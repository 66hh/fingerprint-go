package module

import (
	"fingerprint/api"
	"fingerprint/api/console"
)

func Applepay_test() {

	console.Log("applepay test start")

	api.JsEval("window.ApplePaySession = undefined")
	if GetApplePayState() != NoAPI {
		console.Warn("NoAPI Test Error")
	} else {
		console.Log("NoAPI OK")
	}

	api.JsEval("window.ApplePaySession = []")
	api.JsEval("window.ApplePaySession.canMakePayments = function() { return true }")
	if GetApplePayState() != Enabled {
		console.Warn("Enabled Test Error")
	} else {
		console.Log("Enabled OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { return false }")
	if GetApplePayState() != Disabled {
		console.Warn("Disabled Test Error")
	} else {
		console.Log("Disabled OK")
	}

	api.JsEval("window.ApplePaySession = []")
	api.JsEval("window.ApplePaySession.canMakePayments = function() { return 1 }")
	if GetApplePayState() != Enabled {
		console.Warn("Enabled Test Error")
	} else {
		console.Log("Enabled OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { return 0 }")
	if GetApplePayState() != Disabled {
		console.Warn("Disabled Test Error")
	} else {
		console.Log("Disabled OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { throw new Error(\"InvalidAccessError: Trying to call an ApplePaySession API from an insecure document.\"); }")
	if GetApplePayState() != NotAvailableInInsecureContext {
		console.Warn("NotAvailableInInsecureContext Test Error")
	} else {
		console.Log("NotAvailableInInsecureContext OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { throw new Error(\"InvalidAccessError: Trying to start an Apple Pay session from an insecure document.\"); }")
	if GetApplePayState() != NotAvailableInInsecureContext {
		console.Warn("NotAvailableInInsecureContext Test Error")
	} else {
		console.Log("NotAvailableInInsecureContext OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { throw new Error(\"InvalidAccessError: Trying to call an ApplePaySession API from a document with an different security origin than its top-level frame.\"); }")
	if GetApplePayState() != NotAvailableInFrame {
		console.Warn("NotAvailableInFrame Test Error")
	} else {
		console.Log("NotAvailableInFrame OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { throw new Error(\"InvalidAccessError: Trying to start an Apple Pay session from a document with an different security origin than its top-level frame.\"); }")
	if GetApplePayState() != NotAvailableInFrame {
		console.Warn("NotAvailableInFrame Test Error")
	} else {
		console.Log("NotAvailableInFrame OK")
	}

	api.JsEval("window.ApplePaySession.canMakePayments = function() { throw new Error(\"SecurityError: Third-party iframes are not allowed to request payments unless explicitly allowed via Feature-Policy (payment)\"); }")
	if GetApplePayState() != NotAvailableInFrame {
		console.Warn("NotAvailableInFrame Test Error")
	} else {
		console.Log("NotAvailableInFrame OK")
	}

	api.JsEval("window.ApplePaySession = undefined")

	console.Log("applepay end")
}
