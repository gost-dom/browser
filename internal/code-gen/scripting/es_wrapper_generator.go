package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/scripting/configuration"
	"github.com/gost-dom/code-gen/scripting/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

// createData combines the specific configuration with the corresponding web IDL
// specification, containins information about the intention about _what_ to
// generate, which methods, etc.
func createData(
	spec idl.Spec,
	interfaceConfig *configuration.WebIDLConfig,
	extra []idl.Spec,
) model.ESConstructorData {
	idlName, ok := spec.GetType(interfaceConfig.TypeName)
	if !ok {
		panic(fmt.Sprintf("cannot find type: %s", interfaceConfig.TypeName))
	}
	idlInterface := idlName.IdlInterface
	if idlInterface.Name != interfaceConfig.TypeName {
		panic(fmt.Sprintf("createData error: %s = %s", idlInterface.Name, interfaceConfig.TypeName))
	}
	for _, e := range extra {
		idlInterface = idlInterface.MergePartials(e)
	}
	specRules := customrules.GetSpecRules(interfaceConfig.DomSpec.Name)
	intfRules := specRules[interfaceConfig.TypeName]
	return model.ESConstructorData{
		Spec:          interfaceConfig,
		CustomRule:    intfRules,
		RunCustomCode: interfaceConfig.RunCustomCode,
		IdlInterface:  idlInterface,
		Constructor:   CreateConstructor(idlInterface, intfRules, interfaceConfig, idlName),
		Operations:    CreateInstanceMethods(idlInterface, intfRules, interfaceConfig),
		Attributes:    CreateAttributes(idlInterface, intfRules, interfaceConfig),
	}
}

func CreateConstructor(
	idlInterface idl.Interface,
	intfRule customrules.InterfaceRule,
	interfaceConfig *configuration.WebIDLConfig,
	idlName idl.TypeSpec) *model.Callback {
	if c, ok := idlName.Constructor(); ok {
		c.Name = "constructor"
		// TODO: Fix for constructor overloads
		result := createOperation(
			model.CallbackKindCtor,
			idl.Operation{
				InterfaceMember: idl.InterfaceMember{Name: "constructor"},
				Arguments:       idlInterface.Constructors[0].Arguments,
			},
			intfRule,
			interfaceConfig,
			false,
		)
		return &result
	} else {
		return nil
	}
}

func CreateInstanceMethods(
	idlInterface idl.Interface,
	intfRule customrules.InterfaceRule,
	interfaceConfig *configuration.WebIDLConfig,
) (result []model.Callback) {
	// TODO: Handle overloads, e.g. of XHR.open
	visited := make(map[string]bool)
	for _, operation := range idlInterface.Operations {
		opRules := intfRule.Operations[operation.Name]
		if opRules.Ignore {
			continue
		}
		if operation.Name != "" && !visited[operation.Name] && !operation.Static {
			result = append(
				result,
				createOperation(
					model.CallbackKindOperation,
					operation,
					intfRule,
					interfaceConfig,
					false,
				),
			)
		}
		if operation.Name == "" && operation.Stringifier {
			operation.ReturnType.Name = "USVString"
			result = append(
				result,
				createOperation(
					model.CallbackKindOperation,
					operation,
					intfRule,
					interfaceConfig,
					true,
				),
			)
		}
		visited[operation.Name] = true
	}
	return
}

func CreateAttributes(
	idlInterface idl.Interface,
	intfRules customrules.InterfaceRule,
	interfaceConfig *configuration.WebIDLConfig,
) (res []model.ESAttribute) {
	for attribute := range idlInterface.AllAttributes(interfaceConfig.IncludeIncludes) {
		methodCustomization := interfaceConfig.GetMethodCustomization(attribute.Name)
		customRule := intfRules.Attributes[attribute.Name]
		if methodCustomization.Ignored || attribute.Type.Name == "EventHandler" ||
			customRule.Ignore {
			continue
		}
		var (
			getter *model.Callback
			setter *model.Callback
		)
		attrType := attribute.Type
		if customRule.OverrideType != nil {
			attrType = customRule.OverrideType.IdlType()
		}
		getter = &model.Callback{
			Name:                 attribute.Name,
			Kind:                 model.CallbackKindGetter,
			NotImplemented:       methodCustomization.NotImplemented,
			CustomImplementation: methodCustomization.CustomImplementation,
			RetType:              attrType,
			MethodCustomization:  methodCustomization,
			ZeroAsNull:           customRule.ZeroAsNull,
		}
		if !attribute.Readonly {
			setter = new(model.Callback)
			*setter = *getter
			setter.Name = fmt.Sprintf("set%s", model.IdlNameToGoName(getter.Name))
			setter.Kind = model.CallbackKindSetter
			methodCustomization := interfaceConfig.GetMethodCustomization(setter.Name)
			setter.NotImplemented = setter.NotImplemented || methodCustomization.NotImplemented
			setter.CustomImplementation = setter.CustomImplementation ||
				methodCustomization.CustomImplementation
			setter.RetType = IdlTypeUndefined
			setter.HasError = customRule.SetterHasError
			setter.Arguments = []model.ESOperationArgument{{
				Name: "val",
				Type: model.IdlNameToGoName(attribute.Type.Name),
				IdlArg: idl.Argument{
					Name: "val",
					Type: attrType,
				},
				Optional: false,
				Variadic: false,
			}}
		}
		getterCustomization := interfaceConfig.GetMethodCustomization(getter.Name)
		getter.NotImplemented = getterCustomization.NotImplemented || getter.NotImplemented
		getter.CustomImplementation = getterCustomization.CustomImplementation ||
			getter.CustomImplementation
		res = append(res, model.ESAttribute{
			Name:   attribute.Name,
			Spec:   attribute,
			Getter: getter,
			Setter: setter})
	}
	return
}

