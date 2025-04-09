package v8gen

import (
	"github.com/gost-dom/code-gen/packagenames"
	wrappers "github.com/gost-dom/code-gen/script-wrappers"
	"github.com/gost-dom/code-gen/script-wrappers/configuration"
)

func NewScriptWrapperModulesGenerator() wrappers.ScriptWrapperModulesGenerator {
	specs := configuration.CreateV8Specs()

	return wrappers.ScriptWrapperModulesGenerator{
		Specs:            specs,
		PackagePath:      packagenames.V8host,
		TargetGenerators: V8TargetGenerators{},
	}
}
