package sobekhost

import (
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type iterator struct {
	vm   *sobek.Runtime
	next func() (js.Value[jsTypeParam], error, bool)
	stop func()
}
