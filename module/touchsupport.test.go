package module

import "fingerprint/api/console"

func TouchSupport_test() {
	console.Log("touchsupport test start")

	result := GetTouchSupport()
	if len(result) != 0 && result == GetTouchSupport() {
		console.Log("TouchSupport OK")
	} else {
		console.Warn("TouchSupport Test Error")
	}

	console.Log("touchsupport test end")
}
