// This file is generated. Do not edit.

package html

import js "github.com/gost-dom/browser/scripting/internal/js"

func InitializeWindowOrWorkerGlobalScope[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("setTimeout", WindowOrWorkerGlobalScope_setTimeout)
	jsClass.CreateOperation("clearTimeout", WindowOrWorkerGlobalScope_clearTimeout)
	jsClass.CreateOperation("setInterval", WindowOrWorkerGlobalScope_setInterval)
	jsClass.CreateOperation("clearInterval", WindowOrWorkerGlobalScope_clearInterval)
	jsClass.CreateOperation("queueMicrotask", WindowOrWorkerGlobalScope_queueMicrotask)
}
