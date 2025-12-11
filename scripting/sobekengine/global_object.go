package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

// globalObject is used to configure an object present in global scope in
// JavaScript. For example, the console object in JavaScript is configured as a
// globalObject.
type globalObject struct {
	ctx *scriptContext
	obj *sobek.Object
}

// CreateFunction adds a function to an object in global scope.
func (o globalObject) CreateFunction(name string, cb js.CallbackFunc[jsTypeParam]) {
	o.obj.Set(name, wrapJSCallback(o.ctx, cb.WithLog("", name)))
}