func createOperation(
	kind model.CallbackKind,
	idlOperation idl.Operation,
	intfRules customrules.InterfaceRule,
	typeSpec *configuration.WebIDLConfig,
	stringifier bool,
) model.Callback {
	opRules := intfRules.Operations[idlOperation.Name]
	methodCustomization := typeSpec.GetMethodCustomization(idlOperation.Name)

	op := model.Callback{
		Name:                 idlOperation.Name,
		Spec:                 idlOperation,
		Kind:                 kind,
		NotImplemented:       methodCustomization.NotImplemented,
		CustomImplementation: methodCustomization.CustomImplementation,
		RetType:              idlOperation.ReturnType,
		MethodCustomization:  methodCustomization,
		HasError:             opRules.HasError,
		Arguments:            []model.ESOperationArgument{},
	}
	if stringifier {
		op.Name = "toString"
	}
	for _, idlArg := range idlOperation.Arguments {
		var esArgumentSpec configuration.ESMethodArgument
		if arg := methodCustomization.Argument(idlArg.Name); arg != nil {
			esArgumentSpec = *arg
		}
		esArg := model.ESOperationArgument{
			Name:         idlArg.Name,
			IdlArg:       idlArg,
			Optional:     idlArg.Optional,
			ArgumentSpec: esArgumentSpec,
			Ignore:       esArgumentSpec.Ignored,
			CustomRule:   opRules.Argument(idlArg.Name),
			Variadic:     idlArg.Variadic,
		}
		op.Arguments = append(op.Arguments, esArg)
	}
	return op
}

func IfAnyErrorF(errNames []g.Generator, block TransformerFunc) g.Generator {
	return IfAnyError(errNames, block)
}

func IfAnyError(errNames []g.Generator, block Transformer) g.Generator {
	switch len(errNames) {
	case 0:
		return g.Noop
	case 1:
		return IfError(errNames[0], block)
	default:
		err := g.Id("err")
		return g.StatementList(
			g.Reassign(err, errorsFirst.Call(errNames...)),
			IfError(err, block),
		)
	}
}

func ReturnOnAnyError(errNames []g.Generator) g.Generator {
	switch len(errNames) {
	case 0:
		return g.Noop
	case 1:
		return ReturnIfError(errNames[0])
	default:
		err := g.Id("err")
		return g.StatementList(
			g.Assign(err, stdgen.ErrorsJoin(errNames...)),
			ReturnIfError(err),
		)
	}
}

// SanitizeVarName create a valid go variable name from a variable to avoid
// invalid generated code due to
//
//   - The name is a reserved word, e.g. `type`.
//   - The name already an identifiers in scope (not yet implemented)
func SanitizeVarName(name string) string {
	switch name {
	case "type":
		return "type_"
	case "select":
		return "select_"
	case "interface":
		return "intf"
	}
	return name
}

func ReturnIfError(err g.Generator) g.Generator {
	return IfError(
		err,
		TransformerFunc(func(err g.Generator) g.Generator { return g.Return(g.Nil, err) }),
	)
}

func IfErrorF(err g.Generator, block TransformerFunc) g.Generator { return IfError(err, block) }

func IfError(err g.Generator, block Transformer) g.Generator {
	return g.IfStmt{
		Condition: g.Neq{Lhs: err, Rhs: g.Nil},
		Block:     block.Transform(err),
	}
}

func ReturnTransform(t Transformer) Transformer {
	return TransformerFunc(func(gen g.Generator) g.Generator {
		return g.Return(t.Transform(gen))
	})
}
