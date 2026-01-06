package configuration

// CodeGenStrategy is a "temporary" mechanism in a move towards generating code
// less coupled to the JavaScript engine, but targets an abstraction layer on
// top of the JS engine.
type CodeGenStrategy string

const (
	StrategyDefault     = ""            // Wrapper code targets JS engins
	StrategyAbstraction = "abstraction" // Wrapper code uses an abstraction layer on top of the JS engine
)

type GoType struct {
	Package string
	Name    string
	Pointer bool
}

// WebIDLConfig contains information about how to generate JavaScript mappings
// for a named item in the web IDL specification.
//
// All bindings will be generated using a set of defaults. Data in this
// structure will allow deviating from the defaults.
type WebIDLConfig struct {
	DomSpec             *WebAPIConfig
	TypeName            string
	RunCustomCode       bool
	SkipConstructor     bool
	SkipIterable        bool
	IncludeIncludes     bool
	Customization       map[string]*ESMethodWrapper
	OverrideWrappedType *GoType
	Partial             bool
}

func (c WebIDLConfig) SpecName() string { return c.DomSpec.Name }

func (w *WebIDLConfig) ensureMap() {
	if w.Customization == nil {
		w.Customization = make(map[string]*ESMethodWrapper)
	}
}

func (w *WebIDLConfig) MarkMembersAsNotImplemented(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{NotImplemented: true}
	}
}
func (w *WebIDLConfig) MarkMembersAsIgnored(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{Ignored: true}
	}
}

func (w *WebIDLConfig) GetMethodCustomization(name string) (result ESMethodWrapper) {
	if val, ok := w.Customization[name]; ok {
		result = *val
	}
	return
}

func (w *WebIDLConfig) Method(name string) (result *ESMethodWrapper) {
	w.ensureMap()
	var ok bool
	if result, ok = w.Customization[name]; !ok {
		result = new(ESMethodWrapper)
		w.Customization[name] = result
	}
	return result
}
