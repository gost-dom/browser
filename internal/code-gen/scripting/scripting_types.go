package scripting

import g "github.com/gost-dom/generators"

/* -------- class -------- */

// class struct represents the scripting/internal/js.Class[T] struct. Methods on
// this type generates method call on the corresponding struct.
type class struct{ g.Value }

func (c class) CreateOperation(name string, callback g.Generator) g.Generator {
	return c.Field("CreateOperation").Call(g.Lit(name), callback)
}

func (c class) CreateAttribute(name string, getter g.Generator, setter g.Generator) g.Generator {
	return c.Field("CreatePrototypeAttribute").Call(g.Lit(name), getter, setter)
}

/* -------- CallbackContext -------- */

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
	return c.Field("ScriptCtx").Call()
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
func (c CallbackContext) ReturnWithValueErr(val g.Generator) g.Generator {
	return c.Field("ReturnWithValueErr").Call(val)
}

func (c CallbackContext) ReturnWithError(val g.Generator) g.Generator {
	return c.Field("ReturnWithError").Call(val)
}

func (c CallbackContext) ConsumeArg() g.Generator { return c.Field("ConsumeArg").Call() }
func (c CallbackContext) Scope() g.Generator      { return c.Field("Scope").Call() }
func (c CallbackContext) Logger() Logger          { return Logger{c.Field("Logger").Call()} }

/* -------- Logger -------- */

type Logger struct{ g.Value }

func (l Logger) With(args ...g.Generator) g.Generator {
	return Logger{l.Field("With").Call(args...)}
}

func (l Logger) Debug(msg string, args ...g.Generator) g.Generator {
	return l.Field("Debug").Call(append([]g.Generator{g.Lit(msg)}, args...)...)
}
