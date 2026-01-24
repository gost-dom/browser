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

func (s manifestTestCaseSource) testCases(ctx context.Context) (<-chan TestCase, <-chan error) {
	return s.filteredTests(ctx)
}

func (s manifestTestCaseSource) filteredTests(ctx context.Context) (<-chan TestCase, <-chan error) {
	ch := make(chan TestCase, 8)
	testCases, errors := s.loadManifest(ctx)
	go func() {
		defer func() { close(ch) }()
		for testCase := range testCases {
			if s.filter == nil || s.filter(testCase) {
				select {
				case ch <- testCase:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return ch, errors
}

func (s manifestTestCaseSource) loadManifest(
	ctx context.Context,
) (<-chan TestCase, <-chan error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.href, nil)
	if err != nil {
		panic(fmt.Sprintf("load manifest: create request: %v", err))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("load manifest: %v", err))
	}
	ch := make(chan TestCase, 1)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			res.Body.Close()
			close(ch)
		}()
		if err := ParseManifestTo(ctx, res.Body, ch, s.options.Logger()); err != nil {
			s.options.logger.Error("ERROR LOADING MANIFEST", "err", err)
			errCh <- err
		}
	}()
	return ch, errCh
}
