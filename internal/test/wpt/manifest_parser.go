package main

import (
	"context"
	"encoding/json/jsontext"
	"io"
	"path"
)

func ParseManifest(ctx context.Context, m io.Reader) <-chan TestCase {
	ch := make(chan TestCase)
	go func() {
		d := jsontext.NewDecoder(m)
		var p parser
		p.parse(ctx, ch, d)
		close(ch)
	}()
	return ch
}

type parser struct {
	ignore bool
	prefix []string
}

func (p *parser) parse(ctx context.Context, ch chan TestCase, d *jsontext.Decoder) {
	t, err := d.ReadToken()
	if err != nil {
		panic(err)
	}
	switch t.Kind() {
	case '[':
		for d.PeekKind() != ']' {
			if err := d.SkipValue(); err != nil {
				panic(err)
			}
		}
		t, err := d.ReadToken()
		if err != nil {
			panic(err)
		}
		if t.Kind() != ']' {
			panic("wtf")
		}
		if testCase, ok := p.makeTestCase(); ok {
			ch <- testCase
		}
	case '{':
		p.parseObject(ctx, ch, d)
	case '"':
		p.prefix = append(p.prefix, t.String())
	}
}

func (p parser) makeTestCase() (res TestCase, ok bool) {
	if len(p.prefix) < 3 {
		return
	}
	if p.prefix[0] != "items" || p.prefix[1] != "testharness" {
		return
	}
	if path.Ext(p.prefix[len(p.prefix)-1]) == ".html" {
		name := path.Join(p.prefix[2:]...)
		return TestCase{Path: name}, true
	}
	return
}

func (p parser) parseObject(ctx context.Context, ch chan TestCase, d *jsontext.Decoder) {
	for {
		t, err := d.ReadToken()
		if err != nil {
			panic(err)
		}
		switch t.Kind() {
		case '}':
			return
		case '"':
			prefix := p.prefix
			p.prefix = append(prefix, t.String())
			p.parse(ctx, ch, d)
			p.prefix = prefix
		}
	}

}
