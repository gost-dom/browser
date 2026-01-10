package scripting

import (
	"slices"

	"github.com/gost-dom/webref/idl"
)

// realm represents an isolated JavaScript environment with it's own set of
// values available in global scope. E.g., the initial realm, where scripts on
// the page are executed, have the Window object as global object. Workers run
// in separate realms and have different globals, e.g.
// DedicatedWorkerGlobalScope.
type realm struct {
	global idl.Interface
}

// exposes returns whether interface i is exposed in the realm.
func (r realm) exposes(i idl.Interface) bool {
	if len(i.Exposed) == 0 {
		// Workaround for a webref bug where Exposed=* isn't included in the
		// exposed list
		return true
	}
	if slices.Contains(i.Exposed, "*") {
		return true
	}
	for _, global := range r.global.Global {
		if slices.Contains(i.Exposed, global) {
			return true
		}
	}
	return false
}
