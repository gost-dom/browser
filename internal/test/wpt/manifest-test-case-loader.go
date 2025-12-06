package main

import (
	"context"
	"fmt"
	"net/http"
)

// manifestTestCaseSource is a testCaseSource implementation loading a
// MANIFEST.json from the WPT server
type manifestTestCaseSource struct {
	options options
	href    string
	filter  func(t TestCase) bool
}

func (s manifestTestCaseSource) testCases() <-chan TestCase {
	return s.filteredTests(context.Background())
}

func (s manifestTestCaseSource) filteredTests(ctx context.Context) <-chan TestCase {
	ch := make(chan TestCase, 8)
	go func() {
		defer func() { close(ch) }()
		for testCase := range s.loadManifest(ctx) {
			if s.filter(testCase) {
				ch <- testCase
			}
		}
	}()
	return ch
}

func (s manifestTestCaseSource) loadManifest(
	ctx context.Context,
) <-chan TestCase {
	res, err := http.Get(s.href)
	if err != nil {
		panic(fmt.Sprintf("load manifest: %v", err))
	}
	ch := make(chan TestCase)
	go func() {
		defer func() {
			res.Body.Close()
			close(ch)
		}()
		ParseManifestTo(ctx, res.Body, ch, s.options.Logger())
	}()
	return ch
}
