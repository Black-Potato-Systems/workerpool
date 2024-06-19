package friesframeapi

type FriesFrame struct {
}

// FriesFrameAPI for parallel process the data.......
type FriesFrameInterface interface {
	Show()
	ShowAll()
	Process() *FriesFrame
	LongRunningProcess(workerCount int)
}

func (f *FriesFrame) Show() {
}

func (f *FriesFrame) ShowAll() {

}

func (f *FriesFrame) Process() *FriesFrame {
	return f
}

func (f *FriesFrame) LongRunningProces(workerCount int) {

}
