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
