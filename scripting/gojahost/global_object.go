package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type globalObject struct {
	ctx *GojaContext
	obj *goja.Object
}

// CreateFunction adds a function to an object in global scope.
func (o globalObject) CreateFunction(name string, cb js.FunctionCallback[jsTypeParam]) {
	o.obj.Set(name, wrapJSCallback(o.ctx, cb))
}
