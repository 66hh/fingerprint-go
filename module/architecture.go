package module

import (
	"fingerprint/api"
)

func GetArchitecture() int {

	var result int = -1

	array := api.Global.Get("Float32Array")

	if api.JsCheck(array, api.TypeFunction) {

		floatArray := array.New(1)

		array = api.Global.Get("Uint8Array")

		if api.JsCheck(array, api.TypeFunction) {
			u8Arry := array.New(floatArray.Get("buffer"))

			floatArray.SetIndex(0, api.Global.Get("Infinity"))

			api.Global.Set("arc", floatArray)
			api.JsEval("arc[0] = arc[0] - arc[0]")

			obj := u8Arry.Index(3)

			if api.JsCheck(obj, api.TypeNumber) {
				result = obj.Int()
			}

			api.Global.Set("arc", nil)
		}
	}

	return result
}
