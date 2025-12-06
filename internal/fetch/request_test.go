package fetch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestHeadersInvalidValue(t *testing.T) {
	f := Fetch{}
	r := f.NewRequest("")
	r.Headers.Append("Accept-Encoding", "text/html")
	assert.False(t, r.Headers.Has("accept-charset"))
	assert.False(t, r.Headers.Has("Accept-Charset"))
}
