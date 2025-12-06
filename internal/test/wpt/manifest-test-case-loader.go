package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// manifestTestCaseSource is a testCaseSource implementation loading a
// MANIFEST.json from the WPT server
type manifestTestCaseSource struct {
	o    options
	href string
}

func (s manifestTestCaseSource) testCases() <-chan TestCase {
	return s.filteredTests(context.Background())
}

func (s manifestTestCaseSource) filteredTests(ctx context.Context) <-chan TestCase {
	ch := make(chan TestCase, 8)
	go func() {
		defer func() { close(ch) }()
	testCaseLoop:
		for testCase := range s.loadManifest(ctx) {
			path := testCase.Path
			for _, include := range includeList {
				if strings.HasPrefix(path, include) {
					for _, exclude := range excludeList {
						if strings.HasPrefix(path, exclude) {
							continue testCaseLoop
						}
					}

					ch <- testCase
					continue testCaseLoop
				}
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
		ParseManifestTo(ctx, res.Body, ch, s.o.Logger())
	}()
	return ch
}
