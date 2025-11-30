package fetch

import (
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

func (w Headers[T]) decodeHeadersInit(_ js.Scope[T], v js.Value[T]) (res [][2]string, err error) {
	if v == nil {
		return nil, nil
	}
	if obj, ok := v.AsObject(); ok {
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

func (w Headers[T]) CustomInitializer(jsClass js.Class[T]) {
	iterator := js.NewIterator2(codec.EncodeStringScoped[T], codec.EncodeStringScoped[T])
	iterator.InstallPrototype(jsClass)
}
