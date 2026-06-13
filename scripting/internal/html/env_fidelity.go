package html

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// This file fills in commonly-used parts of the browser environment that the
// Web-IDL code generator does not (yet) produce: a populated navigator, and the
// screen, performance and crypto interfaces. Scripts frequently read these, and
// feature-detect or branch on their presence.
//
// Everything here is implemented with native (FunctionTemplate-backed)
// accessors and operations so that, like a real browser, Function.prototype.
// toString reports "[native code]" rather than leaking a JavaScript polyfill's
// source.
//
// The values describe a single, internally consistent profile: Chrome 131 on
// 64-bit Windows 10. navigator.userAgent, userAgentData, platform and the
// screen metrics are all chosen to agree with one another.

const (
	profileUserAgent   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	profileAppVersion  = "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	profilePlatform    = "Win32"
	profileVendor      = "Google Inc."
	profileProduct     = "Gecko"
	profileProductSub  = "20030107"
	profileAppName     = "Netscape"
	profileAppCodeName = "Mozilla"
	profileLanguage    = "en-US"
	profileChromeMajor = "131"

	profileScreenWidth  = 1920
	profileScreenHeight = 1080
	profileAvailWidth   = 1920
	profileAvailHeight  = 1040 // minus a 40px taskbar
	profileColorDepth   = 24

	profileInnerWidth  = 1920
	profileInnerHeight = 945
	profileOuterWidth  = 1920
	profileOuterHeight = 1040
)

// configureEnvFidelity installs the navigator/screen/performance/crypto surface
// onto the already-configured Window realm.
func configureEnvFidelity[T any](e js.ScriptEngine[T]) {
	if nav, ok := e.Class("Navigator"); ok {
		configureNavigator(nav)
	}

	screen := js.CreateClass(e, "Screen", "", js.IllegalConstructor)
	configureScreen(screen)
	performance := js.CreateClass(e, "Performance", "", js.IllegalConstructor)
	configurePerformance(performance)
	crypto := js.CreateClass(e, "Crypto", "", js.IllegalConstructor)
	configureCrypto(crypto)

	if win, ok := e.Class("Window"); ok {
		configureWindowEnv(win)
	}
	if doc, ok := e.Class("Document"); ok {
		configureDocumentEnv(doc)
	}
	if ev, ok := e.Class("Event"); ok {
		configureEventEnv(ev)
	}

	// HTMLIFrameElement.contentWindow / contentDocument expose the nested
	// browsing context (a separate JavaScript realm). The placeholder class is
	// created by installHtmlElementTypes, which runs before this.
	if iframe, ok := e.Class("HTMLIFrameElement"); ok {
		iframe.CreateAttribute("contentWindow", iframeContentWindow, nil)
		iframe.CreateAttribute("contentDocument", iframeContentDocument, nil)
	}
}

// iframeContentWindow returns the iframe's nested-browsing-context global
// object. Because the child Window's JS global is cached when its script
// context is created, EncodeEntity hands back that very object — a value living
// in the child realm, so reads like contentWindow.String resolve to the child
// realm's genuine native constructors.
func iframeContentWindow[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	el, err := js.As[html.HTMLIFrameElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.Null(), nil
	}
	w := el.ContentWindow()
	if w == nil {
		return cbCtx.Null(), nil
	}
	return codec.EncodeEntity(cbCtx, w)
}

// iframeContentDocument returns the document of the iframe's nested browsing
// context (resolved within the same realm as contentWindow).
func iframeContentDocument[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	el, err := js.As[html.HTMLIFrameElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.Null(), nil
	}
	w := el.ContentWindow()
	if w == nil {
		return cbCtx.Null(), nil
	}
	// Read `document` off the child realm's global so the returned value also
	// belongs to the child realm (rather than a parent-realm wrapper).
	cw, err := codec.EncodeEntity(cbCtx, w)
	if err != nil {
		return nil, err
	}
	obj, ok := cw.AsObject()
	if !ok {
		return cbCtx.Null(), nil
	}
	return obj.Get("document")
}

/* -------------------------------------------------------------------------- */
/* navigator                                                                   */
/* -------------------------------------------------------------------------- */

