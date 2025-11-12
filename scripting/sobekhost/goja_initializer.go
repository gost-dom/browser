package sobekhost

import (
	"github.com/gost-dom/browser/scripting/internal"
)

var initializer *internal.ScriptEngineConfigurer[jsTypeParam]

func init() {
	initializer = internal.DefaultInitializer[jsTypeParam]()
}
