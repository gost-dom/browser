package scripting

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func inheritanceHierarchy(spec idl.Spec, intf idl.Interface) string {
	var res []string
	for intf.Inheritance != "" {
		res = append(res, intf.Name)
		if i, ok := spec.Interfaces[intf.Inheritance]; !ok {
			break
		} else {
			if slices.Contains(res, i.Name) {
				panic(fmt.Sprintf("Cyclic inheritance: %s (%v)", i.Name, res))
			}
			intf = i
		}
	}
	slices.Reverse(res)
	return strings.Join(res, ",")
}

type IntfComparer idl.Spec

func (c IntfComparer) compare(a, b model.ESConstructorData) int {
	inhA := inheritanceHierarchy(idl.Spec(c), a.IdlInterface)
	inhB := inheritanceHierarchy(idl.Spec(c), b.IdlInterface)
	return cmp.Compare(inhA, inhB)
}

func Write(api string, specs configuration.WebIdlConfigurations) error {
	idlSpec, err := idl.Load(api)
	if err != nil {
		return err
	}
	statements := g.StatementList()
	engine := g.Id("e")
	s := slices.Collect(maps.Values(specs))
	slices.SortFunc(
		s,
		func(x, y *configuration.WebAPIConfig) int { return cmp.Compare(x.Name, y.Name) },
	)
	var enriched = make([]model.ESConstructorData, 0)
	for _, spec := range s {
		data, extra, err := configuration.LoadSpecs(spec)
		if err != nil {
			return err
		}
		types := spec.GetTypesSorted()
		for _, t := range types {
			enriched = append(enriched, createData(data, t, extra))
		}
	}
	slices.SortStableFunc(enriched, IntfComparer(idlSpec).compare)
	for _, typeInfo := range enriched {
		statements.Append(
			jsRegisterClass.Call(
				engine,
				g.Lit(typeInfo.Name()),
				g.Lit(typeInfo.Extends()),
				g.Id(ConstructorNameForInterface(typeInfo.Name())),
			))
	}

	bootstrap := g.Raw(jen.Func().Id("Bootstrap").Types(jen.Id("T").Any()).Params(
		jen.Add(engine.Generate()).Add(jsScriptEngine.Generate()),
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
