package console

import (
	"errors"
	"fingerprint/api"
	"syscall/js"
)

var (
	checked bool = false
	console js.Value
)

func init() {
	console = api.Global.Get("console")
	if api.JsCheck(console, api.TypeObject) {
		checked = true
	}
}

func Log(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "log", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}

func Debug(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "debug", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}

func Info(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "info", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}

func Warn(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "warn", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}

func Error(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "error", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}

func Clear(text ...any) (result js.Value, err error) {

	if checked {
		result, err = api.JsCheckAndCall(console, "clear", text...)
	} else {
		err = errors.New("invalid object")
	}

	return
}
