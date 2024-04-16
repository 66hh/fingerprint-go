package fingerprint

import (
	"fingerprint/api/console"
	"fingerprint/module"
	"fmt"
)

func MakeFingerprint() string {

	var result string = ""

	result = result + module.GetApplePayState().String()
	result = result + "|" + fmt.Sprintf("%d", module.GetArchitecture())
	result = result + "|" + module.GetCanvasFingerprint()
	result = result + "|" + fmt.Sprintf("%d", module.GetColorDepth())
	result = result + "|" + module.GetColorGamut()
	result = result + "|" + module.GetContrastPreference().String()
	result = result + "|" + module.GetWebGLFingerprint()
	result = result + "|" + module.GetVendor()
	result = result + "|" + fmt.Sprintf("%v", module.GetSessionStorage())
	result = result + "|" + module.GetVendorFlavors()
	result = result + "|" + module.GetTouchSupport()
	result = result + "|" + module.GetPlatform()
	result = result + "|" + module.GetTimezone()
	result = result + "|" + module.GetLanguage()

	console.Log(result)

	return X64hash128(result)
}
