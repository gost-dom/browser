package url_test

import (
	"testing"

	i "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	. "github.com/gost-dom/browser/url"
	"github.com/stretchr/testify/suite"
)

type URLSearchParamsSuits struct {
	suite.Suite

	// The test uses the implementation through the interface as this brings
	// design-time validation that the type conforms to the IDL specs
	u i.URLSearchParams
}

func (s *URLSearchParamsSuits) SetupTest() {
	s.u = &URLSearchParams{}
}

func (s *URLSearchParamsSuits) TestAddValue() {
	s.u.Append("foo", "bar")
	s.Assert().Equal("foo=bar", s.u.String())
	v, found := s.u.Get("foo")
	s.Assert().True(found)
	s.Assert().Equal("bar", v)

	// Same value can be added twice
	s.u.Append("foo", "baz")
	s.Assert().Equal("foo=bar&foo=baz", s.u.String())
	s.Assert().Equal([]string{"bar", "baz"}, s.u.GetAll("foo"))

	s.u = &URLSearchParams{}
	s.u.Append("foo", "Bar value")
	s.Assert().Equal("foo=Bar+value", s.u.String(), "Query encoding")
}

func (s *URLSearchParamsSuits) TestGet() {
	s.u.Append("foo", "bar")

	foundValue, ok := s.u.Get("foo")
	s.Assert().True(ok)
	s.Assert().Equal("bar", foundValue)

	missingValue, ok := s.u.Get("baz")
	s.Assert().False(ok)
	s.Assert().Equal("", missingValue)
}

func (s *URLSearchParamsSuits) TestParse() {
	// Test percent decoding
	var u, err = ParseURLSearchParams("foo=bar%20value")
	s.Assert().NoError(err)
	v, found := u.Get("foo")
	s.Assert().True(found)
	s.Assert().Equal("bar value", v)

	// Test + decoding
	u, err = ParseURLSearchParams("foo=bar+value")
	s.Assert().NoError(err)
	v, found = u.Get("foo")
	s.Assert().True(found)
	s.Assert().Equal("bar value", v)

	// Test leading ? is ignores
	u, err = ParseURLSearchParams("?foo=bar%20value")
	s.Assert().NoError(err)
	v, found = u.Get("foo")
	s.Assert().True(found)
	s.Assert().Equal("bar value", v)
}

func (s *URLSearchParamsSuits) TestHas() {
	s.u.Append("foo", "bar")
	s.Assert().True(s.u.Has("foo"))
	s.Assert().False(s.u.Has("bar"))

	s.Assert().True(s.u.HasValue("foo", "bar"))
	s.Assert().False(s.u.HasValue("foo", "baz"))
}

func (s *URLSearchParamsSuits) TestDelete() {
	var u = s.u
	u.Append("foo", "a")
	u.Append("foo", "b")
	u.Append("foo", "c")

	u.DeleteValue("foo", "z")
	s.Assert().Equal([]string{"a", "b", "c"}, u.GetAll("foo"))

	u.DeleteValue("foo", "c")
	s.Assert().Equal([]string{"a", "b"}, u.GetAll("foo"))

	u.Delete("foo")
	_, found := u.Get("foo")
	s.Assert().False(found)
}

func (s *URLSearchParamsSuits) TestSet() {
	u := s.u

	// Add a new value
	u.Set("key1", "a")
	val, found := u.Get("key1")
	s.Assert().True(found)
	s.Assert().Equal("a", val)

	u.Append("key2", "a")
	u.Append("key2", "b")
	u.Set("key2", "c")
	s.Assert().Equal([]string{"c"}, u.GetAll("key2"))
}

func TestURLSearchParams(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(URLSearchParamsSuits))
}