// configureNavigator installs the navigator interface's data attributes.
func configureNavigator[T any](nav js.Class[T]) {
	str := func(v string) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewString(v), nil }
	}
	num := func(v int32) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewInt32(v), nil }
	}
	boolean := func(v bool) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewBoolean(v), nil }
	}

	nav.CreateAttribute("userAgent", str(profileUserAgent), nil)
	nav.CreateAttribute("appVersion", str(profileAppVersion), nil)
	nav.CreateAttribute("appName", str(profileAppName), nil)
	nav.CreateAttribute("appCodeName", str(profileAppCodeName), nil)
	nav.CreateAttribute("platform", str(profilePlatform), nil)
	nav.CreateAttribute("vendor", str(profileVendor), nil)
	nav.CreateAttribute("vendorSub", str(""), nil)
	nav.CreateAttribute("product", str(profileProduct), nil)
	nav.CreateAttribute("productSub", str(profileProductSub), nil)
	nav.CreateAttribute("language", str(profileLanguage), nil)
	nav.CreateAttribute("languages", navigatorLanguages, nil)
	nav.CreateAttribute("hardwareConcurrency", num(8), nil)
	nav.CreateAttribute("maxTouchPoints", num(0), nil)
	nav.CreateAttribute("deviceMemory", num(8), nil)
	nav.CreateAttribute("webdriver", boolean(false), nil)
	nav.CreateAttribute("onLine", boolean(true), nil)
	nav.CreateAttribute("cookieEnabled", boolean(true), nil)
	nav.CreateAttribute("pdfViewerEnabled", boolean(true), nil)
	nav.CreateAttribute("userAgentData", navigatorUserAgentData, nil)
}

// navigatorLanguages implements the navigator.languages getter.
func navigatorLanguages[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.NewArray(cbCtx.NewString("en-US"), cbCtx.NewString("en")), nil
}

// navigatorUserAgentData returns a NavigatorUAData-shaped object whose brands
// and platform agree with the userAgent string above.
func navigatorUserAgentData[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	brand := func(name, version string) js.Value[T] {
		o := cbCtx.NewObject()
		_ = o.Set("brand", cbCtx.NewString(name))
		_ = o.Set("version", cbCtx.NewString(version))
		return o
	}
	data := cbCtx.NewObject()
	if err := data.Set("brands", cbCtx.NewArray(
		brand("Google Chrome", profileChromeMajor),
		brand("Chromium", profileChromeMajor),
		brand("Not_A Brand", "24"),
	)); err != nil {
		return nil, err
	}
	_ = data.Set("mobile", cbCtx.NewBoolean(false))
	_ = data.Set("platform", cbCtx.NewString("Windows"))
	return data, nil
}

/* -------------------------------------------------------------------------- */
/* screen                                                                      */
/* -------------------------------------------------------------------------- */

type screenObj struct{ entity.Entity }

// configureScreen installs the Screen interface's metric attributes.
func configureScreen[T any](screen js.Class[T]) {
	num := func(v int32) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewInt32(v), nil }
	}
	screen.CreateAttribute("width", num(profileScreenWidth), nil)
	screen.CreateAttribute("height", num(profileScreenHeight), nil)
	screen.CreateAttribute("availWidth", num(profileAvailWidth), nil)
	screen.CreateAttribute("availHeight", num(profileAvailHeight), nil)
	screen.CreateAttribute("availLeft", num(0), nil)
	screen.CreateAttribute("availTop", num(0), nil)
	screen.CreateAttribute("colorDepth", num(profileColorDepth), nil)
	screen.CreateAttribute("pixelDepth", num(profileColorDepth), nil)
}

/* -------------------------------------------------------------------------- */
/* performance                                                                 */
/* -------------------------------------------------------------------------- */

type performanceObj struct {
	entity.Entity
	origin time.Time
}

// configurePerformance installs the Performance interface's members.
func configurePerformance[T any](performance js.Class[T]) {
	performance.CreateOperation("now", performanceNow)
	performance.CreateAttribute("timeOrigin", performanceTimeOrigin, nil)
}

