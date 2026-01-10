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

// installers has the methods to create the configuration of a specific web API
// in a specific realm. E.g., install the dom API in the Window realm.
type installers struct {
	api   string
	realm realm
	specs configuration.WebIdlConfigurations
}

func newInstallers(api, global string) (installers, error) {
	globalIntf, ok := idlspec.Interface(global)
	if !ok {
		return installers{}, fmt.Errorf(
			"GenerateRegisterFunctions: IDL interface not found: %s",
			global,
		)
	}
	if len(globalIntf.Global) == 0 {
		return installers{}, fmt.Errorf(
			"GenerateRegisterFunctions: IDL interface has no globals: %s",
			global,
		)
	}
	specs := configuration.CreateV8SpecsForSpec(string(api))
	return installers{api, realm{globalIntf}, specs}, nil
}

func (i installers) empty() bool {
	fns, _ := i.InstallersForRealm()
	return len(fns) == 0
}

func (i installers) InstallersForRealm() ([]g.Generator, error) {
	res := make([]g.Generator, 0, 20)
	idlSpec, err := idl.Load(i.api)
	if err != nil {
		return nil, err
	}
	engine := scriptEngine{g.NewValue("e")}
	var enriched []model.ESConstructorData
	for _, spec := range slices.Collect(maps.Values(i.specs)) {
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
			enriched = append(enriched, typeInfo)
		}
	}
	slices.SortStableFunc(enriched, IntfComparer(idlSpec).compare)
	for _, typeInfo := range enriched {
		if typeInfo.InstallConstructor() && i.realm.exposes(typeInfo.IdlInterface) {
			var constructor g.Generator
			if typeInfo.AllowConstructor() {
				constructor = g.Id(JsConstructorForInterface(typeInfo.Name()))
			} else {
				constructor = g.Nil
			}
			baseClass := g.Nil
			if IsGlobal(typeInfo.IdlInterface) {
				if inherits := i.realm.global.Inheritance; inherits != "" {
					baseClass = MustGetClass(engine, inherits)
				}
			}

			res = append(res,
				Initializer(typeInfo).Call(
					renderIfElse(IsGlobal(typeInfo.IdlInterface),
						engine.ConfigureGlobalScope(i.realm.global.Name, baseClass),
						jsCreateClass.Call(
							engine,
							g.Lit(typeInfo.Name()),
							g.Lit(typeInfo.Extends()),
							constructor,
						))),
			)
		}
		if typeInfo.InstallPartial() {
			name := typeInfo.Name()
			if typeInfo.IdlInterface.Mixin {
				name = classNameForMixin(i.realm, typeInfo)
			}
			original, _ := idlspec.Interface(name)
			if i.realm.exposes(original) {
				res = append(res,
					Initializer(typeInfo).Call(MustGetClass(engine, name)),
				)
			}
		}
	}
	return res, nil
}

func (i installers) RegisterRealm() (g.Generator, error) {
	engine := scriptEngine{g.NewValue("e")}
	res, err := i.InstallersForRealm()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return g.Noop, nil
	}
	statements := g.StatementList(res...)
	return gen.NewFunction(
		gen.FunctionName(fmt.Sprintf("Configure%sRealm", i.realm.global.Name)),
		gen.FunctionTypeParam(gen.AnyConstraint(g.Id("T"))),
		gen.FunctionParam(engine, jsScriptEngine),
		gen.FunctionBody(statements),
	), nil
}

func GenerateRegisterFunctions(spec string, globals []string) error {
	gen := g.StatementList()
	for _, global := range globals {
		i, err := newInstallers(spec, global)
		if err != nil {
			return err
		}
		res, err := i.RegisterRealm()
		if err != nil {
			return err
		}
		gen.Append(g.Line)
		gen.Append(res)
	}
	writer, err := os.Create("register_generated.go")
	if err != nil {
		return err
	}
	defer writer.Close()

	return writeGenerator(writer, packagenames.ScriptPackageName(spec), gen)
}

func GenerateCombinedRegisterFunctions(globals []string) error {
	res := g.StatementList()
	for _, global := range globals {
		fnName := fmt.Sprintf("Configure%sRealm", global)
		e := g.Id("e")

		body := g.StatementList()
		for _, name := range customrules.Specs() {
			i, err := newInstallers(string(name), global)
			if err != nil {
				return err
			}
			if !i.empty() {
				pkg := packagenames.ScriptPackageName(string(name))
				body.Append(g.NewValuePackage(fnName, pkg).Call(e))
			}
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
	defer writer.Close()
	return writeGenerator(writer, packagenames.ScriptingInt, res)
}
