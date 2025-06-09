package scripting

import g "github.com/gost-dom/generators"

/* -------- class -------- */

// class struct represents the scripting/internal/js.Class[T] struct. Methods on
// this type generates method call on the corresponding struct.
type class struct{ g.Value }

func (c class) CreatePrototypeMethod(name string, callback g.Generator) g.Generator {
	return c.Field("CreatePrototypeMethod").Call(g.Lit(name), callback)
}

func (c class) CreateAttribute(name string, getter g.Generator, setter g.Generator) g.Generator {
	return c.Field("CreatePrototypeAttribute").Call(g.Lit(name), getter, setter)
}
