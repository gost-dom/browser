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
	var first = true
	for {
		if first {
			res = append(res, "."+intf.Name)
			first = false
		} else {
			res = append(res, intf.Name)
		}
		if i, ok := spec.Interfaces[intf.Inheritance]; !ok {
			res = append(res, intf.Inheritance)
			break
		} else {
			if slices.Contains(res, i.Name) {
				panic(fmt.Sprintf("Cyclic inheritance: %s (%v)", i.Name, res))
			}
			intf = i
		}
	}
	res = append(res, ".Object")
	slices.Reverse(res)
	return strings.Join(res, ",")
}

type IntfComparer idl.Spec

func (c IntfComparer) compare(a, b model.ESConstructorData) int {
	inhA := inheritanceHierarchy(idl.Spec(c), a.IdlInterface)
	inhB := inheritanceHierarchy(idl.Spec(c), b.IdlInterface)
	return cmp.Compare(inhA, inhB)
}

func IsGlobal(intf idl.Interface) bool { return len(intf.Global) > 0 }

func Write(api string, specs configuration.WebIdlConfigurations) error {
	idlSpec, err := idl.Load(api)
	if err != nil {
		return err
	}
	statements := g.StatementList()
	engine := g.Id("e")
	var enriched []model.ESConstructorData
	for _, spec := range slices.Collect(maps.Values(specs)) {
		data, extra, err := configuration.LoadSpecs(spec)
		if err != nil {
			return err
		}
		types := spec.Types()
		for _, t := range types {
			typeInfo := createData(data, t, extra)
			if IsGlobal(typeInfo.IdlInterface) {
				continue
			}
			enriched = append(enriched, typeInfo)
		}
	}
	slices.SortStableFunc(enriched, IntfComparer(idlSpec).compare)
	for _, typeInfo := range enriched {
		if typeInfo.InstallConstructor() {
			statements.Append(
				jsRegisterClass.Call(
					engine,
					g.Lit(typeInfo.Name()),
					g.Lit(typeInfo.Extends()),
					Initializer(typeInfo),
					g.Id(JsConstructorForInterface(typeInfo.Name())),
				))
		}
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
