package mutation_test

import (
	"slices"
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/onsi/gomega"
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
	suite.Suite
}

type MutationRecorder struct {
	Records []Record
}

func (r *MutationRecorder) Clear() { r.Records = nil }

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

func (s *MutationObserverTestSuite) TestObserveChildListNoSubtree() {
	var recorder MutationRecorder
	observer := NewObserver(&recorder)

	doc := html.NewHTMLDocument(nil)
	body := doc.Body()

	observer.Observe(body, ChildList)

	div := doc.CreateElement("div")
	div.SetAttribute("id", "parent")
	body.AppendChild(div)

	observer.Flush()
	s.Assert().Equal([]dom.Node{body}, recorder.Targets())
	s.Assert().Equal([]dom.Node{div}, recorder.AddedNodes())
	s.Assert().Empty(recorder.RemovedNodes())
	recorder.Clear()

	div.AppendChild(doc.CreateElement("div"))
	observer.Flush()
	s.Assert().Empty(recorder.Records, "Child node was notified when subtree=false")

	body.RemoveChild(div)
	observer.Flush()
	s.Assert().Equal([]dom.Node{body}, recorder.Targets())
	g := gomega.NewWithT(s.T())
	g.Expect(recorder.RemovedNodes()).To(gomega.ConsistOf(div))
	s.Assert().Empty(recorder.AddedNodes())
}

func (s *MutationObserverTestSuite) TestObserveChildListWithSubtree() {
	var recorder MutationRecorder
	observer := NewObserver(&recorder)

	doc := html.NewHTMLDocument(nil)
	body := doc.Body()

	div := doc.CreateElement("div")
	body.AppendChild(div)
	observer.Observe(body, ChildList, Subtree)

	child := doc.CreateElement("div")
	div.AppendChild(child)

	observer.Flush()
	s.Assert().Equal([]dom.Node{div}, recorder.Targets())
	s.Assert().Equal([]dom.Node{child}, recorder.AddedNodes())
}

func TestValidOptions(t *testing.T) {
	assert := assert.New(t)
	var rec = new(MutationRecorder)

	doc := html.NewHTMLDocument(nil)

	assert.NoError(NewObserver(rec).Observe(doc, ChildList), "Error when including ChildList")
	assert.NoError(NewObserver(rec).Observe(doc, Attributes), "Error when including Attributes")
	assert.NoError(NewObserver(rec).Observe(doc, CharacterData),
		"Error when including CharacterData")

	observer := NewObserver(rec)
	assert.ErrorIs(observer.Observe(doc), gosterror.TypeError{},
		"Error when calling with no arguments")
	assert.ErrorIs(
		NewObserver(rec).Observe(doc,
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
