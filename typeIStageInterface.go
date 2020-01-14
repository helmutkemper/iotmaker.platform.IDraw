package iotmaker_platform_IDraw

type IStage interface {
	Init()
	SetSleepFrame(value int)
	GetSleepFrame() int
	Set(value int)
	Get() int
	AddCursorDrawFunc(runnerFunc func()) string
	RemoveCursorDrawFunc(id string)
	AddToLowLatency(runnerFunc func()) string
	DeleteFromLowLatency(UId string)
	AddToSystem(runnerFunc func()) string
	DeleteFromSystem(UId string)
	AddToCalculate(runnerFunc func()) string
	DeleteFromCalculate(UId string)
	AddToDraw(runnerFunc func()) string
	DeleteFromDraw(UId string)
	SetZIndex(UId string, index int) int
	ToFront(UId string) int
	ToBack(UId string) int
	GetZIndex(UId string) int
}
