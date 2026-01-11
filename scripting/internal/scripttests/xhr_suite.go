package scripttests

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type XMLHttpRequestSuite struct {
	gosttest.GomegaSuite
	engine     html.ScriptEngine
	body       []byte
	actualPath string
}

func NewXMLHttpRequestSuite(e html.ScriptEngine) *XMLHttpRequestSuite {
	return &XMLHttpRequestSuite{engine: e}
}

func (s *XMLHttpRequestSuite) TestInheritance() {
	w := s.initWindow()
	s.Expect(
		w.Eval("new XMLHttpRequest() instanceof EventTarget"),
	).To(
		BeTrue(),
		"XMLHttpRequest is an EventTarget",
	)
}

func (s *XMLHttpRequestSuite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		var err error
		s.body, err = io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
	}
	s.actualPath = r.URL.Path
}

func (s *XMLHttpRequestSuite) TestLoadEvent() {
	evt := make(chan bool)

	w := s.initWindow()

	w.AddEventListener("go:home", event.NewEventHandlerFunc(func(e *event.Event) error {
		go func() {
			evt <- true
		}()
		return nil
	}))
	w.MustRun(`
		const xhr = new XMLHttpRequest();
		let loadEvent;
		let loadendEvent;
		xhr.addEventListener("load", e => {
			loadEvent = e
			window.dispatchEvent(new CustomEvent("go:home"));
		})
		xhr.addEventListener("loadend", e => {
			loadendEvent = e
			window.dispatchEvent(new CustomEvent("go:home"));
		})
		xhr.open("GET", "/", true);
		xhr.send()
	`)

	ctx, cancel := context.WithTimeout(s.T().Context(), 100*time.Millisecond)
	defer cancel()
	select {
	case <-evt:
	case <-ctx.Done():
		s.T().Error("Timeout waiting for event")
	}
}

func (s *XMLHttpRequestSuite) TestOnloadAttribute() {
	evt := make(chan bool)
	w := s.initWindow()
	w.AddEventListener("go:home", event.NewEventHandlerFunc(func(e *event.Event) error {
		go func() {
			evt <- true
		}()
		return nil
	}))
	w.MustRun(`
		const xhr = new XMLHttpRequest();
		let loadEvent;
		let loadendEvent;
		xhr.onload = function() {
			window.dispatchEvent(new CustomEvent("go:home"));
			loadEvent = e
		}
		xhr.onloadend = function() {
			loadendEvent = e
		}
		xhr.open("GET", "/PATH", true);
		xhr.send()
	`)
	ctx, cancel := context.WithTimeout(s.T().Context(), 100*time.Millisecond)
	defer cancel()
	select {
	case <-evt:
	case <-ctx.Done():
		s.T().Error("Timeout waiting for event")
	}
	s.Expect(s.actualPath).To(Equal("/PATH"))
}

func (s *XMLHttpRequestSuite) TestSendNullBody() {
	w := s.initWindow()
	s.Expect(w.Eval(`
			const xhr = new XMLHttpRequest();
		xhr.open("GET", "/", false);
		xhr.send(null)
		xhr.status
	`)).To(BeEquivalentTo(200))
}

func (s *XMLHttpRequestSuite) initWindow() htmltest.WindowHelper {
	b := browsertest.InitBrowser(s.T(), s, s.engine)
	return b.OpenWindow("http://example.com")
}

func (s *XMLHttpRequestSuite) TestSendFormData() {
	w := s.initWindow()
	s.Expect(w.Eval(`
		const xhr = new XMLHttpRequest();
		const data = new FormData()
		data.append("k1", "v1")
		data.append("k2", "v2")
		xhr.open("GET", "/", false);
		xhr.send(data)
		xhr.status
	`)).To(BeEquivalentTo(200))
	s.Expect(string(s.body)).To(Equal("k1=v1&k2=v2"))
}

func (s *XMLHttpRequestSuite) TestSendString() {
	w := s.initWindow()
	s.Expect(w.Eval(`
		const xhr = new XMLHttpRequest();
		xhr.open("POST", "/", false)
		xhr.send("body contents")
		xhr.status
	`)).To(BeEquivalentTo(200))
	s.Expect(string(s.body)).To(Equal("body contents"))
}
