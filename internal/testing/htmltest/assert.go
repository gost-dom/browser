package htmltest

import (
	"fmt"
	"testing"
)

// Evaler is the interface for the single method Eval that executes JavaScript
// and returns the evaluated value.
type Evaler interface {
	Eval(string) (any, error)
}

// JsAssert helps write tests for the state of values in the JavaScript realm.
type JsAssert struct {
	t   testing.TB
	win Evaler
}

func (a JsAssert) mustEval(script string) any {
	a.t.Helper()
	res, err := a.win.Eval(script)
	if err != nil {
		a.t.Fatalf("Script error: %v. Running %s", err, script)
	}
	return res
}

// Verifies that the target object is an instance of the expected class. Both
// target and expected must be valid identifiers in global JavaScript scope.
func (a JsAssert) InstanceOf(target, expected string) {
	ok := a.mustEval(fmt.Sprintf("%s instanceof %s", target, expected)).(bool)
	if !ok {
		actual := a.mustEval(fmt.Sprintf("Object.getPrototypeOf(%s)?.constructor?.name", target))
		a.t.Errorf("Expected '%s' to be instance of %s\nActual: %s", target, expected, actual)
	}
}
