package configuration

func ConfigureEventSpecs(s *WebAPIConfig) {
	uiEvent := s.Type("UIEvent")
	uiEvent.Method("view").SetNotImplemented()
	uiEvent.Method("detail").SetNotImplemented()
	mouseEvent := s.Type("MouseEvent")
	mouseEvent.SkipWrapper = true

	mouseEvent.Method("getModifierState").SetNotImplemented()
	mouseEvent.Method("screenX").SetNotImplemented()
	mouseEvent.Method("screenY").SetNotImplemented()
	mouseEvent.Method("clientX").SetNotImplemented()
	mouseEvent.Method("clientY").SetNotImplemented()
	mouseEvent.Method("layerX").SetNotImplemented()
	mouseEvent.Method("layerY").SetNotImplemented()
	mouseEvent.Method("ctrlKey").Ignore()
	mouseEvent.Method("shiftKey").Ignore()
	mouseEvent.Method("altKey").Ignore()
	mouseEvent.Method("metaKey").Ignore()
	mouseEvent.Method("button").Ignore()
	mouseEvent.Method("buttons").Ignore()
	mouseEvent.Method("relatedTarget").SetNotImplemented()

	keyboardEvent := s.Type("KeyboardEvent")
	keyboardEvent.SkipWrapper = true

	keyboardEvent.Method("getModifierState").SetNotImplemented()
}

func ConfigurePointerEventSpecs(s *WebAPIConfig) {
	pointerEvent := s.Type("PointerEvent")
	pointerEvent.SkipWrapper = true
	pointerEvent.Method("getCoalescedEvents").Ignore()
	pointerEvent.Method("getPredictedEvents").Ignore()
	pointerEvent.Method("pointerId").Ignore()
	pointerEvent.Method("width").SetNotImplemented()
	pointerEvent.Method("height").SetNotImplemented()
	pointerEvent.Method("pressure").SetNotImplemented()
	pointerEvent.Method("tangentialPressure").SetNotImplemented()
	pointerEvent.Method("tiltX").Ignore()
	pointerEvent.Method("tiltY").Ignore()
	pointerEvent.Method("altitudeAngle").Ignore()
	pointerEvent.Method("azimuthAngle").Ignore()
	pointerEvent.Method("pointerType").Ignore()
	pointerEvent.Method("isPrimary").Ignore()
	pointerEvent.Method("persistentDeviceId").Ignore()
	pointerEvent.Method("twist").Ignore()
}
