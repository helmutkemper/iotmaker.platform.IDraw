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
}
