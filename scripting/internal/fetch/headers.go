package fetch

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w Headers[T]) CreateInstance(
	cbCtx js.CallbackContext[T], options ...[2]string,
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
) (res [][2]string, err error) {
	if v == nil {
		return nil, nil
	}
	if obj, ok := v.AsObject(); ok {
		var it js.Value[T]
		it, err = obj.Iterator()
		if err != nil {
			return nil, err
		}
		if it != nil && !it.IsUndefined() {
			scope.Logger().
				Info("Iterator", "str", it.String(), "it", it, "obj", obj, "objStr", obj.String())
			return w.parseHeaderIterator(scope, obj, it)
		}
		var keys []string
		if keys, err = obj.Keys(); err == nil {
			res = make([][2]string, len(keys))
			for i, key := range keys {
				var val js.Value[T]
				if val, err = obj.Get(key); err != nil {
					return
				}
				res[i] = [2]string{key, val.String()}
			}
		}
		return
	}
	return nil, nil
}

func (w Headers[T]) parseHeaderIterator(
	scope js.Scope[T],
	arg js.Object[T],
	it js.Value[T],
) ([][2]string, error) {
	itFn, ok := it.AsFunction()
	if !ok {
		return nil, errors.New("Not callable")
	}

	itRes, err := itFn.Call(arg)
	if err != nil {
		return nil, err
	}
	obj, ok := itRes.AsObject()
	if !ok {
		return nil, errors.New("Bad iterator object")
	}

	res := make([][2]string, 0)
	next, err := obj.Get("next")
	if err != nil {
		return nil, err
	}
	if next == nil || next.IsUndefined() {
		keys, _ := obj.Keys()
		return nil, fmt.Errorf("Iterator doesn't have function 'next'. Keys: %+v", keys)
	}
	fn, ok := next.AsFunction()
	if !ok {
		return nil, errors.New("Next is not a function")
	}
	for {
		result, err := fn.Call(obj)
		if err != nil {
			return nil, fmt.Errorf("next: %w", err)
		}
		resultObj, ok := result.AsObject()
		if !ok {
			break
		}
		done, err := resultObj.Get("done")
		if err != nil {
			return nil, fmt.Errorf("done: %w", err)
		}
		if done.Boolean() {
			break
		}
		value, err := resultObj.Get("value")
		if err != nil {
			return nil, fmt.Errorf("value: %w", err)
		}
		valObj, ok := value.AsObject()
		if !ok {
			break
		}
		v1, err1 := valObj.Get("0")
		v2, err2 := valObj.Get("1")
		if err = errors.Join(err1, err2); err != nil {
			return nil, err
		}
		res = append(res, [2]string{
			v1.String(),
			v2.String(),
		})
	}
	return res, nil
}

func (w Headers[T]) CustomInitializer(jsClass js.Class[T]) {
	iterator := js.NewIterator2(codec.EncodeStringScoped[T], codec.EncodeStringScoped[T])
	iterator.InstallPrototype(jsClass)
}
