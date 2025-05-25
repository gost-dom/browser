package v8host

import (
	"github.com/gost-dom/browser/dom"
	v8 "github.com/gost-dom/v8go"
)

/*
Instance properties

    length

Instance methods

    entries()  // returns an iterator
    forEach()  // calls a callback
    item()
    keys()
    values()


## Foreach parameters

    // foreach Behaviour in FF (by experimenting)
    // for (i = 0; i < length; i++) {
    //   element = item(i)
    //   if (element) { callback(element, i) }
    // }
    // Inserting an element _before_ current element will iterate current
// element twice (it has a new index), but last element isn't iterated.
    // Removing an element, and it doesn't iterate _past_ the end of the element

callback

    A function to execute on each element of someNodeList. It accepts 3 parameters:

    currentValue

        The current element being processed in someNodeList.
    currentIndex Optional

        The index of the currentValue being processed in someNodeList.
    listObj Optional

        The someNodeList that forEach() is being applied to.

thisArg Optional

    Value to use as this when executing callback.
*/

func createNodeList(host *V8ScriptHost) *v8.FunctionTemplate {
	nodeListIterator := newIterator[dom.Node](
		host,
		func(instance dom.Node, ctx *V8ScriptContext) (*v8.Value, error) {
			return ctx.getInstanceForNode(instance)
		},
	)
	iso := host.iso
	builder := newIllegalConstructorBuilder[dom.NodeList](host)
	builder.SetDefaultInstanceLookup()
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp2(
		"length",
		func(instance dom.NodeList, ctx *V8ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, uint32(instance.Length()))
		},
	)
	proto.CreateFunction(
		"item",
		func(instance dom.NodeList, info *argumentHelper) (*v8.Value, error) {
			index, err := info.consumeInt32()
			if err != nil {
				return nil, v8.NewTypeError(iso, "Index must be an integer")
			}
			result := instance.Item(int(index))
			if result == nil {
				return v8.Null(iso), nil
			}
			return info.ctx.getInstanceForNode(result)
		},
	)
	instanceTemplate := builder.NewInstanceBuilder().proto
	instanceTemplate.SetSymbol(
		v8.SymbolIterator(iso),
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				nodeList, err := getInstanceFromThis[dom.NodeList](ctx, info.This())
				if err != nil {
					return nil, err
				}
				return nodeListIterator.newIteratorInstance(ctx, nodeList.All())
			},
		),
	)
	instanceTemplate.SetIndexedHandler(
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			instance, ok := ctx.getCachedNode(info.This())
			nodemap, ok_2 := instance.(dom.NodeList)
			if ok && ok_2 {
				index := int(info.Index())
				item := nodemap.Item(index)
				if item == nil {
					return v8.Undefined(iso), nil
				}
				return ctx.getInstanceForNode(item)
			}
			return nil, v8.NewTypeError(iso, "dunno")
		},
	)

	return builder.constructor
}