// performanceNow implements performance.now(), returning milliseconds elapsed
// since the Performance time origin.
func performanceNow[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	inst, err := js.As[*performanceObj](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	ms := float64(time.Since(inst.origin).Nanoseconds()) / 1e6
	return cbCtx.NewNumber(ms), nil
}

// performanceTimeOrigin implements the performance.timeOrigin getter.
func performanceTimeOrigin[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	inst, err := js.As[*performanceObj](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.NewNumber(float64(inst.origin.UnixNano()) / 1e6), nil
}

/* -------------------------------------------------------------------------- */
/* crypto                                                                      */
/* -------------------------------------------------------------------------- */

type cryptoObj struct{ entity.Entity }

// configureCrypto installs the Crypto interface's operations.
func configureCrypto[T any](crypto js.Class[T]) {
	crypto.CreateOperation("getRandomValues", cryptoGetRandomValues)
	crypto.CreateOperation("randomUUID", cryptoRandomUUID)
}

// cryptoGetRandomValues fills the passed integer TypedArray with
// cryptographically-strong random values and returns the same array, matching
// Crypto.getRandomValues.
func cryptoGetRandomValues[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	args := cbCtx.Args()
	if len(args) == 0 {
		return nil, cbCtx.NewTypeError("Crypto.getRandomValues: 1 argument required")
	}
	arg := args[0]
	obj, ok := arg.AsObject()
	if !ok {
		return nil, cbCtx.NewTypeError("Crypto.getRandomValues: argument is not an object")
	}
	bpeVal, err := obj.Get("BYTES_PER_ELEMENT")
	if err != nil {
		return nil, err
	}
	bpe := int(bpeVal.Int32())
	if bpe == 0 || bpe > 4 {
		// bpe==0 is a non-TypedArray; >4-byte elements are BigInt64/BigUint64
		// arrays, which need BigInt values this implementation does not produce.
		return nil, cbCtx.NewTypeError(
			"Crypto.getRandomValues: argument is not an integer TypedArray of <= 32-bit elements")
	}
	// Float typed arrays (e.g. Float32Array, which also has a 4-byte element)
	// must be rejected with a TypeMismatchError per the spec, not filled.
	if name := typedArrayName(obj); strings.HasPrefix(name, "Float") {
		return nil, cbCtx.NewTypeError(
			"Crypto.getRandomValues: " + name + " is not an integer TypedArray")
	}
	lengthVal, err := obj.Get("length")
	if err != nil {
		return nil, err
	}
	n := int(lengthVal.Int32())
	if n*bpe > 65536 {
		return nil, cbCtx.NewError(fmt.Errorf("Crypto.getRandomValues: array exceeds 65536 byte quota"))
	}
	raw := make([]byte, n*4)
	if _, err := rand.Read(raw); err != nil {
		return nil, cbCtx.NewError(err)
	}
	for i := 0; i < n; i++ {
		// V8 truncates the assigned Number to the element width, so a random
		// uint32 yields uniform random bits for 8-, 16- and 32-bit elements.
		v := binary.LittleEndian.Uint32(raw[i*4:])
		if err := obj.Set(strconv.Itoa(i), cbCtx.NewUint32(v)); err != nil {
			return nil, err
		}
	}
	return arg, nil
}

// typedArrayName returns the constructor name of a TypedArray (e.g.
// "Uint8Array"), or "" if it cannot be determined.
func typedArrayName[T any](obj js.Object[T]) string {
	ctor, err := obj.Get("constructor")
	if err != nil {
		return ""
	}
	ctorObj, ok := ctor.AsObject()
	if !ok {
		return ""
	}
	name, err := ctorObj.Get("name")
	if err != nil {
		return ""
	}
	return name.String()
}

// cryptoRandomUUID implements crypto.randomUUID, returning a random RFC 4122
// version 4 UUID string.
func cryptoRandomUUID[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, cbCtx.NewError(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
	return cbCtx.NewString(uuid), nil
}

/* -------------------------------------------------------------------------- */
/* window                                                                      */
/* -------------------------------------------------------------------------- */

// configureWindowEnv installs window-level environment attributes (screen,
// performance, crypto, viewport metrics and the frame-less browsing context).
func configureWindowEnv[T any](win js.Class[T]) {
	num := func(v int32) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewInt32(v), nil }
	}
	win.CreateAttribute("screen", windowScreen, nil)
	win.CreateAttribute("performance", windowPerformance, nil)
	win.CreateAttribute("crypto", windowCrypto, nil)
	win.CreateAttribute("devicePixelRatio", func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
		return cbCtx.NewNumber(1), nil
	}, nil)
	win.CreateAttribute("innerWidth", num(profileInnerWidth), nil)
	win.CreateAttribute("innerHeight", num(profileInnerHeight), nil)
	win.CreateAttribute("outerWidth", num(profileOuterWidth), nil)
	win.CreateAttribute("outerHeight", num(profileOuterHeight), nil)
	win.CreateAttribute("screenX", num(0), nil)
	win.CreateAttribute("screenY", num(0), nil)
	win.CreateAttribute("screenLeft", num(0), nil)
	win.CreateAttribute("screenTop", num(0), nil)

	// top/frames/length describe the (top-level, frame-less) browsing context.
	// These were previously unimplemented and *threw*; window.top in particular
	// is read very commonly (e.g. the frame-busting idiom window.top === window),
	// often without a try/catch, so a throw breaks otherwise-correct scripts. A
	// top-level window has top === self, frames === self and length === 0.
	self := func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.This(), nil }
	win.CreateAttribute("top", self, nil)
	win.CreateAttribute("frames", self, nil)
	win.CreateAttribute("length", num(0), nil)
}

