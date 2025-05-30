package wrappers

import (
	"github.com/gost-dom/code-gen/script-wrappers/model"
	g "github.com/gost-dom/generators"
)

type ReturnValueGenerator struct {
	Data     model.ESConstructorData
	Op       model.ESOperation
	Ctx      CallbackContext
	Receiver g.Generator
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
		stmts.Append(g.Return(gen.encodeReturnValue(gen.Ctx, res)))
	}

	return stmts
}

func (gen ReturnValueGenerator) encodeReturnValue(
	cbCtx CallbackContext,
	val []g.Generator,
) g.Generator {
	encoder := gen.Op.Encoder(gen.Data)
	return g.ValueOf(gen.Receiver).Field(encoder).Call(append([]g.Generator{cbCtx}, val...)...)
}

/*
type LogGenerator struct{ g.Generator }

func (gen LogGenerator) Debug(str string, args ...g.Generator) g.Generator {
	gargs := []g.Generator{g.Lit(str)}
	gargs = append(gargs, args...)
	return g.ValueOf(gen).Field("Debug").Call(gargs...)
}

func (gen LogGenerator) DebugJSCall(
	intf model.ESConstructorData,
	op model.ESOperation,
) g.Generator {
	return gen.Debug(fmt.Sprintf("JS call: %s.%s", intf.Name(), op.Name))
}

type AttributeGetterCallbackGenerator struct {
	data        model.ESConstructorData
	attr        model.ESAttribute
	receiver    g.Generator
	wrapperType g.Generator
	lookup      Transformer
}

func (gen AttributeGetterCallbackGenerator) Generate() *jen.Statement {
	ctx := callbackContext{g.NewValue("context")}
	return g.FunctionDefinition{
		Receiver: g.FunctionArgument{Name: gen.receiver, Type: gen.wrapperType},
		Name:     gen.attr.Getter.CallbackMethodName(),
		Args:     g.Arg(ctx, callbackContextType),
		RtnTypes: g.List(callbackReturnType),
		Body:     gen.GenerateBody(ctx),
	}.Generate()
}

func (gen AttributeGetterCallbackGenerator) GenerateBody(ctx callbackContext) g.Generator {
	op := *gen.attr.Getter
	instance := g.NewValue("instance")
	err := g.Id("instErr")
	res := g.StatementList(
		ctx.Logger().DebugJSCall(gen.data, op),
		g.AssignMany(g.List(instance, err), ctx.instance(gen.data.WrappedType())),
		IfError(err, ctx.ReturnError()),
		ReturnValueGenerator{op, ctx, gen.receiver}.Transform(gen.lookup.Transform(instance)),
	)
	return res
}

type AttributeSetterCallback struct {
	data        model.ESConstructorData
	attr        model.ESAttribute
	receiver    g.Generator
	wrapperType g.Generator
}

func (gen AttributeSetterCallback) Generate() *jen.Statement {
	ctx := callbackContext{g.NewValue("context")}
	return g.FunctionDefinition{
		Receiver: g.FunctionArgument{Name: gen.receiver, Type: gen.wrapperType},
		Name:     gen.attr.Getter.CallbackMethodName(),
		Args:     g.Arg(ctx, callbackContextType),
		RtnTypes: g.List(callbackReturnType),
		Body:     gen.GenerateBody(ctx),
	}.Generate()
}

func (gen AttributeSetterCallback) GenerateBody(ctx CallbackContext) g.Generator {
	op := *gen.attr.Getter
	instance := g.NewValue("instance")
	err := g.Id("instErr")
	res := g.StatementList(
		ctx.Logger().DebugJSCall(gen.data, op),
		g.AssignMany(g.List(instance, err), ctx.instance(gen.data.WrappedType())),
	)
	return res
}
*/
