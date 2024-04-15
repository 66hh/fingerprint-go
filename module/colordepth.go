package module

import "fingerprint/api"

func GetColorDepth() int {

	object := api.Global.Get("window")

	if api.JsCheck(object, api.TypeObject) {

		object = object.Get("screen")
		if api.JsCheck(object, api.TypeObject) {

			object = object.Get("colorDepth")
			if api.JsCheck(object, api.TypeNumber) {
				return object.Int()
			}

		}
	}

	return -1
}
