package v8gen

import (
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/scripting"
	"github.com/gost-dom/code-gen/scripting/configuration"
)

func NewScriptWrapperModulesGeneratorForSpec(spec string) scripting.ScriptWrapperModulesGenerator {
	specs := configuration.CreateV8SpecsForSpec(spec)

	return scripting.ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      packagenames.ScriptPackageName(spec),
		TargetGenerators: V8TargetGenerators{},
	}
}
