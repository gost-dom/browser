package v8gen

import (
	"github.com/gost-dom/code-gen/packagenames"
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
)

func NewScriptWrapperModulesGeneratorForSpec(spec string) wrappers.ScriptWrapperModulesGenerator {
	specs := configuration.CreateV8SpecsForSpec(spec)

	return wrappers.ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      packagenames.ScriptPackageName(spec),
		TargetGenerators: V8TargetGenerators{},
	}
}
