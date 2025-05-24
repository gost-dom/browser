package wrappers

import g "github.com/gost-dom/generators"

type CallbackContext struct{ g.Value }

func NewCallbackContext(id g.Generator) CallbackContext {
	return CallbackContext{g.ValueOf(id)}
}

func (c CallbackContext) AssignFrom(host, cbInfo g.Generator) g.Generator {
	return g.Assign(
		c.Value,
		g.NewValue("newArgumentHelper").Call(host, cbInfo),
	)
}

func (c CallbackContext) GetInstance() g.Generator {
	return c.Field("Instance").Call()
}

func (c CallbackContext) Context() g.Value {
	return c.Field("Context").Call()
}

func (c CallbackContext) IllegalConstructor() g.Generator {
	return c.ReturnWithTypeError("Illegal constructor")
}

func (c CallbackContext) ReturnWithTypeError(msg string) g.Generator {
	return c.Field("ReturnWithTypeError").Call(g.Lit(msg))
}
func (c CallbackContext) ReturnWithValue(val g.Generator) g.Generator {
	return c.Field("ReturnWithValue").Call(val)
}
