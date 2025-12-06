package fetch

import (
	"iter"
	"testing"

	"github.com/gost-dom/browser/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestHeadersReturnSeparateSetCookieValues(t *testing.T) {
	h := &Headers{}
	h.Append("set-cookie", "foo=bar")
	h.Append("set-COOKIE", "fizz=buzz; domain=example.com")

	pairs := collectPairs(h.All())
	if !assert.Len(t, pairs, 2) {
		return
	}

	assert.Equal(t, "set-cookie", string(pairs[0].first))
	assert.Equal(t, "foo=bar", string(pairs[0].second), "first cookie")

	assert.Equal(t, "set-cookie", string(pairs[1].first))
	assert.Equal(t, "fizz=buzz; domain=example.com", string(pairs[1].second), "second cookie")
}

func TestHeadersGetWithMultipleSameKeyedValues(t *testing.T) {
	h := &Headers{}
	h.Append("foo", "bar")
	// Insert both a header _before_ and _after_
	h.Append("accept-encoding", "text/html")
	h.Append("x-stuff", "buzz")
	h.Append("foo", "baz")

	foo, _ := h.Get("foo")
	assert.Equal(t, "bar, baz", foo)

	h.Set("foo", "barbaz")

	foo, _ = h.Get("foo")
	assert.Equal(t, "barbaz", foo)
}

func TestHeaderIterWhenModified(t *testing.T) {
	h := &Headers{}
	h.Append("fizz", "buzz")
	h.Append("X-Header", "test")
	next, stop := iter.Pull2(h.All())
	defer stop()

	var key types.ByteString
	var val types.ByteString
	var ok bool

	key, val, ok = next()
	assert.True(t, ok)
	assert.Equal(t, "fizz", string(key))
	assert.Equal(t, "buzz", string(val))

	h.Append("set-cookie", "a=b")
	key, val, ok = next()
	assert.True(t, ok)
	assert.Equal(t, "set-cookie", string(key))

	key, val, ok = next()
	assert.True(t, ok)
	assert.Equal(t, "x-header", string(key))

	_, _, ok = next()
	assert.False(t, ok)
}

func TestHeadersIterateSameValue(t *testing.T) {
	h := &Headers{}
	h.Append("fizz", "buzz")
	h.Append("foo", "bar")
	h.Append("foo", "baz")

	pairs := collectPairs(h.All())
	assert.Len(t, pairs, 2)
}

type pair[T, U any] struct {
	first  T
	second U
}

func collectPairs[T, U any](s iter.Seq2[T, U]) []pair[T, U] {
	var res []pair[T, U]
	for k, v := range s {
		res = append(res, pair[T, U]{k, v})
	}
	return res
}
