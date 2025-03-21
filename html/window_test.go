package html_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

type WindowTestSuite struct {
	gosttest.GomegaSuite
}

func (s *WindowTestSuite) TestDocumentIsAnHTMLDocument() {
	win, err := NewWindowReader(strings.NewReader("<html><body></body></html>"))
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(win.Document().DocumentElement()).To(BeHTMLElement())
}

func (s *WindowTestSuite) TestDocumentWithDOCTYPE() {
	win, err := NewWindowReader(strings.NewReader("<!DOCTYPE HTML><html><body></body></html>"))
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(win.Document().FirstChild().NodeType()).To(Equal(dom.NodeTypeDocumentType))
}

func TestWindow(t *testing.T) {
	suite.Run(t, new(WindowTestSuite))
}
