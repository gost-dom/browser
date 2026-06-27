package html

import (
	"github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// installNavigator adds the data attributes of the Navigator interface. Each
// value is read from the per-window [html.Navigator] instance, so the reported
// profile can be configured per browser via the browser package's WithNavigator
// option.
func installNavigator[T any](e js.ScriptEngine[T]) {
	nav, ok := e.Class("Navigator")
	if !ok {
		return
	}
	nav.CreateAttribute("userAgent", navigatorUserAgent, nil)
	nav.CreateAttribute("platform", navigatorPlatform, nil)
	nav.CreateAttribute("vendor", navigatorVendor, nil)
	nav.CreateAttribute("language", navigatorLanguage, nil)
	nav.CreateAttribute("languages", navigatorLanguages, nil)
	nav.CreateAttribute("hardwareConcurrency", navigatorHardwareConcurrency, nil)
	nav.CreateAttribute("webdriver", navigatorWebdriver, nil)
}

func navigatorInstance[T any](cbCtx js.CallbackContext[T]) (*html.Navigator, error) {
	return js.As[*html.Navigator](cbCtx.Instance())
}

func navigatorUserAgent[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewString(nav.UserAgent()), nil
}

func navigatorPlatform[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewString(nav.Platform()), nil
}

func navigatorVendor[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewString(nav.Vendor()), nil
}

func navigatorLanguage[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewString(nav.Language()), nil
}

func navigatorLanguages[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	langs := nav.Languages()
	vals := make([]js.Value[T], len(langs))
	for i, l := range langs {
		vals[i] = cbCtx.NewString(l)
	}
	return cbCtx.NewArray(vals...), nil
}

func navigatorHardwareConcurrency[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewInt32(int32(nav.HardwareConcurrency())), nil
}

func navigatorWebdriver[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	nav, err := navigatorInstance(cbCtx)
	if err != nil {
		return nil, err
	}
	return cbCtx.NewBoolean(nav.Webdriver()), nil
}
