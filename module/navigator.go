package module

import (
	"fingerprint/api"
)

func GetVendor() string {

	navigator := api.Global.Get("navigator")

	if api.JsCheck(navigator, api.TypeObject) {
		result := navigator.Get("vendor")
		if api.JsCheck(result, api.TypeString) {
			return result.String()
		}
	}

	return "error"
}
