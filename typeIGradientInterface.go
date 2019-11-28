package iotmaker_platform_IDraw

import "image/color"

type IGradient interface {
	AddColorStop(gradient interface{}, stop float64, color color.RGBA)
	CreateLinearGradient(x0, y0, x1, y1 int) interface{}
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 int)
	FillStyle(value interface{})
	StrokeStyle(value interface{})
}
