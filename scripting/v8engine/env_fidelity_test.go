package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/scripting/v8engine"
	"github.com/stretchr/testify/assert"
)

// TestEnvFidelity exercises the browser-environment surface: navigator, screen,
// performance, crypto, the frame-less browsing context accessors, document
// state, legacy event creation and iframe contentWindow.
func TestEnvFidelity(t *testing.T) {
	win := browsertest.InitWindow(t, v8engine.DefaultEngine())
	eq := func(name, script string, want any) {
		t.Helper()
		assert.Equal(t, want, win.MustEval(script), name)
	}

	t.Run("navigator", func(t *testing.T) {
		eq("userAgent non-empty", `navigator.userAgent.length > 0`, true)
		eq("webdriver false", `navigator.webdriver`, false)
		eq("languages", `Array.isArray(navigator.languages) && navigator.languages[0] === "en-US"`, true)
		eq("hardwareConcurrency", `typeof navigator.hardwareConcurrency === "number" && navigator.hardwareConcurrency > 0`, true)
		eq("platform", `navigator.platform.length > 0`, true)
		eq("userAgentData brands", `Array.isArray(navigator.userAgentData.brands)`, true)
	})

	t.Run("screen", func(t *testing.T) {
		eq("width", `screen.width > 0`, true)
		eq("height", `screen.height > 0`, true)
		eq("availWidth <= width", `screen.availWidth <= screen.width`, true)
		eq("colorDepth", `screen.colorDepth`, int32(24))
	})

	t.Run("performance", func(t *testing.T) {
		eq("now is number", `typeof performance.now() === "number"`, true)
		eq("now monotonic", `(() => { const a = performance.now(); const b = performance.now(); return b >= a; })()`, true)
		eq("timeOrigin", `typeof performance.timeOrigin === "number" && performance.timeOrigin > 0`, true)
		eq("now native", `/\[native code\]/.test(performance.now.toString())`, true)
	})

	t.Run("crypto", func(t *testing.T) {
		eq("getRandomValues returns same array", `(() => { const a = new Uint8Array(16); return crypto.getRandomValues(a) === a; })()`, true)
		eq("getRandomValues fills", `(() => { const a = new Uint8Array(32); crypto.getRandomValues(a); return a.some(x => x !== 0); })()`, true)
		eq("randomUUID format", `/^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/.test(crypto.randomUUID())`, true)
		eq("getRandomValues native", `/\[native code\]/.test(crypto.getRandomValues.toString())`, true)
		eq("getRandomValues fills Uint32Array", `(() => { const a = new Uint32Array(8); crypto.getRandomValues(a); return a.some(x => x !== 0); })()`, true)
		eq("getRandomValues rejects Float32Array", `(() => { try { crypto.getRandomValues(new Float32Array(4)); return false; } catch (e) { return true; } })()`, true)
	})

	t.Run("frame-less browsing context", func(t *testing.T) {
		eq("top === self", `window.top === window.self`, true)
		eq("frames === self", `window.frames === window.self`, true)
		eq("length 0", `window.length`, int32(0))
		eq("devicePixelRatio", `typeof window.devicePixelRatio === "number"`, true)
	})

	t.Run("document state", func(t *testing.T) {
		eq("readyState string", `typeof document.readyState === "string"`, true)
		eq("compatMode", `document.compatMode`, "CSS1Compat")
		eq("characterSet", `document.characterSet`, "UTF-8")
		eq("cookie is string", `typeof document.cookie === "string"`, true)
		eq("hasFocus", `document.hasFocus()`, true)
		eq("visibilityState", `document.visibilityState`, "visible")
	})

	t.Run("legacy createEvent", func(t *testing.T) {
		eq("MouseEvent class", `document.createEvent("MouseEvents") instanceof MouseEvent`, true)
		eq("initEvent sets type", `(() => {
			const e = document.createEvent("Event");
			e.initEvent("custom-type", true, false);
			return e.type === "custom-type" && e.bubbles === true;
		})()`, true)
		eq("initMouseEvent sets type", `(() => {
			const e = document.createEvent("MouseEvents");
			e.initMouseEvent("click", true, true);
			return e.type === "click";
		})()`, true)
	})

	t.Run("iframe contentWindow", func(t *testing.T) {
		eq("contentWindow is object", `typeof document.createElement("iframe").contentWindow`, "object")
		eq("contentWindow.String is native", `(() => {
			const f = document.createElement("iframe");
			const cw = f.contentWindow;
			return cw != null && /\[native code\]/.test(cw.String.toString());
		})()`, true)
		eq("contentDocument is object", `typeof document.createElement("iframe").contentDocument`, "object")
	})
}
