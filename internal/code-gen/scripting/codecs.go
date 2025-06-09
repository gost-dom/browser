package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

var decodeString = g.NewValuePackage("DecodeString", packagenames.Codec)
var decodeBoolean = g.NewValuePackage("DecodeBoolean", packagenames.Codec)
var decodeInt = g.NewValuePackage("DecodeInt", packagenames.Codec)
var decodeNode = g.NewValuePackage("DecodeNode", packagenames.Codec)
var decodeHTMLElement = g.NewValuePackage("DecodeHTMLElement", packagenames.Codec)

func decode(s string) g.Generator {
	return g.NewValuePackage(fmt.Sprintf("Decode%s", s), packagenames.Codec)
}

func DecodersForArg(receiver g.Generator, arg model.ESOperationArgument) []g.Generator {
	var convertNames []string

	if d := arg.ArgumentSpec.Decoder; d != "" {
		return g.List(g.Id(d))
	}

	argType := arg.IdlArg.Type
	if argType.Kind == idl.KindUnion {
		convertNames = make([]string, len(argType.Types))
		for i, t := range argType.Types {
			convertNames[i] = fmt.Sprintf("decode%s", model.IdlNameToGoName(t.Name))
		}
	} else {
		switch {
		case arg.CustomRule.OverridesType():
		case arg.IsString():
			return g.List(decodeString)
		case arg.IsBoolean():
			return g.List(decodeBoolean)
		case arg.IsInt():
			return g.List(decodeInt)
		}
		switch arg.GoTypeName() {
		case "Node", "HTMLElement", "EventInit":
			return g.List(decode(arg.GoTypeName()))
		}
		convertNames = []string{fmt.Sprintf("decode%s", model.IdlNameToGoName(arg.GoTypeName()))}
	}

	res := make([]g.Generator, len(convertNames))
	for i, n := range convertNames {
		res[i] = g.ValueOf(receiver).Field(n)
	}
	return res
}
