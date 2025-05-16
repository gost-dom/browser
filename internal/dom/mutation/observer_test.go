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
	matchers "github.com/gost-dom/browser/testing/gomega-matchers"
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
	gosttest.GomegaSuite
}

func (s *MutationObserverTestSuite) TestObserveChildListNoSubtree() {
	doc := html.NewHTMLDocument(nil)
	body := doc.Body()

	recorder := initMutationRecorder(body, ChildList)
	subtreeRecorder := initMutationRecorder(body, ChildList, Subtree)

	div := doc.CreateElement("div")
	grandChild := doc.CreateElement("div")
	grandChild.SetAttribute("id", "grand")
	div.SetAttribute("id", "parent")
	body.AppendChild(div)
	div.AppendChild(grandChild)

	s.Expect(recorder).To(HaveRecorded(
		And(HaveType("childList"), HaveTarget(body), HaveAddedNodes(div)),
	))
	s.Expect(subtreeRecorder).To(HaveRecorded(
		And(HaveType("childList"), HaveTarget(body), HaveAddedNodes(div)),
		And(
			HaveType("childList"),
			HaveTarget(div),
			HaveAddedNodes(grandChild),
			HavePrevSibling(nil),
		),
	))
	s.Expect(subtreeRecorder.RemovedNodes()).To(BeEmpty())

	recorder.Clear() // Clear so we can see in isolation what is removed
	subtreeRecorder.Clear()

	newGrand2 := doc.CreateElement("div")
	newGrand1 := doc.CreateElement("div")
	newGrand1.SetAttribute("id", "new-grand")
	div.AppendChild(newGrand2)

	recorder.Clear() // Clear so we can see in isolation what is removed
	subtreeRecorder.Clear()

	div.InsertBefore(newGrand1, newGrand2)
	s.Expect(subtreeRecorder).To(HaveRecorded(
		And(
			HaveType("childList"),
			HaveTarget(div),
			HavePrevSibling(grandChild),
			HaveNextSibling(newGrand2),
		),
	))

	div.RemoveChild(newGrand2)
	div.RemoveChild(newGrand1)

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

func (s *MutationObserverTestSuite) TestSetInnerHTMLRemovesChildren() {
	doc := html.NewHTMLDocument(html.NewWindow())
	body := doc.Body()
	parent := doc.CreateElement("div")
	body.Append(parent)
	d1 := doc.CreateElement("div")
	d2 := doc.CreateElement("div")

	rec := initMutationRecorder(parent, ChildList)
	parent.Append(d1, d2)

	s.Expect(rec).To(HaveRecorded(HaveAddedNodes(d1, d2)))

	rec.Clear()
	parent.SetInnerHTML("<div>Foo</div>")
	s.Expect(rec).To(HaveRecorded(
		And(
			HaveRemovedNodes(d1, d2),
			HaveAddedNodes(matchers.HaveOuterHTML("<div>Foo</div>")),
		)))
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
			AttributeFilter("dummy")),
		gosterror.TypeError{},
		"Error when _all_ non other options than one of the 3 required is passed",
	)
}

func (s *MutationObserverTestSuite) TestAttributeChanges() {
	doc := html.NewHTMLDocument(nil)
	parent := doc.CreateElement("div")
	child := doc.CreateElement("div")
	doc.Body().AppendChild(parent)
	parent.AppendChild(child)
	parent.SetAttribute("data-x", "Old value")

	childRecorder := initMutationRecorder(parent, ChildList)
	rec1 := initMutationRecorder(parent, Attributes, AttributeOldValue)

	parent.AppendChild(doc.CreateElement("div")) // Should not be recorded
	parent.SetAttribute("data-x", "New x value")
	parent.SetAttribute("data-y", "New y value")
	parent.SetAttribute("data-z", "New z value")

	rec1.Flush()
	childRecorder.Flush()

	s.Assert().Equal(1, len(childRecorder.Records))
	s.Assert().Equal(3, len(rec1.Records))
	s.Expect(rec1.Records).
		To(gomega.HaveEach(gomega.HaveField("Type", string(dom.ChangeEventAttributes))))
	s.Expect(rec1.Records).
		To(gomega.HaveEach(gomega.HaveField("Target", Equal(parent))))
	s.Expect(rec1.Records[0]).To(gomega.HaveField("AttributeName", HaveValue(Equal("data-x"))))

	s.Expect(rec1.Records[0]).To(HaveOldValue("Old value"))
}

func Test(t *testing.T) {
	suite.Run(t, new(MutationObserverTestSuite))
}

func (s *MutationObserverTestSuite) TestChangeCDataValue() {
	doc := html.NewHTMLDocument(nil)
	textNode := doc.CreateText("Original value")
	commentNode := doc.CreateComment("Original comment")
	doc.Body().AppendChild(textNode)
	doc.Body().AppendChild(commentNode)

	rec1 := initMutationRecorder(doc.Body(), CharacterData, Subtree)
	rec2 := initMutationRecorder(doc.Body(), CharacterData)
	rec3 := initMutationRecorder(textNode, CharacterData, CharacterDataOldValue)
	rec4 := initMutationRecorder(commentNode, CharacterData, CharacterDataOldValue)

	textNode.SetData("New value")
	commentNode.SetData("New comment")

	s.Expect(rec2.MutationRecorder.Records).To(BeEmpty())
	s.Expect(rec1).To(HaveRecorded(
		And(
			HaveType(string(dom.ChangeEventCData)),
			HaveTarget(textNode),
			HaveOldValue(BeNil()),
		),
		And(
			HaveType(string(dom.ChangeEventCData)),
			HaveTarget(commentNode),
		), // , HaveOldValue("Original value")),
	))
	s.Expect(rec3).To(HaveRecorded(
		And(HaveType(string(dom.ChangeEventCData)), HaveOldValue("Original value")),
	))
	s.Expect(rec4).To(HaveRecorded(
		And(HaveType(string(dom.ChangeEventCData)), HaveOldValue("Original comment")),
	))
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
	FlusherSet
	Records []Record
}

func (r *MutationRecorder) Clear() {
	r.Records = nil
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
