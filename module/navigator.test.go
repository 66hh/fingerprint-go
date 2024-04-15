package module

import "fingerprint/api/console"

func Navigator_test() {
	console.Log("navigator test start")

	result := GetVendor()
	if result != "error" && len(result) != 0 && result == GetVendor() {
		console.Log("Navigator OK")
	} else {
		console.Warn("Navigator Test Error")
	}

	console.Log("navigator test end")
}
