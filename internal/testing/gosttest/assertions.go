package gosttest

import (
	"context"
	"testing"
	"time"
)

type expectOptions struct {
	ctx     context.Context
	timeout time.Duration
	fatal   bool
}

type ExpectOption func(*expectOptions)

func Timeout(d time.Duration) ExpectOption     { return func(o *expectOptions) { o.timeout = d } }
func Context(ctx context.Context) ExpectOption { return func(o *expectOptions) { o.ctx = ctx } }
func FatalOnError() ExpectOption               { return func(o *expectOptions) { o.fatal = true } }

// ExpectReceive is a test helper verifying that a message can be received
func ExpectReceive[T any](t testing.TB, c <-chan T, opts ...ExpectOption) (msg T) {
	var o expectOptions
	for _, opt := range opts {
		opt(&o)
	}
	t.Helper()
	ctx := o.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	timeout := o.timeout
	_, hasDeadline := ctx.Deadline()
	if timeout == 0 && !hasDeadline {
		// If there is no timeout option, and the context doesn't have a
		// deadline already, create a default timeout
		timeout = 100 * time.Millisecond
	}
	if timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	select {
	case <-ctx.Done():
		if o.fatal {
			t.Fatalf("ExpectReceive: timeout exceeded: %v", timeout)
		} else {
			t.Errorf("ExpectReceive: timeout exceeded: %v", timeout)
		}
	case msg = <-c:
	}
	return
}
