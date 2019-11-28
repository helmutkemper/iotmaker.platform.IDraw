package iotmaker_platform_IDraw

import "image/color"

type IShadow interface {
	ShadowBlur(value int)
	ShadowColor(value color.RGBA)
	ShadowOffsetX(value int)
	ShadowOffsetY(value int)
}
