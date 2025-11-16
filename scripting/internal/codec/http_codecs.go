package codec

import (
	"fmt"
	"io"
	"strings"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// DecodeRequestBody constructs an [io.Reader] used for HTTP request bodies for
// both [XMLHttpRequest.send] and fetch's [RequestInit.body], which supports
// almost identical values.
//
// Note: Not all valid body types aren't yet supported
//
// [XMLHttpRequest.send]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/send
// [RequestInit.body]: https://developer.mozilla.org/en-US/docs/Web/API/RequestInit
func DecodeRequestBody[T any](s js.Scope[T], val js.Value[T]) (io.Reader, error) {
	if val == nil {
		return nil, nil
	}
	if val.IsUndefined() || val.IsNull() {
		return nil, nil
	}
	if val.IsString() {
		return strings.NewReader(val.String()), nil
	}
	if obj, ok := val.AsObject(); ok {
		if res, ok := obj.NativeValue().(*html.FormData); ok {
			return res.GetReader(), nil
		}
	}
	return nil, fmt.Errorf("gost-dom/codec: request body of type not supported (yet): %v", val)
}
