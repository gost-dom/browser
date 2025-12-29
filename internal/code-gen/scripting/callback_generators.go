package scripting

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/customrules"
	variable "github.com/gost-dom/code-gen/gen/var"
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
		Body: body,
	}
}

func (cb CallbackMethods) ConstructorCallback() g.Generator {
	return cb.CallbackFunction("Constructor", cb.ConstructorCallbackBody())
}

func (cb CallbackMethods) ConstructorCallbackBody() g.Generator {
	if cb.Data.IsEventType() {
		return cb.EventConstructorCallbackBody()
	}
	if !cb.Data.AllowConstructor() {
		return g.Return(cb.CbCtx().IllegalConstructor())
	}
	cons := *cb.Data.Constructor
	return cb.OpCallbackGenerators(cons).ConstructorCallbackBody()
}

func (cb CallbackMethods) MethodCallback(op model.Callback) g.Generator {
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

func EventInitDictType(name, domspec string) g.Generator {
	switch name {
	case "EventInit":
		return g.NewTypePackage(name, packagenames.Events)
	}
	return g.NewTypePackage(name, packagenames.PackageName(domspec))
}

func (m CallbackMethods) EventConstructorCallbackBody() g.Generator {
	type_ := g.Id("type_")
	err := g.Id("err")
	errType := g.Id("errType")
	errOpts := g.Id("errOpts")
	ev := g.NewValue("e")
	options := g.Id("options")
	data := g.Id("data")
	cons := *m.Data.Constructor

	fmt.Printf("Constructor: %+v", cons.Arguments[1])
	eventInitType := cons.Arguments[1].IdlArg.Type.Name // "KeyboardEventInit"
	goEventInitType := EventInitDictType(eventInitType, m.SpecName())

	return g.StatementList(
		g.AssignMany(g.List(type_, errType),
			jsConsumeArg.Call(m.CbCtx(), g.Lit("type"), g.Nil, decodeString),
		),
		g.AssignMany(
			g.List(options, errOpts),
			jsConsumeArg.Call(m.CbCtx(), g.Lit("options"), zeroValue, decodeJsObject),
		),
		g.Reassign(err, errorsFirst.Call(errType, errOpts)),
		g.IfStmt{
			Condition: g.Neq{Lhs: err, Rhs: g.Nil},
			Block:     g.Return(g.Nil, err),
		},

		renderIf(m.Data.Name() != "Event",
			variable.New(
				variable.Name(data),
				variable.Type(goEventInitType),
			),
		),
		g.Assign(ev, StructLiteral(event,
			KeyField("Type", type_),
		)),
		g.IfStmt{
			Condition: g.Neq{Lhs: options, Rhs: g.Nil},
			Block: g.StatementList(
				g.Reassign(err, g.NewValuePackage("DecodeEvent", packagenames.Codec).Call(
					m.CbCtx(),
					options,
					g.ValueOf(ev).Reference(),
				)),
				ReturnIfError(err),
				renderIf(m.Data.Name() != "Event",
					g.StatementList(
						g.Reassign(err, g.NewValue(fmt.Sprintf("decode%s", eventInitType)).Call(
							m.CbCtx(),
							options,
							g.ValueOf(data).Reference(),
						)),
						ReturnIfError(err),
					)),
			),
		},
		renderIf(m.Data.Name() != "Event",
			g.Reassign(ev.Field("Data"), data),
		),
		g.Return(EncodeConstructedValue(m.CbCtx(), ev.Reference())),
	)
}

func StructLiteral(t g.Generator, opts ...func(*g.StructLiteral)) g.Generator {
	res := g.StructLiteral{Type: t}
	for _, opt := range opts {
		opt(&res)
	}
	return res
}

func KeyField(name string, val g.Generator) func(*g.StructLiteral) {
	return func(l *g.StructLiteral) { l.KeyField(g.Id(name), val) }
}

func (cb CallbackMethods) OpCallbackGenerators(op model.Callback) OpCallbackMethods {
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
	attrRule := cb.Data.CustomRule.Attributes[attr.Name]
	if cb.Data.CustomRule.OutputType == customrules.OutputTypeStruct && !attrRule.Callable {
		targetTypeRule := customrules.AllRules[attr.Spec.Type.Name]
		if targetTypeRule.OutputType == customrules.OutputTypeStruct {
			call = field.Reference()
		} else {
			call = field
		}
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
		DecodersForType(cb.Receiver(), attr.Spec.Type)...,
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
			g.StatementList(
				g.Reassign(field, val),
				g.Return(g.Nil, g.Nil),
			),
			renderIfElse(
				attr.Setter.HasError,
				g.StatementList(
					g.Return(g.Nil, field.Call(val)),
				),
				g.StatementList(
					field.Call(val),
					g.Return(g.Nil, g.Nil),
				),
			),
		),
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
	c := CallbackContext{g.ValueOf(cbCtx)}
	log := Logger{g.NewValue("l")}
	return g.StatementList(
		g.Assign(log, c.Logger().With(
			slogString.Call(g.Lit("IdlInterface"), g.Lit(cb.IdlName())),
			slogString.Call(g.Lit("Method"), g.Lit(name)),
		)),
		log.Debug("JS function callback enter",
			jsThisLogAttr.Call(cbCtx),
			jsArgsLogAttr.Call(cbCtx),
		),
		g.Raw(
			jen.Defer().
				Add(g.ValueOf(g.FunctionDefinition{Body: log.Debug("JS function callback exit",
					jsLogAttr.Call(g.Lit("res"), g.Id("res")),
					logErrAttr.Call(g.Id("err")),
				)}).Call().Generate()),
		),
	)
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
