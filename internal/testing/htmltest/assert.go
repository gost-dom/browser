package htmltest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func (a JsAssert) True(expr string) {
	a.t.Helper()
	if val, ok := evalAndAssertType[bool](a, expr); ok {
		assert.True(a.t, val)
	}
}

func (a JsAssert) False(expr string) {
	a.t.Helper()
	if val, ok := evalAndAssertType[bool](a, expr); ok {
		assert.False(a.t, val)
	}
}

func evalAndAssertType[T any](a JsAssert, expr string) (T, bool) {
	val := a.mustEval(expr)
	res, ok := val.(T)
	if !ok {
		var dummy T
		a.t.Errorf("Expected type %T evaluating: %s\nGot: %T", dummy, expr, val)
	}
	return res, ok
}
