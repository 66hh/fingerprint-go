package main

import (
	"fingerprint/api/console"
	"fingerprint/fingerprint"
	"fingerprint/module"
)

func test() {
	module.Applepay_test()
	module.Architecture_test()
	module.Colordepth_test()
	module.Canvas_test()
	module.Colorgamut_test()
	module.Contrast_test()
	module.WebGL_test()
	module.Navigator_test()
	module.Storage_test()
	module.VendorFlavors_test()
	module.TouchSupport_test()
	module.Platform_test()
	module.Timezone_test()
	module.Language_test()

	console.Log("result: " + fingerprint.MakeFingerprint())
}

func main() {
	test()
}
