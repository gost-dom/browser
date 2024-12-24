package main

type TypeCustomization []string

// ESMethodWrapper contains information about how to generate ES wrapper code
// around a single class method.
type ESMethodWrapper struct {
	// When set, the ES wrapper will generate an error with the message, "Not implemented"
	NotImplemented bool
}

// ESClassWrapper contains information about how to generate ES wrapper code
// around a class in the web specification.
//
// All classes will be generated using a set of defaults. Data in this structure
// will allow deviating from the defaults.
type ESClassWrapper struct {
	TypeName        string
	InnerTypeName   string
	WrapperTypeName string
	Receiver        string
	Customization   map[string]*ESMethodWrapper
}

func (w *ESClassWrapper) ensureMap() {
	if w.Customization == nil {
		w.Customization = make(map[string]*ESMethodWrapper)
	}
}

func (w *ESClassWrapper) MarkMembersAsNotImplemented(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{NotImplemented: true}
	}
}

func (w *ESClassWrapper) GetMethodCustomization(name string) (result ESMethodWrapper) {
	if val, ok := w.Customization[name]; ok {
		result = *val
	}
	return
}

func (w *ESClassWrapper) Method(name string) (result *ESMethodWrapper) {
	w.ensureMap()
	var ok bool
	if result, ok = w.Customization[name]; !ok {
		result = new(ESMethodWrapper)
		w.Customization[name] = result
	}
	return result
}

// SetNotImplemented is a simple wrapper around [NotImplemented] to support a
// chaning DSL syntax.
func (w *ESMethodWrapper) SetNotImplemented() *ESMethodWrapper {
	w.NotImplemented = true
	return w
}
