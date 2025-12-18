package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeResponse[T any](info js.Scope[T], res *fetch.Response) (js.Value[T], error) {
	return info.Constructor("Response").NewInstance(res)
}

func Fetch[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	info.Logger().Debug("JS Function call: fetch")
	url, err := js.ConsumeArgument(info, "url", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	opts, err := js.ConsumeArgument(
		info,
		"options",
		codec.ZeroValue,
		decodeRequestInit,
	)
	win, err := codec.GetWindow(info)
	if err != nil {
		return nil, err
	}
	f := fetch.New(win)
	info.Logger().Debug("js/fetch: create promise")
	req := f.NewRequest(url, opts...)
	return codec.EncodePromise(info, f.FetchAsync(req), encodeResponse)
}

func decodeRequestInit[T any](
	scope js.Scope[T], val js.Value[T],
) ([]fetch.RequestOption, error) {
	if js.IsUndefined(val) {
		return nil, nil
	}
	options := codec.Options[T, fetch.RequestOption]{
		"signal":  codec.OptDecoder[T](codec.DecodeNativeValue, fetch.WithSignal),
		"method":  codec.OptDecoder[T](codec.DecodeString, fetch.WithMethod),
		"body":    codec.OptDecoder[T](codec.DecodeRequestBody, fetch.WithBody),
		"headers": codec.OptDecoder[T](decodeHeadersInit, fetch.WithHeaders),
	}
	for _, optName := range missingRequestOptions {
		options[optName] = func(js.Scope[T], js.Value[T]) (fetch.RequestOption, error) {
			return nil, fmt.Errorf(
				"gost-dom/fetch: decode RequestInit: %s option not yet supported. %s",
				optName,
				constants.MISSING_FEATURE_ISSUE_URL,
			)
		}
	}
	return codec.DecodeOptions(scope, val, options)
}

var missingRequestOptions = []string{
	"referrer", "referrerPolicy",
	"mode", "credentials", "cache", "redirect", "integrity",
	"keepalive", "duplex", "priority", "window",
}
