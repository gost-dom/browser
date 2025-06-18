package fetch

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/url"
)

type Fetch struct {
	BrowsingContext html.BrowsingContext
}

func New(bc html.BrowsingContext) Fetch { return Fetch{bc} }

func (f Fetch) NewRequest(url string) Request {
	return Request{
		url: url,
		bc:  f.BrowsingContext,
	}
}

type RequestOption func(*Request)

type Request struct {
	url string
	bc  html.BrowsingContext
}

func (r Request) URL() string { return url.ParseURLBase(r.url, r.bc.LocationHREF()).Href() }

func (f Fetch) Fetch(req Request) *Response {
	return &Response{Status: 404}
}

type Response struct {
	Status int
}
