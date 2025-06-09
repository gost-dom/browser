package v8gen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

var ZeroValue = g.NewValuePackage("ZeroValue", packagenames.Codec)
var ConsumeArgument = g.NewValuePackage("ConsumeArgument", packagenames.JS)
var ConsumeOptionalArg = g.NewValuePackage("ConsumeOptionalArg", packagenames.JS)
var ConsumeRestArguments = g.NewValuePackage("ConsumeRestArguments", packagenames.JS)

// V8CallbackGenerators produces code for a function callback, i.e., when a
// JavaScript function with native code is called. This includes
//
// - Operations
// - Accessor attribute getters and setters
// - Constructors
//
// One instance is coupled to a specific callback, i.e. a specific operation,
// constructor, or accessor attribute
type V8CallbackGenerators struct {
	Data     model.ESConstructorData
	Op       model.ESOperation
	Receiver g.Value
}

func (gen V8CallbackGenerators) instance() g.Value {
	return g.NewValue("instance")
}

func (gen V8CallbackGenerators) errInst() g.Value {
	if len(gen.Op.Arguments) == 0 {
		return g.NewValue("err")
	} else {
		return g.NewValue("errInst")
	}
}

func (gen V8CallbackGenerators) assignInstance(cbCtx scripting.CallbackContext) g.Generator {
	err := gen.errInst()
	return g.StatementList(
		g.AssignMany(
			g.List(gen.instance(), err),
			scripting.As.TypeParam(gen.Data.WrappedType()).Call(cbCtx.GetInstance()),
		),

		scripting.IfError(err, scripting.TransformerFunc(returnNilCommaErr)),
	)
}

func (gen V8CallbackGenerators) OperationCallback(
	cbCtx scripting.CallbackContext,
) g.Generator {
	return g.StatementList(
		gen.assignInstance(cbCtx),
		gen.CtxOrOperationCallback(cbCtx, gen.nativeFunctionCall),
	)
}

func (gen V8CallbackGenerators) ConstructorCallback(cbCtx scripting.CallbackContext) g.Generator {
	return gen.CtxOrOperationCallback(cbCtx, gen.NativeConstructorCall)
}

func (gen V8CallbackGenerators) NativeConstructorCall(
	cbCtx scripting.CallbackContext,
	methodPostFix string,
	args []g.Generator,
) g.Generator {
	return g.StatementList(
		g.Return(
			gen.Receiver.
				Field("CreateInstance" + methodPostFix).
				Call(append([]g.Generator{cbCtx}, args...)...),
		),
	)
}

func (gen V8CallbackGenerators) nativeFunctionCall(
	cbCtx scripting.CallbackContext,
	methodPostFix string,
	args []g.Generator,
) g.Generator {
	name := gen.Op.NativeFunctionName() + methodPostFix
	return gen.transformResult(cbCtx, gen.instance().Field(name).Call(args...))
	// callNativeFunc("", reqArgs)),
}

func (gen V8CallbackGenerators) CtxOrOperationCallback(
	cbCtx scripting.CallbackContext,
	callNativeFunc func(cbCtx scripting.CallbackContext, methodPostFix string, args []g.Generator) g.Generator,
) g.Generator {
	op := gen.Op
	receiver := gen.Receiver

	// reqArgs are "required" args on the Go side. This include optional args in
	// the IDL specification that have a default value, or customization
	// indicates a default value.
	reqArgs := []g.Generator{}
	optArgs := []g.Generator{}

	errs := []g.Generator{}

	stmts := g.StatementList()
	var noOfConsumed int
	for i, a := range op.Arguments {
		noOfConsumed = i
		defaultValuer, hasDefault := gen.DefaultValuer(a)
		if a.Optional && !hasDefault && !a.VariadicInGo() {
			break
		}
		noOfConsumed = i + 1
		if a.Ignore {
			stmts.Append(cbCtx.ConsumeArg())
			continue
		}
		arg := g.Id(scripting.SanitizeVarName(a.Name))
		reqArg := arg
		if a.VariadicInGo() {
			reqArg = g.Raw(jen.Id(a.Name).Op("..."))
		}
		err := g.Id(fmt.Sprintf("errArg%d", i+1))
		reqArgs = append(reqArgs, reqArg)
		errs = append(errs, err)
		parseArgs := []g.Generator{cbCtx, g.Lit(a.Name)}
		if !a.Variadic {
			parseArgs = append(parseArgs, defaultValuer)
		}
		var dec = scripting.DecodersForArg(receiver, a)
		parseArgs = append(parseArgs, dec...)
		if a.Variadic {
			stmts.Append(
				g.AssignMany(g.List(arg, err),
					ConsumeRestArguments.Call(parseArgs...)),
			)
		} else {
			stmts.Append(
				g.AssignMany(g.List(arg, err),
					ConsumeArgument.Call(parseArgs...)),
			)
		}
	}

	optArgsBlock := g.StatementList()
	methodPostfix := ""
	for _, a := range op.Arguments[noOfConsumed:] {
		innerBlock := g.StatementList()
		methodPostfix = methodPostfix + internal.UpperCaseFirstLetter(a.Name)

		arg := g.Id(scripting.SanitizeVarName(a.Name))
		found := g.Id("found")
		err := g.Id("errArg")
		optArgs = append(optArgs, arg)
		parseArgs := []g.Generator{cbCtx, g.Lit(a.Name)}
		decoders := scripting.DecodersForArg(receiver, a)
		parseArgs = append(parseArgs, decoders...)
		optArgsBlock.Append(
			g.AssignMany(
				g.List(arg, found, err),
				ConsumeOptionalArg.Call(parseArgs...),
			),
			g.IfStmt{
				Condition: found,
				Block: g.StatementList(
					scripting.IfErrorF(err, returnNilCommaErr),
					innerBlock,
					callNativeFunc(cbCtx, methodPostfix, append(reqArgs, optArgs...)),
				),
			},
		)
	}

	stmts.Append(
		scripting.IfAnyErrorF(errs, returnNilCommaErr),
		optArgsBlock,
		callNativeFunc(cbCtx, "", reqArgs),
	)

	return stmts
}

func (gen V8CallbackGenerators) AttributeGetterCallback(
	cbCtx scripting.CallbackContext,
	eval scripting.Transformer,
) g.Generator {
	return g.StatementList(
		gen.assignInstance(cbCtx),
		scripting.ReturnValueGenerator{
			V8:       true,
			Data:     gen.Data,
			Op:       gen.Op,
			Ctx:      cbCtx,
			Receiver: gen.Receiver,
		}.Transform(eval.Transform(gen.instance())),
	)
}

func (gen V8CallbackGenerators) DefaultValuer(a model.ESOperationArgument) (g.Generator, bool) {
	switch a.IdlArg.Type.Name {
	case "EventInit", "HTMLElement":
		return ZeroValue, true
	}
	defaultName, hasDefault := a.DefaultValueInGo()
	if hasDefault && defaultName != "" {
		return gen.Receiver.Field(defaultName), hasDefault
	} else if a.NullableInIDL() {
		return ZeroValue, hasDefault
	} else {
		return g.Nil, hasDefault
	}
}

func (gen V8CallbackGenerators) transformResult(
	cbCtx scripting.CallbackContext,
	result g.Generator,
) g.Generator {
	return scripting.ReturnValueGenerator{
		V8:       true,
		Data:     gen.Data,
		Op:       gen.Op,
		Ctx:      cbCtx,
		Receiver: gen.Receiver,
	}.Transform(result)
}
