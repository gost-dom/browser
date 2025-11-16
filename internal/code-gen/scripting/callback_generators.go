package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

type CallbackMethods struct{ WrapperStruct }

func (cb CallbackMethods) Receiver() g.Value      { return g.NewValue("w") }
func (cb CallbackMethods) CbCtx() CallbackContext { return NewCallbackContext(g.Id("cbCtx")) }

func (cb CallbackMethods) CallbackFunction(name string, body g.Generator) g.Generator {
	return g.FunctionDefinition{
		Receiver: g.FunctionArgument{
			Name: cb.Receiver(),
			Type: cb.WrapperStructType(),
		},
		Name: name,
		Args: g.Arg(cb.CbCtx(), jsCbCtx),
		RtnTypes: g.List(
			g.Raw(jen.Id("res").Add(jsValue.Generate())),
			g.Raw(jen.Id("err").Add(g.NewType("error").Generate())),
		),
		Body: g.StatementList(
			cb.LogCall(name, cb.CbCtx()),
			body,
		),
	}
}

func (cb CallbackMethods) ConstructorCallback() g.Generator {
	return cb.CallbackFunction("Constructor", cb.ConstructorCallbackBody())
}

func (cb CallbackMethods) ConstructorCallbackBody() g.Generator {
	if !cb.Data.AllowConstructor() {
		return g.Return(cb.CbCtx().IllegalConstructor())
	}
	cons := *cb.Data.Constructor
	return cb.OpCallbackGenerators(cons).ConstructorCallbackBody()
}

func (cb CallbackMethods) MethodCallback(op model.ESOperation) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	name := op.CallbackMethodName()
	return cb.CallbackFunction(
		name,
		renderIfElse(op.NotImplemented,
			cb.ReturnNotImplementedError(name, cbCtx),
			cb.OpCallbackGenerators(op).MethodCallbackBody(),
		),
	)
}

func (cb CallbackMethods) OpCallbackGenerators(op model.ESOperation) OpCallbackMethods {
	return OpCallbackMethods{cb, op}
}

func (cb CallbackMethods) AttributeGetter(attr model.ESAttribute) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	op := attr.Getter
	name := op.CallbackMethodName()
	return cb.CallbackFunction(name,
		renderIfElse(op.NotImplemented,
			cb.ReturnNotImplementedError(name, cbCtx),
			cb.AttributeGetterCallbackBody(attr),
		),
	)
}

func (cb CallbackMethods) AttributeGetterCallbackBody(
	attr model.ESAttribute,
) g.Generator {
	statements := g.StatementList()
	instance := g.NewValue("instance")
	var call g.Generator
	name := model.IdlNameToGoName(attr.Getter.Name)
	field := g.ValueOf(instance).Field(name)
	if cb.Data.CustomRule.OutputType == customrules.OutputTypeStruct {
		call = field
	} else {
		call = field.Call()
	}

	statements.Append(
		cb.assignInstance(nil),
		ReturnValueGenerator{
			Data:     cb.Data,
			Op:       *attr.Getter,
			Ctx:      cb.CbCtx(),
			Receiver: cb.Receiver(),
		}.Transform(call),
	)
	return statements
}

func (cb CallbackMethods) AttributeSetter(attr model.ESAttribute) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	op := attr.Setter
	name := op.CallbackMethodName()
	return cb.CallbackFunction(name,
		renderIfElse(op.NotImplemented,
			cb.ReturnNotImplementedError(name, cbCtx),
			cb.AttributeSetterCallbackBody(attr),
		),
	)
}

func (cb CallbackMethods) AttributeSetterCallbackBody(attr model.ESAttribute) g.Generator {
	var (
		err      = g.Id("err0")
		err1     = g.Id("err1")
		val      = g.Id("val")
		instance = g.NewValue("instance")
		name     = model.IdlNameToGoName(attr.Setter.Name)
		field    = g.ValueOf(instance).Field(name)
	)

	args := append(
		[]g.Generator{cb.CbCtx()},
		DecodersForArg(cb.Receiver(), attr.Setter.Arguments[0])...,
	)
	parsedArg := jsParseSetterArg.Call(args...)

	return g.StatementList(
		g.AssignMany(
			g.List(instance, err),
			jsAs.TypeParam(cb.Data.WrappedType()).Call(cb.CbCtx().GetInstance()),
		),
		g.AssignMany(g.List(val, err1), parsedArg),

		IfAnyErrorF(
			[]g.Generator{err, err1},
			returnNilCommaErr,
		),
		renderIfElse(
			cb.Data.CustomRule.OutputType == customrules.OutputTypeStruct,
			g.Reassign(field, val),
			field.Call(val),
		),
		g.Return(g.Nil, g.Nil),
	)
}

func (cb CallbackMethods) ReturnNotImplementedError(
	name string,
	cbCtx CallbackContext,
) g.Generator {
	errMsg := g.Lit(fmt.Sprintf(
		"%s.%s: Not implemented. Create an issue: %s",
		cb.Data.Name(), name, packagenames.ISSUE_URL,
	))
	return g.Return(
		EncodeCallbackErrorf.Call(cbCtx, errMsg),
	)
}

func (cb CallbackMethods) LogCall(name string, cbCtx g.Generator) g.Generator {
	res := g.ValueOf(cbCtx).Field("Logger").Call().Field("Debug").Call(
		g.Lit(fmt.Sprintf("JS Function call: %s.%s", cb.Data.Name(), name)),
		jsThisLogAttr.Call(cbCtx),
		jsLogAttr.Call(g.Lit("res"), g.Id("res")),
	)
	f := g.ValueOf(g.FunctionDefinition{Body: res})
	return g.Raw(jen.Defer().Add(f.Call().Generate()))
}

func (cb CallbackMethods) assignInstance(
	args []model.ESOperationArgument,
) g.Generator {
	err := cb.errInst(args)
	return g.StatementList(
		g.AssignMany(
			g.List(cb.instance(), err),
			jsAs.TypeParam(cb.Data.WrappedType()).Call(cb.CbCtx().GetInstance()),
		),

		IfError(err, TransformerFunc(returnNilCommaErr)),
	)
}

func (cb CallbackMethods) instance() g.Value {
	return g.NewValue("instance")
}

func (cb CallbackMethods) errInst(args []model.ESOperationArgument) g.Value {
	if len(args) == 0 {
		return g.NewValue("err")
	} else {
		return g.NewValue("errInst")
	}
}
