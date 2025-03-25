package mutation_test

import (
	"slices"
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/dom/mutation"
	. "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMutationObserverConformsToTheIDLSpecs(t *testing.T) {
	var observer dominterfaces.MutationObserver = &Observer{}

	_, observerConforms := observer.(dominterfaces.MutationObserver)
	assert.True(t, observerConforms,
		"mutation.Observer conforms to the MutationObserver IDL interface",
	)
}

type MutationObserverTestSuite struct {
	gosttest.GomegaSuite
}

type MutationTestHelper struct {
	MutationRecorder
	observer *mutation.Observer
}

func (h *MutationTestHelper) Flush() {
	h.MutationRecorder.Flush()
}

func (s *MutationObserverTestSuite) TestObserveChildListNoSubtree() {
	doc := html.NewHTMLDocument(nil)
	body := doc.Body()

	recorder := initMutationRecorder(body, ChildList)
	subtreeRecorder := initMutationRecorder(body, ChildList, Subtree)

	div := doc.CreateElement("div")
	grandChild := doc.CreateElement("div")
	div.SetAttribute("id", "parent")
	body.AppendChild(div)
	div.AppendChild(grandChild)

	recorder.Flush()
	subtreeRecorder.Flush()

	s.Assert().Equal([]dom.Node{body}, recorder.Targets())
	s.Assert().Equal([]dom.Node{div}, recorder.AddedNodes())
	s.Assert().Equal([]dom.Node{body, div}, subtreeRecorder.Targets())
	s.Assert().Equal([]dom.Node{div, grandChild}, subtreeRecorder.AddedNodes())
	s.Assert().Empty(subtreeRecorder.RemovedNodes())

	recorder.Clear()
	subtreeRecorder.Clear()

	div.RemoveChild(grandChild)
	body.RemoveChild(div)
	recorder.Flush()
	subtreeRecorder.Flush()

	s.Assert().Equal([]dom.Node{body}, recorder.Targets())
	s.Assert().Equal([]dom.Node{div}, recorder.RemovedNodes())
	s.Assert().Equal([]dom.Node{div, body}, subtreeRecorder.Targets())
	s.Assert().Equal([]dom.Node{grandChild, div}, subtreeRecorder.RemovedNodes())
	s.Assert().Empty(subtreeRecorder.AddedNodes())
}

func (s *MutationObserverTestSuite) TestValidOptions() {
	assert := s.Assert()
	var rec = new(MutationRecorder)

	doc := html.NewHTMLDocument(nil)

	assert.NoError(NewObserver(rec, rec).Observe(doc, ChildList), "Error when including ChildList")
	assert.NoError(
		NewObserver(rec, rec).Observe(doc, Attributes),
		"Error when including Attributes",
	)
	assert.NoError(NewObserver(rec, rec).Observe(doc, CharacterData),
		"Error when including CharacterData")

	observer := NewObserver(rec, rec)
	assert.ErrorIs(observer.Observe(doc), gosterror.TypeError{},
		"Error when calling with no arguments")
	assert.ErrorIs(
		NewObserver(rec, rec).Observe(doc,
			Subtree,
			AttributeOldValue,
			CharacterDataOldValue,
			AttributeFilter([]string{"dummy"})),
		gosterror.TypeError{},
		"Error when _all_ non other options than one of the 3 required is passed",
	)
}

func Test(t *testing.T) {
	suite.Run(t, new(MutationObserverTestSuite))
}

/*
TODO: Note this about subtree monitoring:

subtree (optional)

Set to true to extend monitoring to the entire subtree of nodes rooted at
target. All of the other properties are then extended to all of the nodes in the
subtree instead of applying solely to the target node. The default value is
false.

**Note that if a descendant of target is removed, changes in that
descendant subtree will continue to be observed, until the notification about
the removal itself has been delivered.**

source: https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver/observe

*/

type MutationRecorder struct {
	flushers map[mutation.Flusher]struct{}
	Records  []Record
}

func (r *MutationRecorder) Clear() { r.Records = nil }

func (r *MutationRecorder) ensureFlushers() {
	if r.flushers == nil {
		r.flushers = make(map[mutation.Flusher]struct{})
	}
}

func (r *MutationRecorder) AddFlusher(f Flusher) {
	if f == nil {
		panic("MutationRecorder.AddFlusher: f is nil")
	}
	r.ensureFlushers()
	r.flushers[f] = struct{}{}
}

func (r *MutationRecorder) RemoveFlusher(f Flusher) {
	if _, found := r.flushers[f]; !found {
		panic("MutationRecorder.RemoveFlusher: flusher is not added")
	}
	delete(r.flushers, f)
}

func (r *MutationRecorder) Flush() {
	for f := range r.flushers {
		f.Flush()
	}
}

func (r *MutationRecorder) HandleMutation(recs []Record, _ *Observer) {
	r.Records = append(r.Records, recs...)
}

func (r MutationRecorder) Targets() []dom.Node {
	result := make([]dom.Node, len(r.Records))
	for i, r := range r.Records {
		result[i] = r.Target
	}
	return result
}

func (r MutationRecorder) AddedNodes() []dom.Node {
	lists := make([][]dom.Node, len(r.Records))
	for i, r := range r.Records {
		if nodes := r.AddedNodes; nodes != nil {
			lists[i] = nodes.All()
		}
	}
	return slices.Concat(lists...)
}

func (r MutationRecorder) RemovedNodes() []dom.Node {
	lists := make([][]dom.Node, len(r.Records))
	for i, r := range r.Records {
		if nodes := r.RemovedNodes; nodes != nil {
			lists[i] = nodes.All()
		}
	}
	return slices.Concat(lists...)
}

func initMutationRecorder(target dom.Node, options ...func(*Options)) *MutationTestHelper {
	var res MutationTestHelper
	res.observer = NewObserver(&res.MutationRecorder, &res.MutationRecorder)
	res.observer.Observe(target, options...)
	return &res
}
