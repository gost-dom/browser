package html

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w HTMLOrSVGElement[T]) InstallPrototype(jsClass js.Class[T]) {
	w.installPrototype(jsClass)
}
