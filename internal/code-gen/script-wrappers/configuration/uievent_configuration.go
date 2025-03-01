package configuration

func ConfigureEventSpecs(specs *WebIdlConfigurations) {
	s := specs.Module("uievents")
	uiEvent := s.Type("UIEvent")
	uiEvent.Method("view").SetNotImplemented()
	uiEvent.Method("detail").SetNotImplemented()
	mouseEvent := s.Type("MouseEvent")
	mouseEvent.SkipWrapper = true
	mouseEvent.Method("getModifierState").Ignore()
	mouseEvent.Method("screenX").Ignore()
	mouseEvent.Method("screenY").Ignore()
	mouseEvent.Method("clientX").Ignore()
	mouseEvent.Method("clientY").Ignore()
	mouseEvent.Method("layerX").Ignore()
	mouseEvent.Method("layerY").Ignore()
	mouseEvent.Method("ctrlKey").Ignore()
	mouseEvent.Method("shiftKey").Ignore()
	mouseEvent.Method("altKey").Ignore()
	mouseEvent.Method("metaKey").Ignore()
	mouseEvent.Method("button").Ignore()
	mouseEvent.Method("buttons").Ignore()
	mouseEvent.Method("relatedTarget").Ignore()

	ConfigurePointerEventSpecs(specs)
}

func ConfigurePointerEventSpecs(specs *WebIdlConfigurations) {
	s := specs.Module("pointerevents4")
	pointerEvent := s.Type("PointerEvent")
	pointerEvent.SkipWrapper = true
}
