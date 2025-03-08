package url_test

import (
	"errors"
	"testing"

	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/url"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/suite"
)

type urlWrapper struct {
	*URL
}

func (w urlWrapper) SearchParams() urlinterfaces.URLSearchParams {
	return w.SearchParams()
}

type URLTestSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *URLTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *URLTestSuite) TestURLIsValidInterface() {
	var result urlinterfaces.URL = urlWrapper{ParseURL("http://example.com")}
	s.Assert().Equal("http://example.com", result.Href())
}

func (s *URLTestSuite) TestParseURL() {
	result := ParseURL("http://example.com")
	s.Expect(result).To(HaveHRef(Equal("http://example.com")))
}

func (s *URLTestSuite) TestCanParseURL() {
	s.Expect(CanParseURL("http://example.com")).To(BeTrue())
}

func (s *URLTestSuite) TestParseInvalidURL() {
	s.Expect(ParseURL(":foo")).To(BeNil())
	s.Expect(NewUrl(":foo")).Error().To(HaveOccurred())
	s.Expect(CanParseURL(":foo")).To(BeFalse())
}

func (s *URLTestSuite) TestParseURLWithBase() {
	// Following examples are taken from MDN documentation
	// https://developer.mozilla.org/en-US/docs/Web/API/URL_API/Resolving_relative_references

	s.Expect(
		ParseURLBase("articles", "https://developer.mozilla.org/some/path"),
	).To(HaveHref("https://developer.mozilla.org/some/articles"), "URL.parse")

	// Example 2
	s.Expect(NewUrlBase(
		"./article",
		"https://test.example.org/api/",
	)).To(HaveHref("https://test.example.org/api/article"))
	s.Expect(NewUrlBase("article", "https://test.example.org/api/v1")).To(HaveHref(
		"https://test.example.org/api/article"))

	// Example 3
	s.Expect(
		NewUrlBase("./story/", "https://test.example.org/api/v2/"),
	).To(HaveHref("https://test.example.org/api/v2/story/"))
	s.Expect(
		NewUrlBase("./story", "https://test.example.org/api/v2/v3"),
	).To(HaveHref("https://test.example.org/api/v2/story"))

	// Example 4 - parent directory relative
	s.Expect(
		NewUrlBase("../path", "https://test.example.org/api/v1/v2/"),
	).To(HaveHref("https://test.example.org/api/v1/path"))
	s.Expect(
		NewUrlBase("../../path", "https://test.example.org/api/v1/v2/v3"),
	).To(HaveHref("https://test.example.org/api/path"))
	s.Expect(
		NewUrlBase("../../../../path", "https://test.example.org/api/v1/v2/"),
	).To(HaveHref("https://test.example.org/path"))

	// Should handle example 5 - root relative
	s.Expect(
		NewUrlBase("/some/path", "https://test.example.org/api/"),
	).To(HaveHref("https://test.example.org/some/path"))
	s.Expect(
		NewUrlBase("/", "https://test.example.org/api/v1/"),
	).To(HaveHref("https://test.example.org/"))
	s.Expect(
		NewUrlBase("/article", "https://example.com/api/v1/"),
	).To(HaveHref("https://example.com/article"))

	// Should ignore base for a complete URL
	s.Expect(
		NewUrlBase("http://foo.example.com/bar?q=baz", "http://bar.example.com/foobar"),
	).To(HaveHref("http://foo.example.com/bar?q=baz"))

	// Should handle query in the relative URL
	s.Expect(
		NewUrlBase("/bar?q=baz", "http://bar.example.com/foobar"),
	).To(HaveHref("http://bar.example.com/bar?q=baz"))

	// Should handle query in the relative URL
	s.Expect(
		NewUrlBase("/bar#baz", "http://bar.example.com/foobar"),
	).To(HaveHref("http://bar.example.com/bar#baz"))
}

func (s *URLTestSuite) TestCreateObjectURL() {
	s.T().Skip("URL.CreateObjectURL not yet implemented")
}

func (s *URLTestSuite) TestRevokeObjectURL() {
	s.T().Skip("URL.CreateObjectURL not yet implemented")
}

func (s *URLTestSuite) TestSearch() {
	u := ParseURL("https://example.com:1234/path/name?foo=bar")
	u.SetSearch("q=help")
	s.Expect(u.Search()).To(Equal("?q=help"))
	u.SetSearch("?q=help")
	s.Expect(u.Search()).To(Equal("?q=help"))
}

func (s *URLTestSuite) TestPort() {
	u := ParseURL("https://example.com:1234/path/name?foo=bar")
	u.SetPort("")
	s.Expect(u.Host()).To(Equal("example.com"))
	s.Expect(u.Hostname()).To(Equal("example.com"))

	u.SetPort("1234")
	s.Expect(u.Host()).To(Equal("example.com:1234"))
	s.Expect(u.Hostname()).To(Equal("example.com"))
}

func (s *URLTestSuite) TestHostname() {
	u := ParseURL("https://example.com:1234/path/name?foo=bar")
	u.SetHostname("m.example.com")
	s.Expect(u.Host()).To(Equal("m.example.com:1234"))
	s.Expect(u.Hostname()).To(Equal("m.example.com"))
}

func (s *URLTestSuite) TestSearchParams() {
	u := ParseURL("https://example.com:1234/path/name?foo=bar")
	u.SearchParams().Set("foo", "baz")
	s.Expect(u).To(HaveHref("https://example.com:1234/path/name?foo=baz"))
}

func HaveHRef(expected interface{}) types.GomegaMatcher {
	if m, ok := expected.(types.GomegaMatcher); ok {
		return WithTransform(func(u *URL) string { return u.Href() }, m)
	} else {
		return HaveHRef(Equal(expected))
	}
}

func HaveHref(expected string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(u *URL) (bool, error) {
		if u == nil {
			return false, errors.New("URL is nil")
		}
		return u.Href() == expected, nil
	})
}

func TestURL(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(URLTestSuite))
}
