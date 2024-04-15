package module

import "fingerprint/api"

func GetSessionStorage() bool {

	object := api.Global.Get("window")

	if api.JsCheck(object, api.TypeObject) {
		if api.JsCheck(object.Get("sessionStorage"), api.TypeObject) {
			return true
		}
	}

	return false
}
