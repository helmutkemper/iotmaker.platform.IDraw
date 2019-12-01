package iotmaker_platform_IDraw

import "image/color"

type ICanvasGradient interface {
	AddColorStop(gradient interface{}, stop float64, color color.RGBA)
	CreateLinearGradient(x0, y0, x1, y1 interface{}) interface{}
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 interface{}) interface{}
	FillStyle(value interface{})
	StrokeStyle(value interface{})
}
