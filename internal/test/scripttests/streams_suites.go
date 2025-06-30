package scripttests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func testStreams(t *testing.T, shf ScriptHostFactory) {
	suite.Run(t, NewStreamsSuite(shf))
}

type StreamsSuite struct {
	ScriptHostFactorySuite
}

func NewStreamsSuite(f ScriptHostFactory) *StreamsSuite {
	return &StreamsSuite{ScriptHostFactorySuite: *NewScriptHostFactorySuite(f)}
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
