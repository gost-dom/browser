package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/suite"
)

func testStreams(t *testing.T, engine html.ScriptEngine) {
	suite.Run(t, NewStreamsSuite(engine))
}

type StreamsSuite struct {
	ScriptHostFactorySuite
}

func NewStreamsSuite(e html.ScriptEngine) *StreamsSuite {
	return &StreamsSuite{ScriptHostFactorySuite: *NewScriptHostFactorySuite(e)}
}

func (s *StreamsSuite) TestPrototypes() {
	s.Assert().Equal("function",
		s.MustEval("typeof ReadableStream"), "ReadableStream is a function")

	s.Assert().Equal("function",
		s.MustEval("typeof ReadableStreamBYOBReader"), "ReadableStreamBYOBReader is a function")

	s.Assert().Equal("function",
		s.MustEval("typeof ReadableStreamDefaultReader"),
		"ReadableStreamDefaultReader is a function")
}
