package scripting

import (
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

type ReturnValueGenerator struct {
	Data     model.ESConstructorData
	Op       model.ESOperation
	Ctx      CallbackContext
	Receiver g.Generator
	V8       bool
}

func (gen ReturnValueGenerator) Transform(call g.Generator) g.Generator {
	if !gen.Op.HasError && !gen.Op.HasResult() {
		return g.StatementList(
			call,
			g.Return(g.Nil, g.Nil),
		)
	}

	res := gen.Op.RetValues(gen.Data)
	err := g.Id("errCall")
	var vals []g.Generator
	if gen.Op.HasResult() {
		vals = append(vals, res...)
	}
	if gen.Op.HasError {
		vals = append(vals, err)
	}

	stmts := g.StatementList(
		g.AssignMany(vals, call),
	)

	if gen.Op.HasError {
		if gen.Op.HasResult() {
			stmts.Append(
				IfErrorF(err, func(err g.Generator) g.Generator { return g.Return(g.Nil, err) }),
			)
		} else {
			stmts.Append(g.Return(g.Nil, err))
		}
	}

	if gen.Op.HasResult() {
		stmts.Append(g.Return(gen.encodeReturnValue(gen.V8, gen.Ctx, res)))
	}

	return stmts
}

func (gen ReturnValueGenerator) encodeReturnValue(
	v8 bool,
	cbCtx CallbackContext,
	val []g.Generator,
) g.Generator {
	encoder := gen.Op.Encoder(v8, g.ValueOf(gen.Receiver), cbCtx, gen.Data)
	return encoder.Call(val...)
}
