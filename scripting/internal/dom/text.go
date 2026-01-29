package dom

import (
	dom "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func CreateText[T any](s js.CallbackScope[T]) (js.Value[T], error) {
	return CreateTextData(s, "")
}
func CreateTextData[T any](s js.CallbackScope[T], data string) (js.Value[T], error) {
	win, err := codec.GetWindow(s)
	if err != nil {
		return nil, err
	}
	text := dom.NewText(data, win.Document())
	return codec.EncodeConstructedValue(s, text)
}
