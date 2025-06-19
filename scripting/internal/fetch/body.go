package fetch

import (
	"io"

	"github.com/gost-dom/browser/internal/fetch"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w Body[T]) json(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	p := cbCtx.NewPromise()
	res = p
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(instance)
	if err != nil {
		p.Reject(err)
	} else {
		if js, err := cbCtx.JSONParse(string(b)); err == nil {
			p.Resolve(js)
		} else {
			p.Reject(err)
		}
	}
	return
}
