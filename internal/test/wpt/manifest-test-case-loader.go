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

func (s manifestTestCaseSource) testCases(
	ctx context.Context,
	cancel context.CancelCauseFunc,
) <-chan TestCase {
	return s.filteredTests(ctx, cancel)
}

func (s manifestTestCaseSource) filteredTests(
	ctx context.Context,
	cancelCause context.CancelCauseFunc,
) <-chan TestCase {
	ch := make(chan TestCase, 8)
	testCases := s.loadManifest(ctx, cancelCause)
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
	return ch
}

func (s manifestTestCaseSource) loadManifest(
	ctx context.Context,
	cancelCause context.CancelCauseFunc,
) <-chan TestCase {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.href, nil)
	if err != nil {
		cancelCause(fmt.Errorf("load manifest: create request: %w", err))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		cancelCause(fmt.Errorf("load manifest: %w", err))
	}
	ch := make(chan TestCase, 1)
	go func() {
		defer func() {
			res.Body.Close()
			close(ch)
		}()
		if err := ParseManifestTo(ctx, res.Body, ch, s.options.Logger()); err != nil {
			s.options.logger.Error("ERROR LOADING MANIFEST", "err", err)
			cancelCause(err)
		}
	}()
	return ch
}
