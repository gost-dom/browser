package gojahost

import (
	"fmt"
	"iter"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaValueFactory struct {
	*GojaContext
}

func newGojaValueFactory(c *GojaContext) js.ValueFactory[jsTypeParam] {
	return gojaValueFactory{c}
}

func (f gojaValueFactory) JSONParse(s string) (js.Value[jsTypeParam], error) {
	o := f.vm.NewObject()
	err := o.UnmarshalJSON([]byte(s))
	return newGojaObject(f.GojaContext, o), err
}

func (f gojaValueFactory) JSONStringify(v js.Value[jsTypeParam]) string {
	if o := v.Self().value.ToObject(f.vm); o != nil {
		b, err := o.MarshalJSON()
		if err == nil {
			panic(fmt.Sprintf("Goja JSON marshalling failed: %v", err))
		}
		return string(b)
	}
	panic("Goja only support JSON for objects")
}

func (f gojaValueFactory) NewArray(v ...js.Value[jsTypeParam]) js.Value[jsTypeParam] {
	panic("NewArray not implemented")
}

func (f gojaValueFactory) NewBoolean(v bool) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) Null() js.Value[jsTypeParam] {
	panic("NewArray not implemented")
}

func (f gojaValueFactory) NewUint32(v uint32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewInt32(v int32) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewInt64(v int64) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewString(v string) js.Value[jsTypeParam] {
	return newGojaValue(f.GojaContext, f.vm.ToValue(v))
}

func (f gojaValueFactory) NewTypeError(v string) error {
	panic(f.vm.NewTypeError(v))
}

func (f gojaValueFactory) NewIterator(
	iter.Seq2[js.Value[jsTypeParam], error],
) js.Value[jsTypeParam] {
	panic("NewIterator not implemented")
}
