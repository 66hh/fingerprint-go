package module

import (
	"fingerprint/api"
)

func GetColorGamut() string {

	object := api.Global.Get("window")

	if api.JsCheck(object, api.TypeObject) {

		var gamuts []string = []string{"rec2020", "p3", "srgb"}

		for _, value := range gamuts {
			data, err := api.JsCheckAndCall(object, "matchMedia", "(color-gamut: "+value+")")
			if err == nil {
				result := data.Get("matches")
				if api.JsCheck(result, api.TypeBoolean) {
					if result.Bool() {
						return value
					}
				}
			}
		}

	}

	return "error"
}
