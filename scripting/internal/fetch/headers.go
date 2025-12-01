package fetch

import (
	"errors"
	"fmt"
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
	scope js.Scope[T],
	v js.Value[T],
) (res [][2]types.ByteString, err error) {
	if v == nil || v.IsUndefined() {
		return nil, nil
	}
	if v.IsNull() {
		return nil, scope.NewTypeError("invalid value: null")
	}
	res, err = w.parseHeaderIterator2(scope, v)
	if err == nil || (!errors.Is(err, ErrNotIterable)) {
		return
	}
	if obj, ok := v.AsObject(); ok {
		var keys []string
		if keys, err = obj.Keys(); err == nil {
			res = make([][2]types.ByteString, len(keys))
			for i, key := range keys {
				if res[i][0], err = types.ToByteString(key); err != nil {
					return
				}
				var val js.Value[T]
				if val, err = obj.Get(key); err != nil {
					return
				}
				if res[i][1], err = codec.DecodeByteString(scope, val); err != nil {
					return
				}
			}
		}
		return
	}
	return nil, nil
}

func (w Headers[T]) parseHeaderIterator2(
	scope js.Scope[T], val js.Value[T],
) (res [][2]types.ByteString, err error) {
	var seq iter.Seq2[js.Value[T], error]
	if seq, err = iterate(val); err != nil {
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

// iterate returns a seq.Iter2 exposing a JavaScript iterable as a Go iterator.
// It will return an ErrNotIterable error if the JavaScript value is not an
// object implementing the [Iterable] protocol. An error is returned if
// obtaining the [Iterator] itself resulted in an error. The returned Seq will
// return yield an error value if the JavaScript iterator throws an error during
// iteration.
//
// [Iterable]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterable_protocol
// [Iterator]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols#the_iterator_protocol
func iterate[T any](v js.Value[T]) (iter.Seq2[js.Value[T], error], error) {
	obj, ok := v.AsObject()
	if !ok {
		return nil, ErrNotIterable
	}
	symIter, ok, err := js.ObjectGetIterator(obj)
	if err == nil && !ok {
		err = ErrNotIterable
	}
	if err != nil {
		return nil, err
	}
	iterVal, err := symIter.Call(obj)
	if err != nil {
		return nil, err
	}
	iter, ok := iterVal.AsObject()
	if !ok {
		return nil, ErrNotIterable
	}
	next, ok, err := js.ObjectGetFunction(iter, "next")
	if err == nil && !ok {
		err = ErrNotIterable
	}
	if err != nil {
		return nil, errNotIterable("next is not a function")
	}
	return func(yield func(js.Value[T], error) bool) {
		for {
			result, err := next.Call(iter)
			if err != nil {
				yield(nil, err)
				return
			}
			resultObj, ok := result.AsObject()
			if !ok {
				break
			}
			done, err := resultObj.Get("done")
			if err != nil {
				yield(nil, err)
				return
			}
			if done.Boolean() {
				return
			}
			val, err := resultObj.Get("value")
			if err != nil || !yield(val, err) {
				return
			}
		}
	}, nil
}

var ErrNotIterable = errors.New("gost-dom/scripting: value not iterable")

func errNotIterable(msg string) error {
	return fmt.Errorf("%w: %s", ErrNotIterable, msg)
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
