package fetch

import (
	"io"

	"github.com/gost-dom/browser/internal/streams"
)

// Body represents the Body IDL mixin defined in the [fetch API].
//
// [fetch API]: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
type Body interface {
	io.Reader
	Body() streams.ReadableStream
}
