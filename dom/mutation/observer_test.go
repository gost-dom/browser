package mutation_test

import (
	"slices"
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/mutation"
	"github.com/gost-dom/browser/html"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMutationObserverConformsToTheIDLSpecs(t *testing.T) {
	var observer dominterfaces.MutationObserver = &mutation.Observer{}

	_, observerConforms := observer.(dominterfaces.MutationObserver)
	assert.True(t, observerConforms,
		"mutation.Observer conforms to the MutationObserver IDL interface",
	)
}

type MutationObserverTestSuite struct {
	suite.Suite
}

type MutationRecorder struct {
	Records []mutation.Record
}

func (r *MutationRecorder) HandleMutation(recs []mutation.Record, _ *mutation.Observer) {
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
		lists[i] = r.AddedNodes.All()
	}
	return slices.Concat(lists...)
}

func (s *MutationObserverTestSuite) TestAddDomNode() {
	var recorder MutationRecorder
	observer := mutation.NewObserver(&recorder)

	doc := html.NewHTMLDocument(nil)
	body := doc.Body()

	observer.Observe(body, mutation.ChildList)

	div := doc.CreateElement("div")
	body.AppendChild(div)

	observer.Flush()
	// s.Assert().Equal(1, len(recorder.Targets()))
	s.Assert().Equal([]dom.Node{body}, recorder.Targets())
	s.Assert().Equal([]dom.Node{div}, recorder.AddedNodes())
}

func (s *MutationObserverTestSuite) TestRemoveDomNode() {
	s.T().Skip("TODO: Implement")
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
