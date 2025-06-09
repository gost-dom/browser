package model

import (
	"strings"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var encodeString = g.NewValuePackage("EncodeString", packagenames.Codec)
var encodeNillableString = g.NewValuePackage("EncodeNillableString", packagenames.Codec)
var encodeNullableString = g.NewValuePackage("EncodeNullableString", packagenames.Codec)
var encodeInt = g.NewValuePackage("EncodeInt", packagenames.Codec)
var encodeBoolean = g.NewValuePackage("EncodeBoolean", packagenames.Codec)
var encodeEntity = g.NewValuePackage("EncodeEntity", packagenames.Codec)

type ESOperation struct {
	Name                 string
	Spec                 idl.Operation
	NotImplemented       bool
	RetType              idl.Type
	HasError             bool
	CustomImplementation bool
	MethodCustomization  configuration.ESMethodWrapper
	Arguments            []ESOperationArgument
}

// CallbackMethodName gets the name for the unexported function that serves as a
// function callback, i.e. the Go function to be executed when JavaScript code
// calls a native function.
func (o ESOperation) CallbackMethodName() string {
	return idl.SanitizeName(o.Name)
}

// NativeFunctionName gets the name of the method in Go that implements the
// behaviour.
func (o ESOperation) NativeFunctionName() string {
	if o.Name == "toString" {
		return "String"
	}
	return IdlNameToGoName(o.Name)
}

func (op ESOperation) GetHasError() bool {
	return op.HasError
}

func (op ESOperation) HasResult() bool {
	if op.Name == "" {
		return false
	}
	return op.RetType.Name != "undefined"
}

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

func idlTypeNameToGoName(t idl.Type) string {
	switch t.Name {
	case "DOMString", "USVString", "ByteString":
		{
			return "string_"
		}
	default:
		return t.Name
	}
}

func (o ESOperation) EncodeAsSimpleJSLookup() bool {
	if IsNodeType(o.RetType.Name) {
		return true
	}
	switch o.RetType.Name {
	case "Attr", "NodeList", "HTMLFormControlsCollection", "Comment":
		return true
	default:
		return false
	}
}

func (o ESOperation) Encoder(
	receiver g.Value,
	cbCtx g.Generator,
	data ESConstructorData,
) internal.BoundFunction {
	if o.EncodeAsSimpleJSLookup() {
		return internal.BindValues(encodeEntity, cbCtx)
	}
	if e := o.MethodCustomization.Encoder; e != "" {
		return internal.BindValues(receiver.Field(e))
	}
	t := o.RetType
	idlType := idltransform.IdlType(t)
	switch {
	case idlType.IsInt():
		return internal.BindValues(encodeInt, cbCtx)
	case idlType.IsBoolean():
		return internal.BindValues(encodeBoolean, cbCtx)
	case idlType.IsString():
		if t.Nullable {
			if data.CustomRule.OutputType == customrules.OutputTypeStruct {
				return internal.BindValues(encodeNullableString, cbCtx)
			} else {
				return internal.BindValues(encodeNillableString, cbCtx)
			}
		} else {
			return internal.BindValues(encodeString, cbCtx)
		}
	}
	var boundArgs []g.Generator
	converter := "to"
	if t.Kind == idl.KindSequence {
		converter += "Sequence"
		t = *t.TypeParam
	}
	if t.Nullable && !idlType.Nillable() {
		if data.CustomRule.OutputType == customrules.OutputTypeStruct {
			converter += "Nullable"
		} else {
			converter += "Nillable"
			boundArgs = append(boundArgs, cbCtx)
		}
	}
	converter += IdlNameToGoName(idlTypeNameToGoName(t))
	return internal.BindValues(receiver.Field(converter), cbCtx)
}

func (o ESOperation) RetValues(data ESConstructorData) []g.Generator {
	if !o.HasResult() {
		return nil
	}
	t := o.RetType
	res := g.Id("result")
	hasValue := g.Id("hasValue")
	if t.Nullable && !idltransform.IdlType(t).Nillable() {
		if data.CustomRule.OutputType == customrules.OutputTypeStruct {
			return g.List(res)
		} else {
			return g.List(res, hasValue)
		}
	}
	return g.List(res)
}

func (o ESOperation) RetTypeName() string {
	return o.RetType.Name
}
