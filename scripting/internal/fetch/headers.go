package fetch

import (
	"errors"

	"github.com/gost-dom/browser/internal/fetch"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func CreateHeaders[T any](cbCtx js.CallbackContext[T], options ...[2]types.ByteString,
) (js.Value[T], error) {
	res := &fetch.Headers{}
	for _, h := range options {
		res.Append(h[0], h[1])
	}
	return codec.EncodeConstructedValue(cbCtx, res)
}

func (w Headers[T]) decodeHeadersInit(
	s js.Scope[T],
	v js.Value[T],
) (res [][2]types.ByteString, err error) {
	return decodeHeadersInit(s, v)
}

func decodeHeadersInit[T any](
	s js.Scope[T],
	v js.Value[T],
) (res [][2]types.ByteString, err error) {
	scope := s.(js.CallbackScope[T])
	if v == nil || v.IsUndefined() {
		return nil, nil
	}
	if v.IsNull() {
		return nil, scope.NewTypeError("invalid value: null")
	}
	res, err = parseHeaderIterator2(scope, v)
	if err == nil || (!errors.Is(err, js.ErrNotIterable)) {
		return
	}
	if obj, ok := v.AsObject(); ok {
		var key js.Value[T]
		for key, err = range js.ObjectEnumerableOwnPropertyKeys(scope, obj) {
			if err == nil && key.IsSymbol() {
				err = scope.NewTypeError("Non-string key")
			}
			if err != nil {
				break
			}

			var item [2]types.ByteString
			if item[0], err = types.ToByteString(key.String()); err != nil {
				return nil, err
			}

			var val js.Value[T]
			if val, err = obj.Get(key.String()); err == nil {
				item[1], err = types.ToByteString(val.String())
			}
			if err != nil {
				return
			}
			res = append(res, item)
		}
	}
	return
}

func parseHeaderIterator2[T any](
	scope js.Scope[T], val js.Value[T],
) (res [][2]types.ByteString, err error) {
	for v, err := range js.Iterate(val) {
		if err != nil {
			return nil, err
		}
		obj, ok := v.AsObject()
		if !ok {
			return nil, scope.NewTypeError("Not an object")
		}
		v1, err1 := obj.Get("0")
		v2, err2 := obj.Get("1")
		s1, err3 := codec.DecodeByteString(scope, v1)
		s2, err4 := codec.DecodeByteString(scope, v2)
		if err = gosterror.First(err1, err2, err3, err4); err != nil {
			return nil, err
		}
		res = append(res, [2]types.ByteString{s1, s2})
	}
	return
}

func (w Headers[T]) CustomInitializer(jsClass js.Class[T]) {
	iterator := js.NewIterator2(codec.EncodeByteString[T], codec.EncodeByteString[T])
	iterator.InstallPrototype(jsClass)
}
