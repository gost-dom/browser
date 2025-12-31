// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

type WindowOrWorkerGlobalScope[T any] struct{}

func NewWindowOrWorkerGlobalScope[T any](scriptHost js.ScriptEngine[T]) WindowOrWorkerGlobalScope[T] {
	return WindowOrWorkerGlobalScope[T]{}
}

func (wrapper WindowOrWorkerGlobalScope[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w WindowOrWorkerGlobalScope[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("setTimeout", WindowOrWorkerGlobalScope_setTimeout)
	jsClass.CreateOperation("clearTimeout", WindowOrWorkerGlobalScope_clearTimeout)
	jsClass.CreateOperation("setInterval", WindowOrWorkerGlobalScope_setInterval)
	jsClass.CreateOperation("clearInterval", WindowOrWorkerGlobalScope_clearInterval)
	jsClass.CreateOperation("queueMicrotask", WindowOrWorkerGlobalScope_queueMicrotask)
}
