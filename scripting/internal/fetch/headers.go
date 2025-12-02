package fetch

import (
	"errors"
	"iter"

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
		var entries js.Value[T]
		entries, err = js.ObjectEntries(scope, obj)
		if err != nil {
			return nil, err
		}
		res, err = w.parseHeaderIterator2(scope, entries)
		// var keys []string
		// if keys, err = obj.Keys(); err == nil {
		// 	res = make([][2]types.ByteString, len(keys))
		// 	for i, key := range keys {
		// 		if res[i][0], err = types.ToByteString(key); err != nil {
		// 			return
		// 		}
		// 		var val js.Value[T]
		// 		if val, err = obj.Get(key); err != nil {
		// 			return
		// 		}
		// 		if res[i][1], err = codec.DecodeByteString(scope, val); err != nil {
		// 			return
		// 		}
		// 	}
		// }
		return
	}
	return nil, nil
}

func (w Headers[T]) parseHeaderIterator2(
	scope js.Scope[T], val js.Value[T],
) (res [][2]types.ByteString, err error) {
	var seq iter.Seq2[js.Value[T], error]
	if seq, err = js.Iterate(val); err != nil {
		return
	}
	for v, err := range seq {
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

/*
func (w Headers[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	var init [][2]string
	if arg, ok := cbCtx.ConsumeArg(); ok {
		init, err = w.decodeHeadersInit(cbCtx, arg)
		if err != nil {
			return
		}
	}
	return w.CreateInstance(cbCtx, init...)
}
*/
