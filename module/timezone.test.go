package module

import "fingerprint/api/console"

func Timezone_test() {

	console.Log("timezone test start")

	text := GetPlatform()
	if text != "error" && text == GetPlatform() {
		console.Log("Timezone OK")
	} else {
		console.Warn("Timezone Test Error")
	}

	console.Log("timezone test end")

}
