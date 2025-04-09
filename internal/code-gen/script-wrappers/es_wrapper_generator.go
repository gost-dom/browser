package wrappers

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/gost-dom/code-gen/customrules"
	. "github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	. "github.com/gost-dom/code-gen/script-wrappers/model"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"

	"github.com/dave/jennifer/jen"
)

var (
	v8FunctionTemplatePtr     = g.NewTypePackage("FunctionTemplate", v8).Pointer()
	v8FunctionCallbackInfoPtr = g.NewTypePackage("FunctionCallbackInfo", v8).Pointer()
	v8ObjectTemplatePtr       = g.NewTypePackage("ObjectTemplate", v8).Pointer()
	v8Value                   = g.NewTypePackage("Value", v8).Pointer()
	v8ReadOnly                = g.Raw(jen.Qual(v8, "ReadOnly"))
	v8None                    = g.Raw(jen.Qual(v8, "None"))
	scriptHostPtr             = g.NewType("V8ScriptHost").Pointer()
)

const (
	v8      = "github.com/gost-dom/v8go"
	gojaSrc = "github.com/dop251/goja"
)

func createData(
	spec idl.Spec,
	interfaceConfig *configuration.IdlInterfaceConfiguration,
) ESConstructorData {
	idlName, ok := spec.GetType(interfaceConfig.TypeName)
	if !ok {
		panic("Missing type")
	}
	idlInterface := idlName.IdlInterface
	wrappedTypeName := idlInterface.Name
	return ESConstructorData{
		Spec:             interfaceConfig,
		IdlInterfaceName: wrappedTypeName,
		RunCustomCode:    interfaceConfig.RunCustomCode,
		Inheritance:      idlInterface.Inheritance,
		IdlInterface:     idlInterface,
		Constructor:      CreateConstructor(interfaceConfig, idlName),
		Operations:       CreateInstanceMethods(interfaceConfig, idlName),
		Attributes:       CreateAttributes(interfaceConfig, idlName),
	}
}

func CreateConstructor(
	interfaceConfig *configuration.IdlInterfaceConfiguration,
	idlName idl.TypeSpec) *ESOperation {
	if c, ok := idlName.Constructor(); ok {
		fmt.Printf("Create constructor %s '%s'\n", interfaceConfig.TypeName, c.Name)
		c.Name = "constructor"
		result := createOperation(interfaceConfig, c)
		return &result
	} else {
		return nil
	}
}

func CreateInstanceMethods(
	interfaceConfig *configuration.IdlInterfaceConfiguration,
	idlName idl.TypeSpec) (result []ESOperation) {
	for instanceMethod := range idlName.InstanceMethods() {
		op := createOperation(interfaceConfig, instanceMethod)
		result = append(result, op)
	}
	return
}

func CreateAttributes(
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
		// r := attribute.AttributeType()
		// rtnType := r.TypeName
		getter = &ESOperation{
			Name:                 attribute.Name,
			NotImplemented:       methodCustomization.NotImplemented,
			CustomImplementation: methodCustomization.CustomImplementation,
			RetType: idl.RetType{
				TypeName: attribute.Type.Name,
				Nullable: attribute.Type.Nullable,
			},
			MethodCustomization: methodCustomization,
		}
		if !attribute.Readonly {
			setter = new(ESOperation)
			*setter = *getter
			setter.Name = fmt.Sprintf("set%s", idlNameToGoName(getter.Name))
			methodCustomization := interfaceConfig.GetMethodCustomization(setter.Name)
			setter.NotImplemented = setter.NotImplemented || methodCustomization.NotImplemented
			setter.CustomImplementation = setter.CustomImplementation ||
				methodCustomization.CustomImplementation
			setter.RetType = idl.NewRetTypeUndefined()
			setter.Arguments = []ESOperationArgument{{
				Name:     "val",
				Type:     idlNameToGoName(attribute.Type.Name),
				Optional: false,
				Variadic: false,
			}}
		}
		getterCustomization := interfaceConfig.GetMethodCustomization(getter.Name)
		getter.NotImplemented = getterCustomization.NotImplemented || getter.NotImplemented
		getter.CustomImplementation = getterCustomization.CustomImplementation ||
			getter.CustomImplementation
		res = append(res, ESAttribute{attribute.Name, getter, setter})
	}
	return
}

func createOperation(
	typeSpec *configuration.IdlInterfaceConfiguration,
	member idl.MemberSpec,
) ESOperation {
	specRules := customrules.GetSpecRules(typeSpec.DomSpec.Name)
	intfRules := specRules[typeSpec.TypeName]
	opRules := intfRules.Operations[member.Name]

	methodCustomization := typeSpec.GetMethodCustomization(member.Name)
	op := ESOperation{
		Name:                 member.Name,
		NotImplemented:       methodCustomization.NotImplemented,
		CustomImplementation: methodCustomization.CustomImplementation,
		RetType:              member.ReturnType(),
		MethodCustomization:  methodCustomization,
		HasError:             opRules.HasError,
		Arguments:            []ESOperationArgument{},
	}
	for _, arg := range member.Arguments {
		var esArgumentSpec configuration.ESMethodArgument
		if arg := methodCustomization.Argument(arg.Name); arg != nil {
			esArgumentSpec = *arg
		}
		esArg := ESOperationArgument{
			Name:         arg.Name,
			Optional:     arg.Optional && !esArgumentSpec.Required,
			IdlType:      arg.IdlType,
			ArgumentSpec: esArgumentSpec,
			Ignore:       esArgumentSpec.Ignored,
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
	if len(errNames) == 0 {
		return g.Noop
	}
	if len(errNames) == 1 {
		return ReturnOnError{err: errNames[0]}
	}
	return g.StatementList(
		g.Assign(g.Id("err"),
			stdgen.ErrorsJoin(errNames...),
		),
		ReturnOnError{},
	)
}

type JenGenerator = g.Generator

func IsNodeType(typeName string) bool {
	loweredName := strings.ToLower(typeName)
	switch loweredName {
	case "node":
		return true
	case "document":
		return true
	case "documentfragment":
		return true
	}
	if strings.HasSuffix(loweredName, "element") {
		return true
	}
	return false
}

// sanitizeVarName create a valid go variable name from a variable to avoid
// invalid generated code due to
//
//   - The name is a reserved word, e.g. `type`.
//   - The name already an identifiers in scope (not yet implemented)
func sanitizeVarName(name string) string {
	switch name {
	case "type":
		return "type_"
	case "select":
		return "select_"
	}
	return name
}

func idlNameToGoName(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = UpperCaseFirstLetter(word)
	}
	return strings.Join(words, "")
}

func idlNameToUnexportedGoName(s string) string {
	return LowerCaseFirstLetter(idlNameToGoName(s))
}

type ReturnOnError struct {
	err g.Generator
}

func (ret ReturnOnError) Generate() *jen.Statement {
	err := ret.err
	if err == nil {
		err = g.Id("err")
	}
	return g.IfStmt{
		Condition: g.Neq{Lhs: err, Rhs: g.Nil}, //g.Raw(err.Generate().Op("!=").Nil()),
		Block:     g.Return(g.Nil, err),
	}.Generate()
}
