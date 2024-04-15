package module

import (
	"fingerprint/api"
	"syscall/js"
)

func drawCircle(ctx js.Value, x, y, radius int) {
	Pi := 3.14159265358979323846264338327950288419716939937510582097494459
	api.JsCheckAndCall(ctx, "beginPath")
	api.JsCheckAndCall(ctx, "arc", x, y, radius, 0, Pi*2, true)
	api.JsCheckAndCall(ctx, "closePath")
	api.JsCheckAndCall(ctx, "fill")
}

func GetCanvasFingerprint() string {

	doc := api.Global.Get("document")

	if api.JsCheck(doc, api.TypeObject) {

		canvas, _ := api.JsCheckAndCall(doc, "createElement", "canvas")
		canvas.Set("width", 2000)
		canvas.Set("height", 200)

		style := canvas.Get("style")
		if api.JsCheck(style, api.TypeObject) {
			style.Set("display", "inline")
		}

		ctx, _ := api.JsCheckAndCall(canvas, "getContext", "2d")
		api.JsCheckAndCall(ctx, "rect", 0, 0, 10, 10)
		api.JsCheckAndCall(ctx, "rect", 2, 2, 6, 6)

		ctx.Set("textBaseline", "alphabetic")
		ctx.Set("fillStyle", "#f60")
		api.JsCheckAndCall(ctx, "fillRect", 125, 1, 62, 20)
		ctx.Set("fillStyle", "#069")
		ctx.Set("font", "11pt Arial")
		api.JsCheckAndCall(ctx, "fillText", "Cwm fjordbank glyphs vext quiz, 😂😂", 2, 15)

		ctx.Set("fillStyle", "rgba(102, 204, 0, 0.2)")
		ctx.Set("font", "18pt Arial")
		api.JsCheckAndCall(ctx, "fillText", "Cwm fjordbank glyphs vext quiz, 😂😂", 4, 45)

		ctx.Set("globalCompositeOperation", "multiply")
		ctx.Set("fillStyle", "rgb(255,0,255)")
		drawCircle(ctx, 50, 50, 50)
		ctx.Set("fillStyle", "rgb(0,255,255)")
		drawCircle(ctx, 100, 50, 50)
		ctx.Set("fillStyle", "rgb(255,255,0)")
		drawCircle(ctx, 75, 100, 50)
		ctx.Set("fillStyle", "rgb(255,0,255)")

		drawCircle(ctx, 75, 75, 75)
		drawCircle(ctx, 75, 75, 25)
		api.JsCheckAndCall(ctx, "fill", "evenodd")

		result, _ := api.JsCheckAndCall(canvas, "toDataURL")

		return result.String()
	}

	return "error"
}
