package scripting

import g "github.com/gost-dom/generators"

type scriptEngine struct{ g.Value }

func (e scriptEngine) ConfigureGlobalScope(scope string, base g.Generator) g.Generator {
	return e.Field("ConfigureGlobalScope").Call(g.Lit(scope), base)
}

func (e scriptEngine) Class(className string) g.Generator {
	return e.Field("Class").Call(g.Lit(className))
}
