package sobekengine

import (
	"errors"

	"github.com/grafana/sobek"
)

type array struct {
	value
	obj *sobek.Object
}

func newArray(c *scriptContext, o *sobek.Object) jsArray {
	return array{value{c, o}, o}
}

func (a array) Push(v jsValue) error {
	push := a.obj.Get("push")
	if push == nil {
		return errors.New("gost-dom/sobekengine: array.push: Underlying object doesn't have a push")
	}
	p, ok := sobek.AssertFunction(push)
	if !ok {
		return errors.New(
			"gost-dom/sobekengine: array.push: Underlying object's push is not a function",
		)
	}
	_, err := p(a.obj, v.Self().value)
	return err
}
