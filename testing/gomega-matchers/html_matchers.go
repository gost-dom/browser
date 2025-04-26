package matchers

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
)

func BeHTMLElement() GomegaMatcher { return HtmlElementMatcher{} }

type HtmlElementMatcher struct{}

func (m HtmlElementMatcher) Match(value any) (bool, error) {
	_, ok := value.(HTMLElement)
	return ok, nil
}

func (m HtmlElementMatcher) FailureMessage(actual interface{}) (message string) {
	return "Should be en HTMLElement"
}

func (m HtmlElementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Should not be an HTMLElement"
}

func HaveTextContent(expected any) GomegaMatcher {
	matcher, ok := expected.(GomegaMatcher)
	if !ok {
		return HaveTextContent(gomega.Equal(expected))
	}

	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		if e == nil {
			return false, errors.New("Element is nil")
		}
		return matcher.Match(e.TextContent())
	}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have textContent {{.Data.FailureMessage .Actual.TextContent}}", matcher)
}

func HaveInnerHTML(matcher GomegaMatcher) GomegaMatcher {
	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		return matcher.Match(e.InnerHTML())
	}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have textContent {{.Data.FailureMessage .Actual.InnerHTML}}", matcher)
}

func HaveOuterHTML(expected interface{}) GomegaMatcher {
	matcher, ok := expected.(GomegaMatcher)
	if !ok {
		return HaveOuterHTML(gomega.Equal(expected))
	}
	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		return matcher.Match(e.OuterHTML())
	}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have outerHTML {{.Data.FailureMessage .Actual.OuterHTML}}", matcher)
}

func HaveTag(expected string) GomegaMatcher {
	matcher := gomega.Equal(expected)
	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		return matcher.Match(e.TagName())
	}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have tag {{.Data.FailureMessage .Actual.TagName}}", matcher)
}
