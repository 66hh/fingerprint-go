package module

import (
	"fingerprint/api"
	"fmt"
)

func GetTouchSupport() string {

	var maxTouchPoints int = -1
	var touchEvent bool = false
	var touchStart bool = false

	navigator := api.Global.Get("navigator")

	if api.JsCheck(navigator, api.TypeObject) {

		result := navigator.Get("maxTouchPoints")
		if api.JsCheck(result, api.TypeNumber) {
			maxTouchPoints = result.Int()
		}
		result = navigator.Get("msMaxTouchPoints")
		if api.JsCheck(result, api.TypeNumber) {
			maxTouchPoints = result.Int()
		}

		doc := api.Global.Get("document")
		if api.JsCheck(doc, api.TypeObject) {
			_, err := api.JsCheckAndCall(doc, "createEvent", "TouchEvent")
			if err == nil {
				touchEvent = true
			}
		}

		window := api.Global.Get("window")
		if api.JsCheck(window, api.TypeObject) {
			if !api.JsCheck(window.Get("ontouchstart"), api.TypeUndefined) {
				touchStart = true
			}
		}

	}

	return fmt.Sprintf("%v|%v|%v", maxTouchPoints, touchEvent, touchStart)

}
