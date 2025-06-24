package scripttests

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type XMLHttpRequestSuite struct {
	ScriptHostSuite
	body       []byte
	actualPath string
}

func NewXMLHttpRequestSuite(h html.ScriptHost) *XMLHttpRequestSuite {
	return &XMLHttpRequestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *XMLHttpRequestSuite) TestInheritance() {
	s.Expect(
		s.Eval("new XMLHttpRequest() instanceof EventTarget"),
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

	s.OpenWindow("http://example.com", s)

	s.Window.AddEventListener("go:home", event.NewEventHandlerFunc(func(e *event.Event) error {
		go func() {
			evt <- true
		}()
		return nil
	}))
	s.MustRunScript(`
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

	s.OpenWindow("http://example.com", s)
	s.Window.AddEventListener("go:home", event.NewEventHandlerFunc(func(e *event.Event) error {
		go func() {
			evt <- true
		}()
		return nil
	}))
	s.Expect(s.Window.Run(`
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
		`)).To(Succeed())
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
	s.OpenWindow("http://example.com", s)
	s.Expect(s.Window.Eval(`
			const xhr = new XMLHttpRequest();
		xhr.open("GET", "/", false);
		xhr.send(null)
		xhr.status
	`)).To(BeEquivalentTo(200))
}

func (s *XMLHttpRequestSuite) TestSendFormData() {
	s.OpenWindow("http://example.com", s)
	s.Expect(s.Window.Eval(`
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
	s.OpenWindow("http://example.com", s)
	s.Expect(s.Window.Eval(`
		const xhr = new XMLHttpRequest();
		xhr.open("POST", "/", false)
		xhr.send("body contents")
		xhr.status
	`)).To(BeEquivalentTo(200))
	s.Expect(string(s.body)).To(Equal("body contents"))
}
