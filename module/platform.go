package module

import "fingerprint/api"

func GetPlatform() string {

	navigator := api.Global.Get("navigator")

	if api.JsCheck(navigator, api.TypeObject) {
		data := navigator.Get("platform")
		if api.JsCheck(data, api.TypeString) {
			return data.String()
		}
	}

	return "error"
}
