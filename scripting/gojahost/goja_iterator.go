package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaIteratorInstance struct {
	vm   *goja.Runtime
	next func() (js.Value[jsTypeParam], error, bool)
	stop func()
}
