package iotmaker_platform_IDraw

type IFilterGradientInterface interface {
	PrepareFilter(platform ICanvasGradient)
	Fill(gradient interface{})
	Stroke(gradient interface{})
}
