package gojahost

import "github.com/gost-dom/browser/scripting/internal/js"

type gojaValueFactory struct {
	*GojaContext
}

func newGojaValueFactory(c *GojaContext) js.ValueFactory[jsTypeParam] {
	return nil
}
