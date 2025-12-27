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
var encodeOptionalString = g.NewValuePackage("EncodeOptionalString", packagenames.Codec)
var encodeInt = g.NewValuePackage("EncodeInt", packagenames.Codec)
var encodeBoolean = g.NewValuePackage("EncodeBoolean", packagenames.Codec)
var encodeEntity = g.NewValuePackage("EncodeEntity", packagenames.Codec)

type CallbackKind int

const (
	CallbackKindCtor CallbackKind = iota
	CallbackKindOperation
	CallbackKindGetter
	CallbackKindSetter
)

// Callback represents a Go function callback implementing a JS function.
type Callback struct {
	Name                 string
	Spec                 idl.Operation
	Kind                 CallbackKind
	NotImplemented       bool
	RetType              idl.Type
	HasError             bool
	CustomImplementation bool
	MethodCustomization  configuration.ESMethodWrapper
	Arguments            []ESOperationArgument
	ZeroAsNull           bool
}

// CallbackMethodName gets the name for the unexported function that serves as a
// function callback, i.e. the Go function to be executed when JavaScript code
// calls a native function.
func (o Callback) CallbackMethodName() string {
	return idl.SanitizeName(o.Name)
}

// NativeFunctionName gets the name of the method in Go that implements the
// behaviour.
func (o Callback) NativeFunctionName() string {
	if o.Name == "toString" {
		return "String"
	}
	return IdlNameToGoName(o.Name)
}

func (op Callback) GetHasError() bool {
	return op.HasError
}

func (op Callback) HasResult() bool {
	if op.Name == "" {
		return false
	}
	return op.RetType.Name != "undefined"
}

func IsNodeType(typeName string) bool {
	loweredName := strings.ToLower(typeName)
	switch loweredName {
	case "node", "text", "cdata", "comment", "processinginstruction":
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

func (o Callback) EncodeAsSimpleJSLookup() bool {
	if IsNodeType(o.RetType.Name) {
		return true
	}
	switch o.RetType.Name {
	case "Attr", "NodeList", "HTMLFormControlsCollection", "Comment", "DOMStringMap", "Location":
		return true
	default:
		return false
	}
}

func zeroAsNull(_ ESConstructorData, cb Callback) bool { return cb.ZeroAsNull }

func hasStringOkReturn(data ESConstructorData, cb Callback) bool {
	return data.CustomRule.OutputType == customrules.OutputTypeStruct &&
		cb.Kind == CallbackKindGetter
}

func (o Callback) Encoder(
	receiver g.Value,
	cbCtx g.Generator,
	data ESConstructorData,
) internal.BoundFunction {
	if o.EncodeAsSimpleJSLookup() {
		return internal.BindValues(encodeEntity, cbCtx)
	}
	t := idltransform.FilterType(o.RetType)
	idlType := idltransform.NewIdlType(t)
	switch {
	case idlType.IsInt():
		return internal.BindValues(encodeInt, cbCtx)
	case idlType.IsBoolean():
		return internal.BindValues(encodeBoolean, cbCtx)
	case idlType.IsString():
		if t.Nullable {
			if zeroAsNull(data, o) {
				return internal.BindValues(encodeOptionalString, cbCtx)
			} else if hasStringOkReturn(data, o) {
				return internal.BindValues(encodeNullableString, cbCtx)
			} else {
				return internal.BindValues(encodeNillableString, cbCtx)
			}
		} else {
			return internal.BindValues(encodeString, cbCtx)
		}
	}
	var boundArgs []g.Generator
	converter := "encode"
	if t.Kind == idl.KindSequence {
		converter += "Sequence"
		t = *t.TypeParam
	}
	if t.Kind == idl.KindPromise {
		converter += "Promise"
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
	return internal.BindValues(g.NewValue(converter), cbCtx)
}

func (o Callback) RetValues(data ESConstructorData) []g.Generator {
	if !o.HasResult() {
		return nil
	}
	t := o.RetType
	res := g.Id("result")
	hasValue := g.Id("hasValue")
	if t.Nullable && !idltransform.NewIdlType(t).Nillable() && !o.ZeroAsNull {
		if hasStringOkReturn(data, o) {
			return g.List(res)
		} else {
			return g.List(res, hasValue)
		}
	}
	return g.List(res)
}

func (o Callback) RetTypeName() string {
	return o.RetType.Name
}
