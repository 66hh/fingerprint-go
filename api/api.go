package api

import (
	"errors"
	"fmt"
	"syscall/js"
)

var (
	Global js.Value
)

type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string {
	switch t {
	case TypeUndefined:
		return "undefined"
	case TypeNull:
		return "null"
	case TypeBoolean:
		return "boolean"
	case TypeNumber:
		return "number"
	case TypeString:
		return "string"
	case TypeSymbol:
		return "symbol"
	case TypeObject:
		return "object"
	case TypeFunction:
		return "function"
	default:
		return "unknown"
	}
}

func JsEval(code string) (result js.Value, err error) {

	result, err = JsCheckAndCall(Global, "eval", code)

	return
}

func JsCheck(value js.Value, valuetype Type) bool {
	if value.Type().String() == valuetype.String() {
		return true
	}
	return false
}

func JsCheckAndCall(object js.Value, name string, param ...any) (result js.Value, err error) {

	if JsCheck(object.Get(name), TypeFunction) {
		defer func() {
			if e := recover(); e != nil {
				err = errors.New(fmt.Sprintf("%v", e))
			}
		}()
		result = object.Call(name, param...)
	} else {
		err = errors.New("invalid type")
	}

	return
}

func init() {
	Global = js.Global()
}
