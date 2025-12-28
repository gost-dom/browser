package configuration

// type TypeCustomization []string

type ESMethodArgument struct {
	HasDefault   bool
	DefaultValue string
	Ignored      bool
	Decoder      string
}

func (a *ESMethodArgument) SetHasDefault() *ESMethodArgument { a.HasDefault = true; return a }

func (a *ESMethodArgument) HasDefaultValue(value string) {
	a.HasDefault = true
	a.DefaultValue = value
}

func (a *ESMethodArgument) Ignore() {
	a.Ignored = true
}

func (a *ESMethodArgument) SetDecoder(d string) *ESMethodArgument {
	a.Decoder = d
	return a
}
