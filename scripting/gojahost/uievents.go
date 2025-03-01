package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/uievents"
)

func (w uIEventWrapper) constructor(
	c goja.ConstructorCall, vm *goja.Runtime,
) *goja.Object {
	type_ := c.Arguments[0].String()
	w.storeInternal(uievents.NewUIEvent(type_), c.This)
	return nil
}

type mouseEventWrapper struct{ uIEventWrapper }

func createUIEventWrapper(instance *GojaContext) uIEventWrapper {
	return uIEventWrapper{newBaseInstanceWrapper[uievents.UIEvent](instance)}
}

func newMouseEventWrapper(instance *GojaContext) wrapper {
	return &mouseEventWrapper{createUIEventWrapper(instance)}
}
