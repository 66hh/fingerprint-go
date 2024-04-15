package module

import "fingerprint/api/console"

func WebGL_test() {
	console.Log("webgl test start")

	result := GetWebGLFingerprint()
	if len(result) != 0 && result == GetWebGLFingerprint() {
		console.Log("WebGL OK")
	} else {
		console.Warn("WebGL Test Error")
	}

	console.Log("webgl test end")
}
