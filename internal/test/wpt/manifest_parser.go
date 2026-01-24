package main

import (
	"context"
	"encoding/json/jsontext"
	"io"
	"log/slog"
	"path"
	"strings"
)

type MaybeError[T any] struct {
	Value T
	Err   error
}

func ParseManifestTo(
	ctx context.Context,
	m io.Reader,
	ch chan<- TestCase,
	l *slog.Logger,
) error {
	d := jsontext.NewDecoder(m)
	var p = parser{logger: l}
	return p.parse(ctx, ch, d)
}

type parser struct {
	logger *slog.Logger
	ignore bool
	prefix []string
}

func (p *parser) parse(
	ctx context.Context,
	ch chan<- TestCase,
	d *jsontext.Decoder,
) error {
	t, err := d.ReadToken()
	if err != nil {
		return err
	}
	switch t.Kind() {
	case '[':
		if err := p.parseArray(ctx, ch, d); err != nil {
			return err
		}
	case '{':
		if err := p.parseObject(ctx, ch, d); err != nil {
			return err
		}
	case '"':
		p.prefix = append(p.prefix, t.String())
	}
	return nil
}

func (p *parser) parseArray(
	ctx context.Context,
	ch chan<- TestCase,
	d *jsontext.Decoder,
) error {
	for {
		t, err := d.ReadToken()
		if err != nil {
			return err
		}
		switch t.Kind() {
		case ']':
			return nil
		case '[':
			if err := p.parseArray(ctx, ch, d); err != nil {
				return err
			}
		case '"':
			name := t.String()
			if isTestFile(name) {
				select {
				case ch <- TestCase{
					Path:         name,
					PathElements: strings.Split(name, "/"),
				}:
				case <-ctx.Done():
					return nil
				}
			}
		}
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
		elements := p.prefix[2:]
		name := path.Join(elements...)
		return TestCase{Path: name, PathElements: elements}, true
	}
	return
}

func isTestFile(fileName string) bool {
	return path.Ext(fileName) == ".html"
}

func (p parser) parseObject(
	ctx context.Context,
	ch chan<- TestCase,
	d *jsontext.Decoder,
) error {
	for {
		t, err := d.ReadToken()
		if err != nil {
			return err
		}
		switch t.Kind() {
		case '}':
			return nil
		case '"':
			logger := p.logger
			prefix := p.prefix
			name := t.String()

			p.prefix = append(prefix, name)
			p.logger = p.logger.WithGroup(name)
			if isTestFile(name) {
				if testCase, ok := p.makeTestCase(); ok {
					select {
					case ch <- testCase:
					case <-ctx.Done():
						return nil
					}
				}
			}
			if err := p.parse(ctx, ch, d); err != nil {
				return err
			}

			p.prefix = prefix
			p.logger = logger
		}
	}
}
