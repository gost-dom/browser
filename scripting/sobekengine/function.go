package sobekengine

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type function struct {
	value
	f sobek.Callable
}

func newFunction(c *scriptContext, v sobek.Value) function {
	f, ok := sobek.AssertFunction(v)
	if !ok {
		panic(fmt.Sprintf("gost-dom/sobekengine: %v: not a function value", v))
	}
	return function{value{c, v}, f}
}

func (f function) Call(
	this jsObject,
	args ...js.Value[jsTypeParam],
) (js.Value[jsTypeParam], error) {
	v := make([]sobek.Value, len(args))
	for i, a := range args {
		v[i] = unwrapValue(a)
	}
	res, err := f.f(unwrapValue(this), v...)
	err = errors.Join(err, f.ctx.tick())
	return newValue(f.ctx, res), err
}
