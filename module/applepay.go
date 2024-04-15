package module

import (
	"fingerprint/api"
	"fingerprint/utils"
)

type ApplePayState int

const (
	Disabled                      ApplePayState = 0
	Enabled                       ApplePayState = 1
	NoAPI                         ApplePayState = -1
	NotAvailableInInsecureContext ApplePayState = -2
	NotAvailableInFrame           ApplePayState = -3
)

func (s ApplePayState) String() string {
	switch s {
	case Disabled:
		return "Disabled"
	case Enabled:
		return "Enabled"
	case NoAPI:
		return "NoAPI"
	case NotAvailableInInsecureContext:
		return "NotAvailableInInsecureContext"
	case NotAvailableInFrame:
		return "NotAvailableInFrame"
	default:
		return "unknown"
	}
}

func GetApplePayState() ApplePayState {

	object := api.Global.Get("window")
	if api.JsCheck(object, api.TypeObject) {

		object = object.Get("ApplePaySession")
		if api.JsCheck(object, api.TypeObject) {

			data, err := api.JsCheckAndCall(object, "canMakePayments")
			if err != nil {
				text := err.Error()

				if utils.FindText(text, "InvalidAccessError") {
					if utils.FindText(text, "from") && utils.FindText(text, "insecure") {
						return NotAvailableInInsecureContext
					}
					if utils.FindText(text, "different") && utils.FindText(text, "origin") && utils.FindText(text, "top") && utils.FindText(text, "level") && utils.FindText(text, "frame") {
						return NotAvailableInFrame
					}
				}

				if utils.FindText(text, "SecurityError") {
					if utils.FindText(text, "third") && utils.FindText(text, "party") && utils.FindText(text, "iframes") && utils.FindText(text, "not") && utils.FindText(text, "allowed") {
						return NotAvailableInFrame
					}
				}

			}

			if api.JsCheck(data, api.TypeBoolean) {
				if data.Bool() {
					return Enabled
				} else {
					return Disabled
				}
			}

			if api.JsCheck(data, api.TypeNumber) {
				if data.Int() == 1 {
					return Enabled
				} else {
					return Disabled
				}
			}

		}
	}

	return NoAPI
}
