package v8gen

import (
	g "github.com/gost-dom/generators"
)

const scriptHostName = "scriptHost"

type WrapperInstance struct{ g.Value }

func (i WrapperInstance) GetScriptHost() g.Value { return i.Field(scriptHostName) }

func (i WrapperInstance) MustGetContext(info g.Generator) g.Generator {
	return i.Method("mustGetContext").Call(info)
}

func (i WrapperInstance) InitializeClass(cls jsClass) g.Generator {
	return i.Field("installPrototype").Call(cls)
}
