package module

import "fingerprint/api"

type ContrastPreference int

const (
	Less         ContrastPreference = -1
	None         ContrastPreference = 0
	More         ContrastPreference = 1
	ForcedColors ContrastPreference = 10
)

func (s ContrastPreference) String() string {
	switch s {
	case Less:
		return "Less"
	case None:
		return "None"
	case More:
		return "More"
	case ForcedColors:
		return "ForcedColors"
	default:
		return "unknown"
	}
}

func doesMatch(value string) bool {
	object := api.Global.Get("window")

	if api.JsCheck(object, api.TypeObject) {

		data, err := api.JsCheckAndCall(object, "matchMedia", "(prefers-contrast: "+value+")")

		if err == nil {
			result := data.Get("matches")
			if api.JsCheck(result, api.TypeBoolean) {
				return result.Bool()
			}
		}
	}

	return false
}

func GetContrastPreference() ContrastPreference {
	if doesMatch("no-preference") {
		return None
	}
	if doesMatch("high") || doesMatch("more") {
		return More
	}
	if doesMatch("low") || doesMatch("less") {
		return Less
	}
	if doesMatch("no-forced") {
		return ForcedColors
	}
	return None
}
