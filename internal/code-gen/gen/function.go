package gen

import (
	"github.com/dave/jennifer/jen"
	g "github.com/gost-dom/generators"
)

// Function represents a function, not a function type.
type Function struct {
	receiver   g.Generator
	name       string
	typeParams []g.Generator
	params     []g.Generator
	retValues  []g.Generator
	body       g.Generator
}
type FunctionOption func(*Function)

func NewFunction(opts ...FunctionOption) Function {
	var res Function
	for _, opt := range opts {
		opt(&res)
	}
	return res
}

func FunctionName(name string) FunctionOption {
	return func(fn *Function) { fn.name = name }
}

func pairToGen(name, t g.Generator) g.Generator {
	if name == nil {
		panic("nil name")
	}
	if t == nil {
		return name
	} else {
		return g.Raw(name.Generate().Add(t.Generate()))
	}
}

func FunctionTypeParam(t g.Generator) FunctionOption {
	return func(fn *Function) {
		fn.typeParams = append(fn.typeParams, t)
	}
}

func FunctionParam(name, type_ g.Generator) FunctionOption {
	return func(fn *Function) { fn.params = append(fn.params, pairToGen(name, type_)) }
}

func FunctionReturnType(type_ g.Generator) FunctionOption {
	return func(fn *Function) { fn.retValues = append(fn.retValues, type_) }
}
func FunctionNamedReturn(name, type_ g.Generator) FunctionOption {
	return func(fn *Function) { fn.retValues = append(fn.retValues, pairToGen(name, type_)) }
}

func FunctionBody(body g.Generator) FunctionOption {
	return func(fn *Function) { fn.body = body }
}

func (f Function) Generate() *jen.Statement {
	res := jen.Func()
	if f.receiver != nil {
		res.Params(f.receiver.Generate())
	}
	if f.name != "" {
		res.Id(f.name)
	}
	if f.typeParams != nil {
		res.Index(g.ToJenCodes(f.typeParams)...)
	}
	res.Params(g.ToJenCodes(f.params)...)
	if len(f.retValues) > 0 {
		res.Params(g.ToJenCodes(f.retValues)...)
	}
	// switch len(f.retValues) {
	// case 0:
	// case 1:
	// 	res.Add(f.retValues[0].Generate())
	// default:
	// 	res.Params(g.ToJenCodes(f.retValues)...)
	// }
	if f.body != nil {
		res.Block(f.body.Generate())
	}
	return res
}
