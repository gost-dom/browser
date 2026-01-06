package htmltest

import (
	"fmt"
	"testing"

	"github.com/gost-dom/browser/html"
)

type JsAssert struct {
	t   testing.TB
	win html.Window
}

func (a JsAssert) mustEval(script string) any {
	a.t.Helper()
	res, err := a.win.Eval(script)
	if err != nil {
		a.t.Errorf("Script error: %s", script)
		a.t.Fatal(err)
	}
	return res
}

func (a JsAssert) InstanceOf(target, expected string) {
	ok := a.mustEval(fmt.Sprintf("%s instanceof %s", target, expected)).(bool)
	if !ok {
		actual := a.mustEval(fmt.Sprintf("Object.getPrototypeOf(%s)?.constructor?.name", target))
		a.t.Errorf("Expected '%s' to be instance of %s\nActual: %s", target, expected, actual)
	}
}
