package scripting

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/code-gen/gen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func exposedTo(intf idl.Interface, globals []string) bool {
	for _, global := range globals {
		if slices.Contains(intf.Exposed, global) {
			return true
		}
	}
	return false
}

// classNameForMixin is necessary in order to install the operations and
// attributes specified on a mixin. E.g., module "fetch" defines operation
// "fetch" to be installed on "WindowOrWorkerGlobalScope". This means in a
// window scope, the operation belongs to the Window interface. In a Worker
// scope, the operation belongs on the Worker interface.
func classNameForMixin(globals []string, data model.ESConstructorData) string {
	for _, name := range customrules.SpecNames() {
		spec, err := idl.Load(name)
		if err != nil {
			panic(fmt.Sprintf("Unknown spec in custom_rules: %v", err))
		}
		for _, intf := range spec.Interfaces {
			for _, incl := range intf.Includes {
				if exposedTo(intf, globals) && incl.Name == data.Name() {
					return intf.Name
				}
			}
		}
	}
	panic(fmt.Sprintf("Class name for mixin not found: %s", data.Name()))
}

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

func Write(api string, globals []string, specs configuration.WebIdlConfigurations) error {
	idlSpec, err := idl.Load(api)
	if err != nil {
		return err
	}
	statements := g.StatementList()
	engine := g.NewValue("e")
	var enriched []model.ESConstructorData
	for _, spec := range slices.Collect(maps.Values(specs)) {
		data, extra, err := configuration.LoadSpecs(spec)
		if err != nil {
			return err
		}
		types := spec.Types()
		for _, t := range types {
			typeInfo, err := createData(data, t, extra)
			if err != nil {
				return err
			}
			if IsGlobal(typeInfo.IdlInterface) {
				continue
			}
			enriched = append(enriched, typeInfo)
		}
	}
	slices.SortStableFunc(enriched, IntfComparer(idlSpec).compare)
	for _, typeInfo := range enriched {
		if typeInfo.InstallConstructor() {
			var constructor g.Generator
			if typeInfo.AllowConstructor() {
				constructor = g.Id(JsConstructorForInterface(typeInfo.Name()))
			} else {
				constructor = g.Nil
			}
			statements.Append(
				Initializer(typeInfo).Call(
					jsCreateClass.Call(
						engine,
						g.Lit(typeInfo.Name()),
						g.Lit(typeInfo.Extends()),
						constructor,
					)),
			)
		}
		if typeInfo.InstallPartial() {
			name := classNameForMixin(globals, typeInfo)
			instance := g.Id(internal.LowerCaseFirstLetter(name))
			ok := g.Id("ok")
			statements.Append(
				g.AssignMany(g.List(instance, ok), engine.Field("Class").Call(g.Lit(name))),
				g.IfStmt{Condition: gen.Not(ok), Block: gen.Panic(g.Lit(""))},
				Initializer(typeInfo).Call(instance),
			)
		}
	}
	bootstrap := gen.NewFunction(
		gen.FunctionName("Bootstrap"),
		gen.FunctionTypeParam(gen.AnyConstraint(g.Id("T"))),
		gen.FunctionParam(engine, jsScriptEngine),
		gen.FunctionBody(statements),
	)

	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}

	return writeGenerator(writer, packagenames.ScriptPackageName(api), bootstrap)
}

func GenerateRegisterFunctions(spec string, globals []string) error {
	specs := configuration.CreateV8SpecsForSpec(spec)
	return Write(spec, globals, specs)
}
