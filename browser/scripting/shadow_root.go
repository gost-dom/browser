package scripting

import (
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

// createShadowRootPrototype currently only exists to allow code to check
// for inheritence, i.e., `node instanceof DocumentFragment`.
// This is performed by HTMX; but it doesn't itself _create_ a shadow root.
func createShadowRootPrototype(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[ShadowRoot](host)
	return builder.constructor
}
