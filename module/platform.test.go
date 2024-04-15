package module

import "fingerprint/api/console"

func Platform_test() {

	console.Log("platform test start")

	text := GetPlatform()
	if text != "error" && text == GetPlatform() {
		console.Log("Platform OK")
	} else {
		console.Warn("Platform Test Error")
	}

	console.Log("platform test end")

}
