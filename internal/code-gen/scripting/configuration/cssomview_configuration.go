package configuration

func ConfigureCssomView(domSpecs *WebAPIConfig) {
	element := domSpecs.Type("Element")
	element.Partial = true
	element.Method("getBoundingClientRect").SetCustomImplementation()
	element.Method("getClientRects").SetCustomImplementation()
	rng := domSpecs.Type("Range")
	rng.Partial = true
	rng.Method("getBoundingClientRect").SetCustomImplementation()
	rng.Method("getClientRects").SetCustomImplementation()
}
