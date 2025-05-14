package wrappers

import (
	"fmt"
	"log/slog"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func createData(
	spec idl.Spec,
	interfaceConfig *configuration.IdlInterfaceConfiguration,
) ESConstructorData {
	idlName, ok := spec.GetType(interfaceConfig.TypeName)
	if !ok {
		panic("Missing type")
	}
	// if interfaceConfig.IdlInterface.Name != interfaceConfig.TypeName {
	// 	panic(
	// 		fmt.Sprintf(
	// 			"Type mismatch, %s = %s",
	// 			interfaceConfig.IdlInterface.Name,
	// 			interfaceConfig.TypeName,
	// 		),
	// 	)
	// }
	idlInterface := idlName.IdlInterface
	wrappedTypeName := idlInterface.Name
	if idlInterface.Name != interfaceConfig.TypeName {
		panic(fmt.Sprintf("createData error: %s = %s", idlInterface.Name, interfaceConfig.TypeName))
	}
	return ESConstructorData{
		Spec:             interfaceConfig,
		IdlInterfaceName: wrappedTypeName,
		RunCustomCode:    interfaceConfig.RunCustomCode,
		Inheritance:      idlInterface.Inheritance,
		IdlInterface:     idlInterface,
		Constructor:      CreateConstructor(idlInterface, interfaceConfig, idlName),
		Operations:       CreateInstanceMethods(idlInterface, interfaceConfig, idlName),
		Attributes:       CreateAttributes(idlInterface, interfaceConfig, idlName),
	}
}

func CreateConstructor(
	idlInterface idl.Interface,
	interfaceConfig *configuration.IdlInterfaceConfiguration,
	idlName idl.TypeSpec) *ESOperation {
	if c, ok := idlName.Constructor(); ok {
		fmt.Printf("Create constructor %s '%s'\n", interfaceConfig.TypeName, c.Name)
		c.Name = "constructor"
		// TODO: Fix for constructor overloads
		result := createOperation(
			idlInterface,
			interfaceConfig,
			c,
			idlInterface.Constructors[0].Arguments,
		)
		return &result
	} else {
		return nil
	}
}

func CreateInstanceMethods(
	idlInterface idl.Interface,
	interfaceConfig *configuration.IdlInterfaceConfiguration,
	idlName idl.TypeSpec) (result []ESOperation) {
	for instanceMethod := range idlName.InstanceMethods() {
		idlOperation, found := idlInterface.GetOperation(instanceMethod.Name)
		if !found {
			panic("Method not found: " + instanceMethod.Name)
		}

		op := createOperation(idlInterface, interfaceConfig, instanceMethod, idlOperation.Arguments)
		result = append(result, op)
	}
	return
}

func CreateAttributes(
	idlInterface idl.Interface,
	interfaceConfig *configuration.IdlInterfaceConfiguration,
	idlName idl.TypeSpec,
) (res []ESAttribute) {
	for attribute := range idlName.IdlInterface.AllAttributes(interfaceConfig.IncludeIncludes) {
		methodCustomization := interfaceConfig.GetMethodCustomization(attribute.Name)
		if methodCustomization.Ignored || attribute.Type.Name == "EventHandler" {
			continue
		}
		var (
			getter *ESOperation
			setter *ESOperation
		)
		getter = &ESOperation{
			Name:                 attribute.Name,
			NotImplemented:       methodCustomization.NotImplemented,
			CustomImplementation: methodCustomization.CustomImplementation,
			RetType:              attribute.Type,
			MethodCustomization:  methodCustomization,
		}
		if !attribute.Readonly {
			setter = new(ESOperation)
			*setter = *getter
			setter.Name = fmt.Sprintf("set%s", IdlNameToGoName(getter.Name))
			methodCustomization := interfaceConfig.GetMethodCustomization(setter.Name)
			setter.NotImplemented = setter.NotImplemented || methodCustomization.NotImplemented
			setter.CustomImplementation = setter.CustomImplementation ||
				methodCustomization.CustomImplementation
			setter.RetType = IdlTypeUndefined
			setter.Arguments = []ESOperationArgument{{
				Name:     "val",
				Type:     IdlNameToGoName(attribute.Type.Name),
				Optional: false,
				Variadic: false,
			}}
		}
		getterCustomization := interfaceConfig.GetMethodCustomization(getter.Name)
		getter.NotImplemented = getterCustomization.NotImplemented || getter.NotImplemented
		getter.CustomImplementation = getterCustomization.CustomImplementation ||
			getter.CustomImplementation
		res = append(res, ESAttribute{Name: attribute.Name, Getter: getter, Setter: setter})
	}
	return
}

func createOperation(
	idlInterface idl.Interface,
	typeSpec *configuration.IdlInterfaceConfiguration,
	member idl.MemberSpec,
	idlArgs []idl.Argument,
) ESOperation {
	specRules := customrules.GetSpecRules(typeSpec.DomSpec.Name)
	intfRules := specRules[typeSpec.TypeName]
	opRules := intfRules.Operations[member.Name]
	methodCustomization := typeSpec.GetMethodCustomization(member.Name)
	idlOperation, _ := idlInterface.GetOperation(member.Name)

	op := ESOperation{
		Name:                 member.Name,
		NotImplemented:       methodCustomization.NotImplemented,
		CustomImplementation: methodCustomization.CustomImplementation,
		RetType:              idlOperation.ReturnType,
		MethodCustomization:  methodCustomization,
		HasError:             opRules.HasError,
		Arguments:            []ESOperationArgument{},
	}
	for i, arg := range member.Arguments {
		var esArgumentSpec configuration.ESMethodArgument
		if arg := methodCustomization.Argument(arg.Name); arg != nil {
			esArgumentSpec = *arg
		}
		esArg := ESOperationArgument{
			Name:         arg.Name,
			IdlArg:       idlArgs[i],
			Optional:     arg.Optional && !esArgumentSpec.Required,
			IdlType:      arg.IdlType,
			ArgumentSpec: esArgumentSpec,
			Ignore:       esArgumentSpec.Ignored,
			CustomRule:   opRules.Argument(arg.Name),
		}
		if len(arg.IdlType.Types) > 0 {
			slog.Warn(
				"Multiple argument types",
				"Operation",
				member.Name,
				"Argument",
				arg.Name,
			)
		}
		if arg.IdlType.IdlType != nil {
			esArg.Type = arg.IdlType.IdlType.IType.TypeName
		}
		op.Arguments = append(op.Arguments, esArg)
	}
	return op
}

func ReturnOnAnyError(errNames []g.Generator) g.Generator {
	switch len(errNames) {
	case 0:
		return g.Noop
	case 1:
		return returnIfError(errNames[0])
	default:
		err := g.Id("err")
		return g.StatementList(
			g.Assign(err, stdgen.ErrorsJoin(errNames...)),
			returnIfError(err),
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
	}
	return name
}

func returnIfError(err g.Generator) g.Generator {
	return g.IfStmt{
		Condition: g.Neq{Lhs: err, Rhs: g.Nil},
		Block:     g.Return(g.Nil, err),
	}
}
