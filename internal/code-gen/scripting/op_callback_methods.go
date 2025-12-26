package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

// OpCallbackMethods is used to generate both web IDL interface constructor and
// operation callbacks.
//
// Common to both is that they can accept a long list of arguments, some can be
// optional or variadic.
type OpCallbackMethods struct {
	CallbackMethods
	Op model.Callback
}

func (gen OpCallbackMethods) ConstructorCallbackBody() g.Generator {
	return gen.CtorOrOperationCallback(gen.NativeConstructorCall)
}

func (gen OpCallbackMethods) NativeConstructorCall(
	op model.Callback,
	methodPostFix string,
	args []g.Generator,
) g.Generator {
	return g.Return(
		gen.Receiver().
			Field("CreateInstance" + methodPostFix).
			Call(append([]g.Generator{gen.CbCtx()}, args...)...),
	)
}

func (b OpCallbackMethods) MethodCallbackBody() g.Generator {
	return g.StatementList(
		g.StatementList(
			b.assignInstance(b.Op.Arguments),
			b.CtorOrOperationCallback(b.NativeMethodCall),
		),
	)
}

func (gen OpCallbackMethods) NativeMethodCall(
	op model.Callback,
	methodPostFix string,
	args []g.Generator,
) g.Generator {
	name := op.NativeFunctionName() + methodPostFix
	eval := gen.instance().Field(name).Call(args...)
	return ReturnValueGenerator{
		Data:     gen.Data,
		Op:       op,
		Ctx:      gen.CbCtx(),
		Receiver: gen.Receiver(),
	}.Transform(eval)
}

func (gen OpCallbackMethods) CtorOrOperationCallback(
	callNativeFunc func(op model.Callback, methodPostFix string, args []g.Generator) g.Generator,
) g.Generator {
	// reqArgs are "required" args on the Go side. This include optional args in
	// the IDL specification that have a default value, or customization
	// indicates a default value.
	reqArgs := []g.Generator{}
	optArgs := []g.Generator{}

	errs := []g.Generator{}

	stmts := g.StatementList()
	var noOfConsumed int
	for i, a := range gen.Op.Arguments {
		noOfConsumed = i
		defaultValuer, hasDefault := gen.DefaultValuer(a)
		if a.Optional && !hasDefault && !a.VariadicInGo() {
			break
		}
		noOfConsumed = i + 1
		if a.Ignore {
			stmts.Append(gen.CbCtx().ConsumeArg())
			continue
		}
		arg := g.Id(SanitizeVarName(a.Name))
		reqArg := arg
		if a.VariadicInGo() {
			reqArg = g.Raw(jen.Id(a.Name).Op("..."))
		}
		err := g.Id(fmt.Sprintf("errArg%d", i+1))
		reqArgs = append(reqArgs, reqArg)
		errs = append(errs, err)
		parseArgs := []g.Generator{gen.CbCtx(), g.Lit(a.Name)}
		if !a.Variadic {
			parseArgs = append(parseArgs, defaultValuer)
		}
		var dec = DecodersForArg(gen.Receiver(), a)
		parseArgs = append(parseArgs, dec...)
		if a.Variadic {
			stmts.Append(
				g.AssignMany(g.List(arg, err),
					jsConsumeRestArgs.Call(parseArgs...)),
			)
		} else {
			stmts.Append(
				g.AssignMany(g.List(arg, err),
					jsConsumeArg.Call(parseArgs...)),
			)
		}
	}

	optArgsBlock := g.StatementList()
	methodPostfix := ""
	for _, a := range gen.Op.Arguments[noOfConsumed:] {
		if a.CustomRule.Ignore {
			continue
		}
		innerBlock := g.StatementList()
		methodPostfix = methodPostfix + internal.UpperCaseFirstLetter(a.Name)

		arg := g.Id(SanitizeVarName(a.Name))
		found := g.Id("found")
		err := g.Id("errArg")
		optArgs = append(optArgs, arg)
		parseArgs := []g.Generator{gen.CbCtx(), g.Lit(a.Name)}
		decoders := DecodersForArg(gen.Receiver(), a)
		parseArgs = append(parseArgs, decoders...)
		optArgsBlock.Append(
			g.AssignMany(
				g.List(arg, found, err),
				jsConsumeOptionalArg.Call(parseArgs...),
			),
			g.IfStmt{
				Condition: found,
				Block: g.StatementList(
					IfErrorF(err, returnNilCommaErr),
					innerBlock,
					callNativeFunc(gen.Op, methodPostfix, append(reqArgs, optArgs...)),
				),
			},
		)
	}

	stmts.Append(
		IfAnyErrorF(errs, returnNilCommaErr),
		optArgsBlock,
		callNativeFunc(gen.Op, "", reqArgs),
	)

	return stmts
}

func (gen OpCallbackMethods) DefaultValuer(a model.ESOperationArgument) (g.Generator, bool) {
	if a.CustomRule.ZeroAsDefault {
		return zeroValue, true
	}
	switch a.IdlArg.Type.Name {
	case "EventInit", "HTMLElement":
		return zeroValue, true
	}

	if d := a.IdlArg.Default; d != nil {
		if d.Type == "null" {
			return zeroValue, true
		}
		if d.Type == "number" && d.Value == "0" {
			return zeroValue, true
		}
		if d.Type == "boolean" && d.Value == false {
			return zeroValue, true
		}
	}
	defaultName, hasDefault := a.DefaultValueInGo()
	if hasDefault && defaultName != "" {
		return gen.Receiver().Field(defaultName), hasDefault
	} else if a.NullableInIDL() {
		return zeroValue, hasDefault
	} else {
		return g.Nil, hasDefault
	}
}
