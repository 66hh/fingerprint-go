package module

import "fingerprint/api/console"

func Colordepth_test() {

	console.Log("colordepth test start")

	num := GetColorDepth()
	if num == 48 || num == 36 || num == 32 || num == 30 || num == 24 || num == 18 || num == 16 || num == 15 || num == 12 || num == 8 || num == 0 {
		console.Log("Colordepth OK")
	} else {
		console.Warn("Colordepth Test Error")
	}

	console.Log("colordepth test end")
}