// configureDocumentEnv installs commonly-read Document attributes and the
// legacy hasFocus operation.
func configureDocumentEnv[T any](doc js.Class[T]) {
	str := func(v string) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewString(v), nil }
	}
	boolean := func(v bool) js.CallbackFunc[T] {
		return func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return cbCtx.NewBoolean(v), nil }
	}
	// These were either unimplemented (and threw) or ignored (undefined), but a
	// real browser always exposes them as plain string/boolean values.
	doc.CreateAttribute("compatMode", str("CSS1Compat"), nil)
	doc.CreateAttribute("characterSet", str("UTF-8"), nil)
	doc.CreateAttribute("charset", str("UTF-8"), nil)
	doc.CreateAttribute("inputEncoding", str("UTF-8"), nil)
	doc.CreateAttribute("contentType", str("text/html"), nil)
	// readyState reflects the real document load state ("loading" while parsing,
	// "complete" afterwards) rather than a constant, so scripts that defer work
	// until the DOM is ready (e.g. htmx) still behave correctly.
	doc.CreateAttribute("readyState", documentReadyState, nil)
	doc.CreateAttribute("visibilityState", str("visible"), nil)
	doc.CreateAttribute("hidden", boolean(false), nil)
	doc.CreateOperation("hasFocus", func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
		return cbCtx.NewBoolean(true), nil
	})
	// cookie and title are always strings in a real browser. Provide no-op
	// setters so assignments (document.cookie = ..., document.title = ...) don't
	// throw.
	noopSetter := func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return nil, nil }
	doc.CreateAttribute("cookie", str(""), noopSetter)
	doc.CreateAttribute("title", str(""), noopSetter)
	doc.CreateAttribute("referrer", str(""), nil)
	// Note: Document.createEvent is implemented natively in the dom package
	// (it overrides a generated stub, which CreateOperation here cannot do).
}

// documentReadyState returns the owning document's real load state.
func documentReadyState[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	doc, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return cbCtx.NewString("complete"), nil
	}
	return cbCtx.NewString(doc.ReadyState()), nil
}

// configureEventEnv adds the legacy init* event initializers to the Event
// prototype (inherited by MouseEvent, UIEvent, ...), so events created via
// document.createEvent can be initialized the old way before being dispatched.
func configureEventEnv[T any](ev js.Class[T]) {
	ev.CreateOperation("initEvent", eventInit)
	ev.CreateOperation("initUIEvent", eventInit)
	ev.CreateOperation("initMouseEvent", eventInit)
	ev.CreateOperation("initKeyboardEvent", eventInit)
	ev.CreateOperation("initCustomEvent", eventInit)
}

// eventInit implements the common prefix (type, bubbles, cancelable) of all the
// legacy init*Event methods; trailing interface-specific arguments are ignored.
func eventInit[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	e, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	args := cbCtx.Args()
	if len(args) > 0 {
		e.Type = args[0].String()
	}
	if len(args) > 1 {
		e.Bubbles = args[1].Boolean()
	}
	if len(args) > 2 {
		e.Cancelable = args[2].Boolean()
	}
	return nil, nil
}

// windowScreen implements the window.screen getter, returning a per-window
// Screen instance.
func windowScreen[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := codec.GetWindow(cbCtx)
	if err != nil {
		return nil, err
	}
	s, ok := entity.ComponentType[*screenObj](win)
	if !ok {
		s = &screenObj{}
		entity.SetComponentType(win, s)
	}
	return codec.EncodeEntityScopedWithPrototype(cbCtx, s, "Screen")
}

// windowPerformance implements the window.performance getter, returning a
// per-window Performance instance.
func windowPerformance[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := codec.GetWindow(cbCtx)
	if err != nil {
		return nil, err
	}
	p, ok := entity.ComponentType[*performanceObj](win)
	if !ok {
		p = &performanceObj{origin: time.Now()}
		entity.SetComponentType(win, p)
	}
	return codec.EncodeEntityScopedWithPrototype(cbCtx, p, "Performance")
}

// windowCrypto implements the window.crypto getter, returning a per-window
// Crypto instance.
func windowCrypto[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win, err := codec.GetWindow(cbCtx)
	if err != nil {
		return nil, err
	}
	c, ok := entity.ComponentType[*cryptoObj](win)
	if !ok {
		c = &cryptoObj{}
		entity.SetComponentType(win, c)
	}
	return codec.EncodeEntityScopedWithPrototype(cbCtx, c, "Crypto")
}
