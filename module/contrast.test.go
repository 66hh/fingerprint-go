package module

import "fingerprint/api/console"

func Contrast_test() {
	console.Log("contrast test start")

	result := GetContrastPreference()
	if result == GetContrastPreference() {
		console.Log("Contrast OK")
	} else {
		console.Warn("Contrast Test Error")
	}

	console.Log("contrast test end")
}
