package module

import "fingerprint/api/console"

func Language_test() {
	console.Log("language test start")

	result := GetLanguage()
	if result != "error" && len(result) != 0 && result == GetLanguage() {
		console.Log("Language OK")
	} else {
		console.Warn("Language Test Error")
	}

	console.Log("language test end")
}
