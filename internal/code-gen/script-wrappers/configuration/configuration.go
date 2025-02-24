// Package configuration is part of an internal code generation tool for
// Gost-DOM. It is not to be used in other context, and is not used at runtime.
package configuration

func CreateSpecs() WebIdlConfigurations {
	specs := NewWrapperGeneratorsSpec()

	ConfigureDOMSpecs(&specs)
	ConfigureHTMLSpecs(&specs)

	return specs
}
