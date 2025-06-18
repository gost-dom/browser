package fetch

import (
	"net/http"

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

func (r *Request) URL() string { return url.ParseURLBase(r.url, r.bc.LocationHREF()).Href() }

func (r *Request) do() (*http.Response, error) {
	r.bc.Logger().Info("Get", "url", r.URL())
	c := r.bc.HTTPClient()
	return c.Get(r.URL())
}

func (f Fetch) Fetch(req Request) (*Response, error) {
	resp, err := req.do()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &Response{Status: resp.StatusCode}, nil
}

type Response struct {
	Status int
}
