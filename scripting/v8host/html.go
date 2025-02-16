package v8host

import (
	. "github.com/gost-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type esElementContainerWrapper[T ElementContainer] struct {
	nodeV8WrapperBase[T]
	parentNode *parentNodeV8Wrapper
}

func newESContainerWrapper[T ElementContainer](host *V8ScriptHost) esElementContainerWrapper[T] {
	return esElementContainerWrapper[T]{newNodeV8WrapperBase[T](host),
		newParentNodeV8Wrapper(host),
	}
}

func (e esElementContainerWrapper[T]) Install(ft *v8.FunctionTemplate) {
	prototype := ft.PrototypeTemplate()
	e.parentNode.installPrototype(prototype)
}
