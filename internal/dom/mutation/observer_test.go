package mutation_test

import (
	"slices"
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/dom/mutation"
	"github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	_ "github.com/gost-dom/browser/testing/gomega-matchers"
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

	s.Expect(recorder).To(HaveRecorded(
		And(HaveType("childList"), HaveTarget(body), HaveAddedNodes(div)),
	))
	s.Expect(subtreeRecorder).To(HaveRecorded(
		And(HaveType("childList"), HaveTarget(body), HaveAddedNodes(div)),
		And(HaveType("childList"), HaveTarget(div), HaveAddedNodes(grandChild)),
	))
	s.Expect(subtreeRecorder.RemovedNodes()).To(BeEmpty())

	recorder.Clear() // Clear so we can see in isolation what is removed
	subtreeRecorder.Clear()

	div.RemoveChild(grandChild)
	body.RemoveChild(div)

	s.Expect(recorder).To(HaveRecorded(
		And(HaveType("childList"), HaveTarget(body), HaveRemovedNodes(div), HaveNoAddedNodes())))

	s.Expect(subtreeRecorder).To(HaveRecorded(
		And(HaveTarget(div), HaveRemovedNodes(grandChild), HaveNoAddedNodes()),
		And(HaveType("childList"), HaveTarget(body), HaveRemovedNodes(div), HaveNoAddedNodes()),
	))

	s.Assert().Empty(subtreeRecorder.AddedNodes())
}

func (s *MutationObserverTestSuite) TestValidOptions() {
	assert := s.Assert()
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
