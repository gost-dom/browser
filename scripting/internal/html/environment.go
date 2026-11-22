package html

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// installEnvironment fills in functional parts of the browser environment that
// the Web IDL code generator does not yet produce: the Performance and Crypto
// interfaces (and their window accessors) and the live document.readyState.
func installEnvironment[T any](e js.ScriptEngine[T]) {
	configurePerformance(js.CreateClass(e, "Performance", "", js.IllegalConstructor))
	configureCrypto(js.CreateClass(e, "Crypto", "", js.IllegalConstructor))

	if win, ok := e.Class("Window"); ok {
		win.CreateAttribute("performance", windowPerformance, nil)
		win.CreateAttribute("crypto", windowCrypto, nil)
	}
	if doc, ok := e.Class("Document"); ok {
		doc.CreateAttribute("readyState", documentReadyState, nil)
	}
}

/* -------------------------------------------------------------------------- */
/* performance                                                                */
/* -------------------------------------------------------------------------- */

type performanceObj struct {
	entity.Entity
	origin time.Time
}

func configurePerformance[T any](performance js.Class[T]) {
	performance.CreateOperation("now", performanceNow)
	performance.CreateAttribute("timeOrigin", performanceTimeOrigin, nil)
}

func performanceNow[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	inst, err := js.As[*performanceObj](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	ms := float64(time.Since(inst.origin).Nanoseconds()) / 1e6
	return cbCtx.NewNumber(ms), nil
}

func performanceTimeOrigin[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	inst, err := js.As[*performanceObj](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.NewNumber(float64(inst.origin.UnixNano()) / 1e6), nil
}

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

/* -------------------------------------------------------------------------- */
/* crypto                                                                     */
/* -------------------------------------------------------------------------- */

type cryptoObj struct{ entity.Entity }

func configureCrypto[T any](crypto js.Class[T]) {
	crypto.CreateOperation("getRandomValues", cryptoGetRandomValues)
	crypto.CreateOperation("randomUUID", cryptoRandomUUID)
}

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
	if isFloatTypedArray(cbCtx, obj) {
		return nil, cbCtx.NewTypeError(
			"Crypto.getRandomValues: argument is not an integer TypedArray")
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
		// The engine truncates the assigned Number to the element width, so a
		// random uint32 yields uniform random bits for 8-, 16- and 32-bit
		// elements.
		v := binary.LittleEndian.Uint32(raw[i*4:])
		if err := obj.Set(strconv.Itoa(i), cbCtx.NewUint32(v)); err != nil {
			return nil, err
		}
	}
	return arg, nil
}

// isFloatTypedArray reports whether obj is a Float16/32/64 typed array, whose
// elements getRandomValues must reject. It compares the array's constructor to
// the global float typed-array constructors. This works regardless of whether
// the engine lets a constructor function be introspected as an object (Sobek
// does not), unlike reading constructor.name.
func isFloatTypedArray[T any](cbCtx js.CallbackContext[T], obj js.Object[T]) bool {
	ctor, err := obj.Get("constructor")
	if err != nil || ctor == nil {
		return false
	}
	globalThis := cbCtx.GlobalThis()
	for _, name := range []string{"Float16Array", "Float32Array", "Float64Array"} {
		floatCtor, err := globalThis.Get(name)
		if err != nil || floatCtor == nil || floatCtor.IsUndefined() {
			continue
		}
		if ctor.StrictEquals(floatCtor) {
			return true
		}
	}
	return false
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
/* document                                                                   */
/* -------------------------------------------------------------------------- */

// documentReadyState returns the owning document's real load state ("loading"
// while parsing, "complete" afterwards) rather than a constant, so scripts that
// defer work until the DOM is ready behave correctly.
func documentReadyState[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	doc, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return cbCtx.NewString("complete"), nil
	}
	return cbCtx.NewString(doc.ReadyState()), nil
}
