package fetch

import (
	"errors"

	"github.com/gost-dom/browser/internal/fetch"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	"github.com/gost-dom/browser/internal/types"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w Headers[T]) CreateInstance(
	cbCtx js.CallbackContext[T], options ...[2]types.ByteString,
) (js.Value[T], error) {
	res := fetch.Headers{}
	for _, h := range options {
		res.Append(h[0], h[1])
	}
	return codec.EncodeConstrucedValue(cbCtx, res)
}

func (w Headers[T]) decodeHeadersInit(
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
	res, err = w.parseHeaderIterator2(scope, v)
	if err == nil || (!errors.Is(err, js.ErrNotIterable)) {
		return
	}
	if obj, ok := v.AsObject(); ok {
		var keys js.Value[T]
		keys, err = js.ObjectKeys(scope, obj)
		if err != nil {
			return nil, scope.NewTypeError(err.Error())
		}
		for key, err := range js.Iterate(keys) {
			if err != nil {
				return nil, err
			}
			var desc js.PropertyDescriptor[T]
			desc, err = js.ObjectOwnPropertyDescriptor(scope, obj, key)
			if err != nil {
				return nil, err
			}
			if desc != nil && desc.Enumerable() {
				if !key.IsString() {
					return nil, scope.NewTypeError("Non-string key")
				}
				var item [2]types.ByteString
				var val js.Value[T]
				if item[0], err = types.ToByteString(key.String()); err != nil {
					return nil, err
				}
				val, err = obj.Get(key.String())
				if err == nil {
					item[1], err = types.ToByteString(val.String())
				}
				if err != nil {
					return nil, err
				}
				res = append(res, item)
			}
		}
		return
	}
	return nil, nil
}

func (w Headers[T]) parseHeaderIterator2(
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
