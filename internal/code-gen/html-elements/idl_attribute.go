package htmlelements

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	. "github.com/gost-dom/code-gen/internal"
	g "github.com/gost-dom/generators"
)

type Receiver struct {
	Name g.Generator
	Type g.Generator
}

type IDLAttribute struct {
	AttributeName string
	Receiver      Receiver
	ReadOnly      bool
}

func (a IDLAttribute) Generate() *jen.Statement {
	attrType := g.NewType("string")
	receiver := g.ValueOf(a.Receiver.Name)
	result := g.Id("result")
	getter := g.FunctionDefinition{
		Receiver: g.FunctionArgument(a.Receiver),
		Name:     UpperCaseFirstLetter(a.AttributeName),
		RtnTypes: g.List(attrType),
		Body: g.StatementList(
			g.AssignMany(
				g.List(result, g.Id("_")),
				receiver.Field("GetAttribute").Call(g.Lit(a.AttributeName)),
			),
			g.Return(result),
		),
	}
	l := g.StatementList(
		getter,
	)
	if !a.ReadOnly {
		argument := g.NewValue("val")
		l.Append(
			g.Line,
			g.FunctionDefinition{
				Receiver: getter.Receiver,
				Name:     fmt.Sprintf("Set%s", getter.Name),
				Args:     g.Arg(argument, attrType),
				Body:     receiver.Field("SetAttribute").Call(g.Lit(a.AttributeName), argument),
			})
	}
	return l.Generate()
}
