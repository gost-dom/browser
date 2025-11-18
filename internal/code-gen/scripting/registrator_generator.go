package scripting

import (
	"cmp"
	"maps"
	"os"
	"slices"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func Write(api string, specs configuration.WebIdlConfigurations) error {
	statements := g.StatementList()
	registrator := g.Id("reg")
	s := slices.Collect(maps.Values(specs))
	slices.SortFunc(
		s,
		func(x, y *configuration.WebAPIConfig) int { return cmp.Compare(x.Name, y.Name) },
	)
	for _, spec := range s {
		data, err := idl.Load(spec.Name)
		if err != nil {
			return err
		}
		extra := make([]idl.Spec, len(spec.PartialSearchModules))
		for i, spec := range spec.PartialSearchModules {
			if extra[i], err = idl.Load(spec); err != nil {
				return err
			}
		}
		types := spec.GetTypesSorted()
		for _, specType := range types {
			typeInfo := createData(data, specType, extra)
			statements.Append(
				jsRegisterClass.Call(
					registrator,
					g.Lit(typeInfo.Name()),
					g.Lit(typeInfo.Extends()),
					g.Id(ConstructorNameForInterface(typeInfo.Name())),
				))
		}
	}

	bootstrap := g.Raw(jen.Func().Id("Bootstrap").Types(jen.Id("T").Any()).Params(
		jen.Add(registrator.Generate()).
			Add(jsClassBuilder.Generate()).Types(jen.Id("T")),
	).Block(statements.Generate()))

	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}

	return writeGenerator(writer, packagenames.ScriptPackageName(api), bootstrap)
}

func GenerateRegisterFunctions(spec string) error {
	specs := configuration.CreateV8SpecsForSpec(spec)
	return Write(spec, specs)
}
