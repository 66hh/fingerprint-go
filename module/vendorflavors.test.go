package module

import "fingerprint/api/console"

func VendorFlavors_test() {
	console.Log("vendorflavors test start")

	result := GetVendorFlavors()
	if len(result) != 0 && result == GetVendorFlavors() {
		console.Log("VendorFlavors OK")
	} else {
		console.Warn("VendorFlavors Test Error")
	}

	console.Log("vendorflavors test end")
}
