package module

import "fingerprint/api/console"

func Colorgamut_test() {

	console.Log("colorgamut test start")

	text := GetColorGamut()
	if text != "error" && text == GetColorGamut() {
		console.Log("Colorgamut OK")
	} else {
		console.Warn("Colorgamut Test Error")
	}

	console.Log("colorgamut test end")

}
