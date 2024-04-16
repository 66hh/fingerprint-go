package module

import (
	"fingerprint/api"
)

func GetTimezone() string {

	Intl := api.Global.Get("Intl")

	if api.JsCheck(Intl, api.TypeObject) {
		data, _ := api.JsCheckAndCall(Intl, "DateTimeFormat")
		data, _ = api.JsCheckAndCall(data, "resolvedOptions")
		data = data.Get("timeZone")

		if api.JsCheck(data, api.TypeString) {
			return data.String()
		}
	}

	return "error"
}
