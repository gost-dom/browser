package configuration

// WebIdlConfigurations is a list of specifications for generating ES wrapper
// code. Each key in the map correspond to a specific IDL file
type WebIdlConfigurations map[string](*WebAPIConfig)

func NewWrapperGeneratorsSpec() WebIdlConfigurations {
	return make(WebIdlConfigurations)
}

// Module returns the configuration for a specific spec. A new configuration is
// created if it doesn't exist.
func (c WebIdlConfigurations) Module(spec string) *WebAPIConfig {
	if mod, ok := c[spec]; ok {
		return mod
	}
	mod := NewWebAPIConfig(spec)
	c[spec] = mod
	return mod
}
