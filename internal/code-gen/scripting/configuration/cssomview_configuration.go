package configuration

func ConfigureCssomView(domSpecs *WebAPIConfig) {
	element := domSpecs.Type("Element")
	element.Partial = true
	element.Method("getBoundingClientRect").SetCustomImplementation()
}
