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

func (h *MutationTestHelper) Clear() {
	h.Flush()
	h.MutationRecorder.Clear()
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

func HaveOldValue(expected string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return (r.OldValue == expected), nil
	}).WithTemplate("Mutation record old attribute value\n    Expected value: {{ .Data }}\n    Actual value: {{ .Actual.OldValue }}", expected)
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

func HavePrevSibling(node dom.Node) types.GomegaMatcher {
	var m types.GomegaMatcher
	if node == nil {
		m = BeNil()
	} else {
		m = Equal(node)
	}
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		return m.Match(r.PreviousSibling)
	})
}

func HaveAddedNodes(nodes ...any) types.GomegaMatcher {
	var Data struct {
		Actual  []dom.Node
		Matcher types.GomegaMatcher
	}
	matchers := make([]types.GomegaMatcher, len(nodes))
	for i, n := range nodes {
		if m, ok := n.(types.GomegaMatcher); ok {
			matchers[i] = m
		} else {
			matchers[i] = Equal(n)
		}
	}
	Data.Matcher = HaveExactElements(nodes)
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		Data.Actual = r.AddedNodes.All()
		return Data.Matcher.Match(Data.Actual)
	}).WithTemplate("Have added nodes\nActual: {{ .Data.Actual }}\n{{ .Data.Matcher.Formatted }}", &Data)
}

func HaveRemovedNodes(nodes ...dom.Node) types.GomegaMatcher {
	var Data struct {
		Actual  []dom.Node
		Matcher types.GomegaMatcher
	}
	Data.Matcher = HaveExactElements(nodes)
	return gcustom.MakeMatcher(func(r mutation.Record) (bool, error) {
		// if r.RemovedNodes != nil {
		Data.Actual = r.RemovedNodes.All()
		// }
		return Data.Matcher.Match(Data.Actual)
	}).WithTemplate("Have removed nodes\nActual: {{ .Data.Actual }}\n{{ .Data.Matcher.Formatted }}", &Data)
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
