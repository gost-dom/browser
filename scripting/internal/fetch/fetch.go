package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func Fetch[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	f := fetch.New(info.Window())
	p := info.NewPromise()
	r := f.Fetch(f.NewRequest(""))
	resp, err := info.Constructor("Response").NewInstance(r)
	if err != nil {
		return nil, err
	}
	p.Resolve(resp)
	return p, nil
}
