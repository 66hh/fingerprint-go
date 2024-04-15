package module

import "fingerprint/api/console"

func Canvas_test() {

	console.Log("canvas test start")

	text := GetCanvasFingerprint()
	if text != "error" && len(text) != 0 && text == GetCanvasFingerprint() {
		console.Log("Canvas OK")
	} else {
		console.Warn("Canvas Test Error")
	}

	console.Log("canvas test end")

}
