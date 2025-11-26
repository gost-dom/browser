package internal

import (
	"log/slog"

	"github.com/gost-dom/browser/scripting/internal/js"
)

func configureConsole[T any](host js.ScriptEngine[T]) {
	console := host.CreateGlobalObject("console")
	console.CreateFunction("debug", consoleDebug)
	console.CreateFunction("log", consoleLog)
	console.CreateFunction("info", consoleInfo)
	console.CreateFunction("warn", consoleWarn)
	console.CreateFunction("error", consoleError)
	console.CreateFunction("assert", consoleAssert)
}

func consoleDebug[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	return consoleWrite(ctx, slog.LevelDebug)
}
func consoleLog[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	return consoleWrite(ctx, slog.LevelInfo)
}
func consoleInfo[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	return consoleWrite(ctx, slog.LevelInfo)
}
func consoleWarn[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	return consoleWrite(ctx, slog.LevelWarn)
}
func consoleError[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	return consoleWrite(ctx, slog.LevelError)
}

func consoleAssert[T any](ctx js.CallbackContext[T]) (js.Value[T], error) {
	if arg, ok := ctx.ConsumeArg(); ok {
		if !arg.Boolean() {
			return consoleWrite(ctx, slog.LevelError)
		}
	}
	return nil, nil
}

func consoleWrite[T any](ctx js.CallbackContext[T], level slog.Level) (js.Value[T], error) {
	logger := ctx.Logger()
	var rest []js.Value[T]
	var msg string
	if jsMsg, ok := ctx.ConsumeArg(); ok {
		msg = jsMsg.String()
		for arg, ok := ctx.ConsumeArg(); ok; arg, ok = ctx.ConsumeArg() {
			rest = append(rest, arg)
		}
	}
	if rest == nil {
		logger.Log(ctx.Window().Context(), level, msg)
	} else {
		logger.Log(ctx.Window().Context(), level, msg, "args", rest)
	}
	return nil, nil
}
