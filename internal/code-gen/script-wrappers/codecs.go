package wrappers

import (
	"fmt"

	"github.com/gost-dom/code-gen/script-wrappers/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

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
		convertNames = []string{fmt.Sprintf("decode%s", model.IdlNameToGoName(arg.GoTypeName()))}
	}

	res := make([]g.Generator, len(convertNames))
	for i, n := range convertNames {
		res[i] = g.ValueOf(receiver).Field(n)
	}
	return res
}
