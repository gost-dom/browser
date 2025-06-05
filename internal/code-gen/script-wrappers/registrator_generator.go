package wrappers

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func Write(specs configuration.WebIdlConfigurations) error {
	statements := g.StatementList()
	registrator := g.Id("reg")
	s := slices.Collect(maps.Values(specs))
	slices.SortFunc(
		s,
		func(x, y *configuration.WebIdlConfiguration) int { return cmp.Compare(x.Name, y.Name) },
	)
	for _, spec := range s {
		data, err := idl.Load(spec.Name)
		if err != nil {
			return err
		}
		types := spec.GetTypesSorted()
		// errs := make([]error, len(types))
		for _, specType := range types {
			typeInfo := createData(data, specType)
			statements.Append(
				JSRegister.Call(
					registrator,
					g.Lit(typeInfo.Name()),
					g.Lit(typeInfo.Extends()),
					g.Id(fmt.Sprintf("new%sV8Wrapper", typeInfo.Name())),
				))
		}
	}

	bootstrap := g.Raw(jen.Func().Id("Bootstrap").Types(jen.Id("T").Any()).Params(
		jen.Add(registrator.Generate()).
			Add(JSClassBiulder.Generate()).Types(jen.Id("T")),
	).Block(statements.Generate()))

	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}

	return writeGenerator(writer, packagenames.V8host, bootstrap)
}

func GenerateRegisterFunctions() error {
	specs := configuration.CreateV8Specs()
	return Write(specs)
}
