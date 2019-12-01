package iotmaker_platform_IDraw

import "image/color"

type ICanvasShadow interface {
	ShadowBlur(value int)
	ShadowColor(value color.RGBA)
	ShadowOffsetX(value int)
	ShadowOffsetY(value int)
}
