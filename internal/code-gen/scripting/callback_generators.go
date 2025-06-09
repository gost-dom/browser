package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

type CallbackGenerators struct {
	WrapperStruct
}

func (s CallbackGenerators) Receiver() g.Value      { return g.NewValue("w") }
func (s CallbackGenerators) CbCtx() CallbackContext { return NewCallbackContext(g.Id("cbCtx")) }

func (gen CallbackGenerators) CallbackFunction(name string, body g.Generator) g.Generator {
	return g.FunctionDefinition{
		Receiver: g.FunctionArgument{
			Name: gen.Receiver(),
			Type: gen.WrapperStructType(),
		},
		Name:     name,
		Args:     g.Arg(gen.CbCtx(), v8CbCtx),
		RtnTypes: g.List(JSValue, g.NewType("error")),
		Body: g.StatementList(
			gen.LogCall(name, gen.CbCtx()),
			body,
		),
	}
}

func (s CallbackGenerators) ConstructorCallback() g.Generator {
	return s.CallbackFunction("Constructor", s.ConstructorCallbackBody())
}

func (s CallbackGenerators) ConstructorCallbackBody() g.Generator {
	if !s.Data.AllowConstructor() {
		return g.Return(s.CbCtx().IllegalConstructor())
	}
	cons := *s.Data.Constructor
	return s.OpCallbackGenerators(cons).ConstructorCallbackBody()
}

func (s CallbackGenerators) MethodCallback(op model.ESOperation) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	name := op.CallbackMethodName()
	return s.CallbackFunction(
		name,
		renderIfElse(op.NotImplemented,
			s.ReturnNotImplementedError(name, cbCtx),
			s.OpCallbackGenerators(op).MethodCallbackBody(),
		),
	)
}

func (s CallbackGenerators) OpCallbackGenerators(op model.ESOperation) OpCallbackGenerators {
	return OpCallbackGenerators{s, op}
}

func (s CallbackGenerators) AttributeGetter(attr model.ESAttribute) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	op := attr.Getter
	name := op.CallbackMethodName()
	return s.CallbackFunction(name,
		renderIfElse(op.NotImplemented,
			s.ReturnNotImplementedError(name, cbCtx),
			s.AttributeGetterCallbackBody(attr),
		),
	)
}
func (b CallbackGenerators) AttributeGetterCallbackBody(
	attr model.ESAttribute,
) Generator {
	statements := g.StatementList()
	instance := g.NewValue("instance")
	var call g.Generator
	name := model.IdlNameToGoName(attr.Getter.Name)
	field := g.ValueOf(instance).Field(name)
	if b.Data.CustomRule.OutputType == customrules.OutputTypeStruct {
		call = field
	} else {
		call = field.Call()
	}

	statements.Append(
		b.assignInstance(nil),
		ReturnValueGenerator{
			Data:     b.Data,
			Op:       *attr.Getter,
			Ctx:      b.CbCtx(),
			Receiver: b.Receiver(),
		}.Transform(call),
	)
	return statements
}

func (s CallbackGenerators) AttributeSetter(attr model.ESAttribute) g.Generator {
	cbCtx := NewCallbackContext(g.Id("cbCtx"))
	op := attr.Setter
	name := op.CallbackMethodName()
	return s.CallbackFunction(name,
		renderIfElse(op.NotImplemented,
			s.ReturnNotImplementedError(name, cbCtx),
			s.AttributeSetterCallbackBody(attr),
		),
	)
}

func (gen CallbackGenerators) AttributeSetterCallbackBody(
	attr model.ESAttribute,
) Generator {
	var (
		err      = g.Id("err0")
		err1     = g.Id("err1")
		val      = g.Id("val")
		instance = g.NewValue("instance")
		name     = model.IdlNameToGoName(attr.Setter.Name)
		field    = g.ValueOf(instance).Field(name)
	)

	args := append(
		[]g.Generator{gen.CbCtx()},
		DecodersForArg(gen.Receiver(), attr.Setter.Arguments[0])...,
	)
	parsedArg := parseSetterArg.Call(args...)

	return g.StatementList(
		g.AssignMany(
			g.List(instance, err),
			As.TypeParam(gen.Data.WrappedType()).Call(gen.CbCtx().GetInstance()),
		),
		g.AssignMany(g.List(val, err1), parsedArg),

		IfAnyErrorF(
			[]g.Generator{err, err1},
			returnNilCommaErr,
		),
		renderIfElse(
			gen.Data.CustomRule.OutputType == customrules.OutputTypeStruct,
			g.Reassign(field, val),
			field.Call(val),
		),
		g.Return(g.Nil, g.Nil),
	)
}

func (c CallbackGenerators) ReturnNotImplementedError(
	name string,
	cbCtx CallbackContext,
) g.Generator {
	errMsg := fmt.Sprintf(
		"%s.%s: Not implemented. Create an issue: %s",
		c.Data.Name(), name, packagenames.ISSUE_URL,
	)
	return g.Return(g.Nil, g.NewValuePackage("New", "errors").Call(g.Lit(errMsg)))
}

func (c CallbackGenerators) LogCall(name string, cbCtx g.Generator) g.Generator {
	return g.ValueOf(cbCtx).Field("Logger").Call().Field("Debug").Call(
		g.Lit(fmt.Sprintf("V8 Function call: %s.%s", c.Data.Name(), name)))
}

func (gen CallbackGenerators) assignInstance(
	args []model.ESOperationArgument,
) g.Generator {
	err := gen.errInst(args)
	return g.StatementList(
		g.AssignMany(
			g.List(gen.instance(), err),
			As.TypeParam(gen.Data.WrappedType()).Call(gen.CbCtx().GetInstance()),
		),

		IfError(err, TransformerFunc(returnNilCommaErr)),
	)
}

func (gen CallbackGenerators) instance() g.Value {
	return g.NewValue("instance")
}

func (gen CallbackGenerators) errInst(args []model.ESOperationArgument) g.Value {
	if len(args) == 0 {
		return g.NewValue("err")
	} else {
		return g.NewValue("errInst")
	}
}
