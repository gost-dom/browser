package mutation_test

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/dom/mutation"
	. "github.com/gost-dom/browser/internal/dom/mutation"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	_ "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

type MutationTestHelper struct {
	MutationRecorder
	observer *mutation.Observer
}

func (h *MutationTestHelper) Flush() {
	h.MutationRecorder.Flush()
}

func initMutationRecorder(target dom.Node, options ...func(*Options)) *MutationTestHelper {
	var res MutationTestHelper
	res.observer = NewObserver(&res.MutationRecorder, &res.MutationRecorder)
	res.observer.Observe(target, options...)
	return &res
}

func HaveType(t string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return (r.Type == t), nil
	})
}

func HaveTarget(t dom.Node) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return (r.Target == t), nil
	})
}

func HaveNoAddedNodes() types.GomegaMatcher {
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		var nodes []dom.Node
		if r.AddedNodes != nil {
			nodes = r.AddedNodes.All()
		}
		return len(nodes) == 0, nil
	})
}

func HaveAddedNodes(nodes ...dom.Node) types.GomegaMatcher {
	m := HaveExactElements(nodes)
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return m.Match(r.AddedNodes.All())
	})
}

func HaveRemovedNodes(nodes ...dom.Node) types.GomegaMatcher {
	m := HaveExactElements(nodes)
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return m.Match(r.RemovedNodes.All())
	})
}

func HaveRecorded(expected ...types.GomegaMatcher) types.GomegaMatcher {
	tmp := make([]any, len(expected))
	for i, e := range expected {
		tmp[i] = e
	}
	return gomega.WithTransform(
		func(h *MutationTestHelper) []mutation.Record {
			h.Flush()
			return h.Records
		},
		HaveExactElements(tmp...),
	)
}
