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
	"github.com/gost-dom/code-gen/idlspec"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting/configuration"
	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

// classNameForMixin is necessary in order to install the operations and
// attributes specified on a mixin. E.g., module "fetch" defines operation
// "fetch" to be installed on "WindowOrWorkerGlobalScope". This means in a
// window scope, the operation belongs to the Window interface. In a Worker
// scope, the operation belongs on the Worker interface.
func classNameForMixin(r realm, data model.ESConstructorData) string {
	for intf := range idlspec.IdlInterfaces() {
		if r.exposes(intf) {
			for _, incl := range intf.Includes {
				if incl.Name == data.Name() {
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

func RegisterRealm(
	api string,
	realm realm,
	specs configuration.WebIdlConfigurations,
) (g.Generator, error) {
	idlSpec, err := idl.Load(api)
	if err != nil {
		return nil, err
	}
	statements := g.StatementList()
	engine := g.NewValue("e")
	var enriched []model.ESConstructorData
	for _, spec := range slices.Collect(maps.Values(specs)) {
		data, err := configuration.LoadSpecs(spec)
		if err != nil {
			return nil, err
		}
		types := spec.Types()
		for _, t := range types {
			typeInfo, err := createData(data, t)
			if err != nil {
				return nil, err
			}
			if IsGlobal(typeInfo.IdlInterface) {
				continue
			}
			enriched = append(enriched, typeInfo)
		}
	}
	slices.SortStableFunc(enriched, IntfComparer(idlSpec).compare)
	for _, typeInfo := range enriched {
		if typeInfo.InstallConstructor() && realm.exposes(typeInfo.IdlInterface) {
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
			name := typeInfo.Name()
			if typeInfo.IdlInterface.Mixin {
				name = classNameForMixin(realm, typeInfo)
			}
			instance := g.Id(internal.LowerCaseFirstLetter(name))
			ok := g.Id("ok")
			statements.Append(
				g.AssignMany(g.List(instance, ok), engine.Field("Class").Call(g.Lit(name))),
				g.IfStmt{Condition: gen.Not(ok), Block: gen.Panic(g.Lit(
					fmt.Sprintf("gost-dom/%s: %s: class not registered", api, name),
				))},
				Initializer(typeInfo).Call(instance),
			)
		}
	}
	return gen.NewFunction(
		gen.FunctionName(fmt.Sprintf("Configure%sRealm", realm.global.Name)),
		gen.FunctionTypeParam(gen.AnyConstraint(g.Id("T"))),
		gen.FunctionParam(engine, jsScriptEngine),
		gen.FunctionBody(statements),
	), nil
}

func GenerateRegisterFunctions(spec string, globals []string) error {
	gen := g.StatementList()
	for _, global := range globals {
		globalIntf, ok := idlspec.Interface(global)
		if !ok {
			return fmt.Errorf("Global interface not found: %s", global)
		}
		if len(globalIntf.Global) == 0 {
			return fmt.Errorf("Specified name has no exposed globals")
		}
		gen.Append(g.Line)
		specs := configuration.CreateV8SpecsForSpec(spec)
		res, err := RegisterRealm(spec, realm{globalIntf}, specs)
		if err != nil {
			return err
		}
		gen.Append(res)
	}
	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}

	return writeGenerator(writer, packagenames.ScriptPackageName(spec), gen)
}

func specNames() []customrules.Spec {
	names := customrules.Specs()
	slices.SortFunc(names, func(a, b customrules.Spec) int {
		if a.DependsOn(b) {
			return 1
		}
		if b.DependsOn(a) {
			return -1
		}
		return strings.Compare(string(a), string(b))
	})
	return names
}

func GenerateCombinedRegisterFunctions(globals []string) error {
	res := g.StatementList()
	for _, global := range globals {
		fnName := fmt.Sprintf("Configure%sRealm", global)
		e := g.Id("e")

		body := g.StatementList()
		for _, name := range specNames() {
			pkg := packagenames.ScriptPackageName(string(name))
			body.Append(g.NewValuePackage(fnName, pkg).Call(e))
		}
		res.Append(g.Line)
		res.Append(gen.NewFunction(
			gen.FunctionName(fnName),
			gen.FunctionTypeParam(gen.AnyConstraint(nil)),
			gen.FunctionParam(e, jsScriptEngine),
			gen.FunctionBody(body),
		))
	}
	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}
	return writeGenerator(writer, packagenames.ScriptingInt, res)
}
