package iotmaker_platform_IDraw

type IHtml interface {
	NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) interface{}
	Append(document, element interface{})
}
