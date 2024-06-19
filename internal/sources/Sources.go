package sources

type SourcesStruct struct {
}

type SourcesInterface interface {
	Read()
	ReadStream()
}

func (sources *SourcesStruct) Read()       {}
func (sources *SourcesStruct) ReadStream() {}
