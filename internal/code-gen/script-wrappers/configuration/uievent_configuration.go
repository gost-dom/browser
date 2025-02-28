package configuration

func ConfigureEventSpecs(specs *WebIdlConfigurations) {
	s := specs.Module("uievents")
	uiEvent := s.Type("UIEvent")
	uiEvent.Method("view").SetNotImplemented()
	uiEvent.Method("detail").SetNotImplemented()
}
