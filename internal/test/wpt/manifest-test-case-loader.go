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

func (s manifestTestCaseSource) fetchManifest(ctx context.Context) (res *http.Response, err error) {
	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, s.href, nil)
	if err != nil {
		err = fmt.Errorf("load manifest: create request: %w", err)
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("load manifest: %w", err)
	}
	return
}

func (s manifestTestCaseSource) loadManifest(
	ctx context.Context,
	cancelCause context.CancelCauseFunc,
) <-chan TestCase {
	ch := make(chan TestCase, 1)
	go func() {
		defer close(ch)
		res, err := s.fetchManifest(ctx)
		if err != nil {
			cancelCause(err)
			return
		}
		defer res.Body.Close()
		if err := ParseManifestTo(ctx, res.Body, ch, s.options.Logger()); err != nil {
			s.options.logger.Error("ERROR LOADING MANIFEST", "err", err)
			cancelCause(err)
		}
	}()
	return ch
}
