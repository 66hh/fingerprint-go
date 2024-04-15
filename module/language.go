package module

import (
	"fingerprint/api"
)

func GetLanguage() string {

	navigator := api.Global.Get("navigator")

	if api.JsCheck(navigator, api.TypeObject) {
		result := navigator.Get("language")
		if api.JsCheck(result, api.TypeString) {
			return result.String()
		}
	}

	return "error"
}
