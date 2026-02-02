package htmltest

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

type ScriptContextHelper struct {
	html.ScriptContext
	t testing.TB
}

func NewScriptContextHelper(t testing.TB, ctx html.ScriptContext) ScriptContextHelper {
	return ScriptContextHelper{ctx, t}
}

func (h ScriptContextHelper) MustEval(script string) any {
	h.t.Helper()
	res, err := h.Eval(script)
	if err != nil {
		h.t.Fatalf("Script error: %v", err)
	}
	return res
}
