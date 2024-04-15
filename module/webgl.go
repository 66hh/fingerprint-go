package module

import (
	"fingerprint/api"
	"syscall/js"
)

func getWebGLEngine() js.Value {
	var gl js.Value

	doc := api.Global.Get("document")
	if api.JsCheck(doc, api.TypeObject) {
		canvas, err := api.JsCheckAndCall(doc, "createElement", "canvas")
		if err == nil {
			gl, err = api.JsCheckAndCall(canvas, "getContext", "webgl")
			if err != nil {
				gl, err = api.JsCheckAndCall(canvas, "getContext", "experimental-webgl")
				if err != nil {
					gl = js.Null()
				}
			}
		} else {
			gl = js.Null()
		}
	}
	return gl
}

func GetWebGLFingerprint() string {

	var result string = ""

	webglContext := getWebGLEngine()

	if webglContext.IsNull() {
		return ""
	}

	vShaderTemplate := "attribute vec2 attrVertex;varying vec2 varyinTexCoordinate;uniform vec2 uniformOffset;void main(){varyinTexCoordinate=attrVertex+uniformOffset;gl_Position=vec4(attrVertex,0,1);}"
	fShaderTemplate := "precision mediump float;varying vec2 varyinTexCoordinate;void main() {gl_FragColor=vec4(varyinTexCoordinate,0,1);}"

	vertexPosBuffer, _ := api.JsCheckAndCall(webglContext, "createBuffer")
	api.JsCheckAndCall(webglContext, "bindBuffer", webglContext.Get("ARRAY_BUFFER"), vertexPosBuffer)

	vertices := []float32{-0.2, -0.9, 0, 0.4, -0.26, 0, 0, 0.732134444, 0}
	array := make([]interface{}, len(vertices))
	for i, v := range vertices {
		array[i] = v
	}
	jsVertices := js.ValueOf(array)

	floatarr := api.Global.Get("Float32Array")

	if api.JsCheck(floatarr, api.TypeFunction) {
		float32Array := floatarr.New(jsVertices)
		api.JsCheckAndCall(webglContext, "bufferData", webglContext.Get("ARRAY_BUFFER"), float32Array, webglContext.Get("STATIC_DRAW"))

		program, _ := api.JsCheckAndCall(webglContext, "createProgram")
		vshader, _ := api.JsCheckAndCall(webglContext, "createShader", webglContext.Get("VERTEX_SHADER"))
		api.JsCheckAndCall(webglContext, "shaderSource", vshader, vShaderTemplate)
		api.JsCheckAndCall(webglContext, "compileShader", vshader)
		fshader, _ := api.JsCheckAndCall(webglContext, "createShader", webglContext.Get("FRAGMENT_SHADER"))
		api.JsCheckAndCall(webglContext, "shaderSource", fshader, fShaderTemplate)
		api.JsCheckAndCall(webglContext, "compileShader", fshader)

		api.JsCheckAndCall(webglContext, "attachShader", program, vshader)
		api.JsCheckAndCall(webglContext, "attachShader", program, fshader)
		api.JsCheckAndCall(webglContext, "linkProgram", program)
		api.JsCheckAndCall(webglContext, "useProgram", program)

		attrVertex, _ := api.JsCheckAndCall(webglContext, "getAttribLocation", program, "attrVertex")
		api.JsCheckAndCall(webglContext, "enableVertexAttribArray", attrVertex)
		api.JsCheckAndCall(webglContext, "vertexAttribPointer", attrVertex, 2, webglContext.Get("FLOAT"), false, 0, 0)

		uniformOffset, _ := api.JsCheckAndCall(webglContext, "getUniformLocation", program, "uniformOffset")
		api.JsCheckAndCall(webglContext, "uniform2f", uniformOffset, 1, 1)

		api.JsCheckAndCall(webglContext, "drawArrays", webglContext.Get("TRIANGLE_STRIP"), 0, len(vertices)/2)

		url, _ := api.JsCheckAndCall(webglContext.Get("canvas"), "toDataURL")

		result = url.String()
	}

	return result
}
