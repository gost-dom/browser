package codec

import (
	"errors"

	"github.com/gost-dom/browser/scripting/internal/js"
)

func DecodeOption[T, U, V any](
	_ js.Scope[T],
	v js.Value[T],
	key string,
	encode func(U) V,
) (res V, ok bool, err error) {
	var obj js.Object[T]
	if obj, ok = v.AsObject(); !ok {
		return
	}
	var opt js.Value[T]
	if opt, err = obj.Get(key); err != nil {
		return
	}
	optObj, ok := opt.AsObject()
	if ok {
		native, succ := optObj.NativeValue().(U)
		if succ {
			res = encode(native)
		} else {
			err = errors.New("Option is not correct type")
		}
	}
	return
}
