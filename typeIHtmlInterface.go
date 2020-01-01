package iotmaker_platform_IDraw

import "github.com/helmutkemper/iotmaker.platform.webbrowser/Html"

type IHtml interface {
	NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) Html.Image
	Append(document, element interface{})
	Remove(document, element interface{})
	GetDocumentWidth(document interface{}) int
	GetDocumentHeight(document interface{}) int
}
