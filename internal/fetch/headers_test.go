package fetch

import (
	"iter"
	"testing"

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
	assert.Equal(t, "foo=bar", string(pairs[0].second))

	assert.Equal(t, "set-cookie", string(pairs[1].first))
	assert.Equal(t, "fizz=buzz; domain=example.com", string(pairs[1].second))
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
