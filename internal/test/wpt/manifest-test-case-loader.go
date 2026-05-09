package main

import (
	"context"
	"os"
)

// manifestTestCaseSource is a testCaseSource implementation loading a
// MANIFEST.json from the WPT server
type manifestTestCaseSource struct {
	options options
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

func (s manifestTestCaseSource) loadManifest(
	ctx context.Context,
	cancelCause context.CancelCauseFunc,
) <-chan TestCase {
	ch := make(chan TestCase, 1)
	go func() {
		defer close(ch)
		res, err := os.Open(s.options.manifest)
		if err != nil {
			cancelCause(err)
			return
		}
		defer res.Close()
		if err := ParseManifestTo(ctx, res, ch, s.options.Logger()); err != nil {
			s.options.logger.Error("ERROR LOADING MANIFEST", "err", err)
			cancelCause(err)
		}
	}()
	return ch
}
