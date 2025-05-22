package wrappers

import (
	"fmt"

	"github.com/gost-dom/code-gen/script-wrappers/model"
	g "github.com/gost-dom/generators"
)

func DecodersForArg(receiver g.Generator, arg model.ESOperationArgument) []g.Generator {
	var convertNames []string
	if arg.Type != "" {
		convertNames = []string{fmt.Sprintf("decode%s", model.IdlNameToGoName(arg.Type))}
	} else {
		types := arg.IdlType.IdlType.IType.Types
		convertNames = make([]string, len(types))
		for i, t := range types {
			convertNames[i] = fmt.Sprintf("decode%s", t.IType.TypeName)
		}
	}
	res := make([]g.Generator, len(convertNames))
	for i, n := range convertNames {
		res[i] = g.ValueOf(receiver).Field(n)
	}
	return res
}
