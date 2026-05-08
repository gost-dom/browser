// This file is generated. Do not edit.

package dom

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeRange[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("setStart", Range_setStart)
	jsClass.CreateOperation("setEnd", Range_setEnd)
	jsClass.CreateOperation("setStartBefore", Range_setStartBefore)
	jsClass.CreateOperation("setStartAfter", Range_setStartAfter)
	jsClass.CreateOperation("setEndBefore", Range_setEndBefore)
	jsClass.CreateOperation("setEndAfter", Range_setEndAfter)
	jsClass.CreateOperation("collapse", Range_collapse)
	jsClass.CreateOperation("selectNode", Range_selectNode)
	jsClass.CreateOperation("selectNodeContents", Range_selectNodeContents)
	jsClass.CreateOperation("compareBoundaryPoints", Range_compareBoundaryPoints)
	jsClass.CreateOperation("deleteContents", Range_deleteContents)
	jsClass.CreateOperation("extractContents", Range_extractContents)
	jsClass.CreateOperation("cloneContents", Range_cloneContents)
	jsClass.CreateOperation("insertNode", Range_insertNode)
	jsClass.CreateOperation("surroundContents", Range_surroundContents)
	jsClass.CreateOperation("cloneRange", Range_cloneRange)
	jsClass.CreateOperation("detach", Range_detach)
	jsClass.CreateOperation("isPointInRange", Range_isPointInRange)
	jsClass.CreateOperation("comparePoint", Range_comparePoint)
	jsClass.CreateOperation("intersectsNode", Range_intersectsNode)
	jsClass.CreateOperation("toString", Range_toString)
	jsClass.CreateAttribute("commonAncestorContainer", Range_commonAncestorContainer, nil)
}

func Range_setStart[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetStart(node, offset)
	return nil, nil
}

func Range_setEnd[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetEnd(node, offset)
	return nil, nil
}

func Range_setStartBefore[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SetStartBefore(node)
	return nil, nil
}

func Range_setStartAfter[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SetStartAfter(node)
	return nil, nil
}

func Range_setEndBefore[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SetEndBefore(node)
	return nil, nil
}

func Range_setEndAfter[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SetEndAfter(node)
	return nil, nil
}

func Range_collapse[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	toStart, errArg1 := js.ConsumeArgument(cbCtx, "toStart", codec.ZeroValue, codec.DecodeBoolean)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Collapse(toStart)
	return nil, nil
}

func Range_selectNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SelectNode(node)
	return nil, nil
}

func Range_selectNodeContents[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SelectNodeContents(node)
	return nil, nil
}

func Range_compareBoundaryPoints[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	how, errArg1 := js.ConsumeArgument(cbCtx, "how", nil, codec.DecodeInt)
	sourceRange, errArg2 := js.ConsumeArgument(cbCtx, "sourceRange", nil, decodeRange)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.CompareBoundaryPoints(how, sourceRange)
	return codec.EncodeInt(cbCtx, result)
}

func Range_deleteContents[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.DeleteContents()
	return nil, nil
}

func Range_extractContents[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ExtractContents()
	return codec.EncodeEntity(cbCtx, result)
}

func Range_cloneContents[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CloneContents()
	return codec.EncodeEntity(cbCtx, result)
}

func Range_insertNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.InsertNode(node)
	return nil, nil
}

func Range_surroundContents[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	newParent, errArg1 := js.ConsumeArgument(cbCtx, "newParent", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.SurroundContents(newParent)
	return nil, nil
}

func Range_cloneRange[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CloneRange()
	return encodeRange(cbCtx, result)
}

func Range_detach[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Detach()
	return nil, nil
}

func Range_isPointInRange[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.IsPointInRange(node, offset)
	return codec.EncodeBoolean(cbCtx, result)
}

func Range_comparePoint[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	offset, errArg2 := js.ConsumeArgument(cbCtx, "offset", nil, codec.DecodeInt)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.ComparePoint(node, offset)
	return codec.EncodeInt(cbCtx, result)
}

func Range_intersectsNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.Range](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.IntersectsNode(node)
	return codec.EncodeBoolean(cbCtx, result)
}

func Range_toString[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.String()
	return codec.EncodeString(cbCtx, result)
}

func Range_commonAncestorContainer[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.Range](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CommonAncestorContainer()
	return codec.EncodeEntity(cbCtx, result)
}
