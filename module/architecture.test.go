package module

import "fingerprint/api/console"

func Architecture_test() {

	console.Log("architecture test start")

	num := GetArchitecture()

	if num == 127 || num == 255 {
		console.Log("Architecture OK")
	} else {
		console.Warn("Architecture Test Error")
	}

	if num == GetArchitecture() {
		console.Log("Architecture OK")
	} else {
		console.Warn("Architecture Test Error")
	}

	console.Log("architecture test end")

}
