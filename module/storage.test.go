package module

import "fingerprint/api/console"

func Storage_test() {
	console.Log("storage test start")

	result := GetSessionStorage()
	if result == GetSessionStorage() {
		console.Log("Storage OK")
	} else {
		console.Warn("Storage Test Error")
	}

	console.Log("storage test end")
}
