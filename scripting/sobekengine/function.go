package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type function struct {
	value
	f sobek.Callable
}

func (f function) Call(
	this js.Object[jsTypeParam],
	args ...js.Value[jsTypeParam],
) (js.Value[jsTypeParam], error) {
	v := make([]sobek.Value, len(args))
	for i, a := range args {
		v[i] = unwrapValue(a)
	}
	res, err := f.f(unwrapValue(this), v...)
	return newValue(f.ctx, res), err
}
