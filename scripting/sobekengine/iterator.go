package sobekengine

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type iterator struct {
	next func() (js.Value[jsTypeParam], error, bool)
	stop func()
}
