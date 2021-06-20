package iotmaker_platform_IDraw

import (
	commonTypes "github.com/helmutkemper/iotmaker.santa_isabel_theater.commonTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
)

type IHtml interface {
	NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) Html.Image
	Append(document, element interface{})
	Remove(document, element interface{})
	GetDocumentWidth(document interface{}) commonTypes.Number
	GetDocumentHeight(document interface{}) commonTypes.Number
}
