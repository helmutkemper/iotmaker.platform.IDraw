package iotmaker_platform_IDraw

import "image/color"

type IDraw interface {
	BeginPath()
	MoveTo(x, y int)
	ArcTo(x, y, radius, startAngle, endAngle int)
	LineTo(x, y int)
	ClosePath(x, y int)
	Stroke()
	LineWidth(value int)
	ShadowBlur(value int)
	ShadowColor(value color.RGBA)
	ShadowOffsetX(value int)
	ShadowOffsetY(value int)
	AddColorStop(gradient interface{}, stop float64, color color.RGBA)
	CreateLinearGradient(x0, y0, x1, y1 int) interface{}
	CreateRadialGradient(x0, y0, r0, x1, y1, r1 int)
	FillStyle(value interface{})
	StrokeStyle(value interface{})
}
