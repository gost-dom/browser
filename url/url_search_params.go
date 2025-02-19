package url

import (
	netURL "net/url"
	"slices"
	"strings"
)

type URLSearchParams struct {
	values netURL.Values
	url    *URL
}

func ParseURLSearchParams(s string) (URLSearchParams, error) {
	if strings.HasPrefix(s, "?") {
		s = s[1:]
	}
	values, err := netURL.ParseQuery(s)
	return URLSearchParams{values, nil}, err
}

// Read

func (p URLSearchParams) Size() int                  { panic("TODO") }
func (p URLSearchParams) String() string             { return p.values.Encode() }
func (p URLSearchParams) GetAll(key string) []string { return p.values[key] }
func (p URLSearchParams) Has(key string) bool        { _, ok := p.Get(key); return ok }

func (p URLSearchParams) Get(key string) (string, bool) {
	return p.values.Get(key), p.values.Has(key)
}

func (p URLSearchParams) HasValue(key string, val string) bool {
	v, ok := p.Get(key)
	return ok && v == val
}

// Mutate

func (p *URLSearchParams) ensureValid() {
	if p.values == nil {
		p.values = make(netURL.Values)
	}
}
func (p *URLSearchParams) postUpdate() {
	if p.url != nil {
		p.url.SetSearch(p.String())
	}
}

func (p *URLSearchParams) Append(key string, val string) {
	p.ensureValid()
	p.values.Add(key, val)
	p.postUpdate()
}
func (p *URLSearchParams) Delete(key string) { p.ensureValid(); p.values.Del(key); p.postUpdate() }
func (p *URLSearchParams) DeleteValue(key string, val string) {
	p.ensureValid()
	if v, ok := p.values[key]; ok {
		if i := slices.Index(v, val); i >= 0 {
			p.values[key] = slices.Delete(v, i, i+1)
		}
	}
	p.postUpdate()
}

func (p *URLSearchParams) Set(key string, val string) {
	p.ensureValid()
	p.values.Set(key, val)
	p.postUpdate()
}
func (p *URLSearchParams) Sort() {
	panic("URLSearchParams doesn't support Sorting in it's current implementation")
}
