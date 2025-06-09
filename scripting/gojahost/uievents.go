package gojahost

/*
func (w uIEventWrapper) constructor(
	c goja.ConstructorCall, vm *goja.Runtime,
) *goja.Object {
	type_ := c.Arguments[0].String()
	w.storeInternal(uievents.NewUIEvent(type_), c.This)
	return nil
}

type mouseEventWrapper struct{ uIEventWrapper }
type pointerEventWrapper struct{ mouseEventWrapper }

func createUIEventWrapper(instance *GojaContext) uIEventWrapper {
	return uIEventWrapper{newBaseInstanceWrapper[uievents.UIEvent](instance)}
}

func newMouseEventWrapper(instance *GojaContext) wrapper {
	return &mouseEventWrapper{createUIEventWrapper(instance)}
}

func newPointerEventWrapper(instance *GojaContext) wrapper {
	return &pointerEventWrapper{mouseEventWrapper{createUIEventWrapper(instance)}}
}
*/
