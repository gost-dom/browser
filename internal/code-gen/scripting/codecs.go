package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/idltransform"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var decodeString = g.NewValuePackage("DecodeString", packagenames.Codec)
var decodeByteString = g.NewValuePackage("DecodeByteString", packagenames.Codec)
var decodeBoolean = g.NewValuePackage("DecodeBoolean", packagenames.Codec)
var decodeInt = g.NewValuePackage("DecodeInt", packagenames.Codec)

func decode(s string) g.Generator {
	return g.NewValuePackage(fmt.Sprintf("Decode%s", s), packagenames.Codec)
}

func DecodersForArg(receiver g.Generator, arg model.ESOperationArgument) []g.Generator {
	if d := arg.ArgumentSpec.Decoder; d != "" {
		return g.List(g.Id(d))
	}

	argType := arg.IdlArg.Type
	if arg.CustomRule.OverridesType() {
		argType = arg.CustomRule.Type
	}
	return DecodersForType(receiver, argType)
}

// DecodersForType generates the decoders to be used for decoding an input of a
// specific JavaScript type into Go based on the web IDL specification.
func DecodersForType(receiver g.Generator, argType idl.Type) []g.Generator {
	argType = idltransform.FilterType(argType)
	if argType.Kind == idl.KindUnion {
		res := make([]g.Generator, len(argType.Types))
		for i, t := range argType.Types {
			res[i] = decoderForType(receiver, t)
		}
		return res
	} else {
		return []g.Generator{decoderForType(receiver, argType)}
	}
}

func decoderForType(receiver g.Generator, argType idl.Type) g.Generator {
	idlType := idltransform.NewIdlType(argType)
	switch {
	case idlType.Name == "ByteString":
		return decodeByteString
	case idlType.IsString():
		return decodeString
	case idlType.IsBoolean():
		return decodeBoolean
	case idlType.IsInt():
		return decodeInt
	}
	switch argType.Name {
	case "Node", "HTMLElement", "EventInit":
		return decode(argType.Name)
	}
	name := fmt.Sprintf("decode%s", model.IdlNameToGoName(argType.Name))
	return g.ValueOf(receiver).Field(name)
}
