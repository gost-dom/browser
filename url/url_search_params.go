package url

import (
	netURL "net/url"
)

type URLSearchParams struct {
	values netURL.Values
}

func (p URLSearchParams) Size() int {
	panic("TODO")
}

// Read

func (p URLSearchParams) String() string          { return p.values.Encode() }
func (p URLSearchParams) Get(string) string       { panic("TODO") }
func (p URLSearchParams) GetAll(string)           { panic("TODO") }
func (p URLSearchParams) Has(string, string) bool { panic("TODO") }

// Mutate

func (p *URLSearchParams) ensureValid() {
	if p.values == nil {
		p.values = make(netURL.Values)
	}
}

func (p *URLSearchParams) Append(key string, val string) { p.ensureValid(); p.values.Add(key, val) }
func (p *URLSearchParams) Delete(string, string)         { panic("TODO") }
func (p *URLSearchParams) Set(string, string)            { panic("TODO") }
func (p *URLSearchParams) Sort()                         { panic("TODO") }
