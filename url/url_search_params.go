package url

import (
	netURL "net/url"
	"strings"
)

type URLSearchParams struct {
	// Using netURL.Values
	values netURL.Values
}

func ParseURLSearchParams(s string) (URLSearchParams, error) {
	if strings.HasPrefix(s, "?") {
		s = s[1:]
	}
	values, err := netURL.ParseQuery(s)
	return URLSearchParams{values}, err
}

// Read

func (p URLSearchParams) Size() int      { panic("TODO") }
func (p URLSearchParams) String() string { return p.values.Encode() }

func (p URLSearchParams) Get(
	key string,
) (string, bool) {
	return p.values.Get(key), p.values.Has(key)
}
func (p URLSearchParams) GetAll(key string) []string   { return p.values[key] }
func (p URLSearchParams) Has(string) bool              { panic("TODO") }
func (p URLSearchParams) HasValue(string, string) bool { panic("TODO") }

// Mutate

func (p *URLSearchParams) ensureValid() {
	if p.values == nil {
		p.values = make(netURL.Values)
	}
}

func (p *URLSearchParams) Append(key string, val string) { p.ensureValid(); p.values.Add(key, val) }
func (p *URLSearchParams) Delete(string)                 { panic("TODO") }
func (p *URLSearchParams) DeleteValue(string, string)    { panic("TODO") }
func (p *URLSearchParams) Set(string, string)            { panic("TODO") }
func (p *URLSearchParams) Sort()                         { panic("TODO") }
