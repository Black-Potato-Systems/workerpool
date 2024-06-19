package schemavalidatorservice

import (
	"io"

	friesframeapi "github.com/blackpotato/workerpool/internal/friesFrameAPI"
	"github.com/blackpotato/workerpool/internal/sources"

	"os"

	"github.com/golang/glog"
)

type SchemaValidatorServiceStruct struct {
	ServiceModeField string
	ReadSource       string
	ReadStruct       sources.SourcesStruct
	WorkerCount      int
}

type SchemValidatorServiceInterface interface {
	SchemaVaidatorServiceSession()
	ServiceMode(serviceMode string) SchemValidatorServiceInterface
	Source(sourceType string) SchemValidatorServiceInterface
	Sink(sinkType string) SchemValidatorServiceInterface
	SechemaValidatorServiceType()
	ErrorWroker()
	NumberOfWorkers(workerCount int) SchemValidatorServiceInterface
	Build() SchemValidatorServiceInterface
	NumberOfchains()
	Nodes()
	Metrics()
	ApplyRegex()
	Read(readFile string) *friesframeapi.FriesFrame
	ReadStream(stream string) *friesframeapi.FriesFrame
}

// nodes? // schema node // parser node // writer node === this service can be run in various modes on a
// different kubertestes containers.....

// to attain maximum throughput from a node

func (s *SchemaValidatorServiceStruct) SchemaVaidatorServiceSession() {

}

func (s *SchemaValidatorServiceStruct) Read(readFile string) *friesframeapi.FriesFrame {

	readSource := s.ReadSource

	if readSource == "local" {
		// init kafka workflow
		// read csv file into chunks using multiple go routines
		glog.Info(readFile)

		inputFile, err := os.Open("/home/blackpotato/Documents/xxxx.txt")

		if err != nil {
			glog.Infoln("panic reading file.......")
			panic(err)
		}

		defer func() {
			err := inputFile.Close()

			if err != nil {
				panic(err)
			}
		}()

		glog.Infoln(inputFile.Name())

		// create output file

		outputFile, err := os.Create("/home/blackpotato/Documents/output.txt")
		outputFile.Chmod(0777)

		if err != nil {
			panic(err)
		}

		defer func() {

			err := outputFile.Close()
			if err != nil {
				panic(err)
			}

		}()

		tempBuffer := make([]byte, 1024)

		// read a chunk of file

		glog.Infoln("going to read file..............")

		for {

			n, err := inputFile.Read(tempBuffer)

			if err != nil && err != io.EOF {
				panic(err)
			}
			glog.Infoln("keep reading.........")
			if n == 0 {
				glog.Infoln("breaking out.......")
				break
			}

			glog.Infoln("N value....", n)
			glog.Info("keep writing.......")
			glog.Infoln("string buffer", string(tempBuffer))
			if _, err := outputFile.Write(tempBuffer[:n]); err != nil {
				glog.Infoln("error in writing............")
				panic(err)
			}

		}
	}

	// read data parllely and convert it into FriesFrame Structure

	return &friesframeapi.FriesFrame{}
}

func (s *SchemaValidatorServiceStruct) Process() *friesframeapi.FriesFrame {

	readSource := s.ReadSource

	if readSource == "local" {
		// init kafka workflow
		glog.Info("readFile")
	}

	// read data parllely and convert it into FriesFrame Structure

	return &friesframeapi.FriesFrame{}
}

func (s *SchemaValidatorServiceStruct) ReadStream(stream string) *friesframeapi.FriesFrame {

	readSource := s.ReadSource

	if readSource == "local" {
		// init kafka workflow
		glog.Info(stream)
	}

	return &friesframeapi.FriesFrame{}
}

func InitFries() SchemValidatorServiceInterface {

	return &SchemaValidatorServiceStruct{}

}

func (s *SchemaValidatorServiceStruct) ServiceMode(serviceMode string) SchemValidatorServiceInterface {

	if serviceMode == "processingNode" {
		// run the framework in the processing mode
		// this should need soy=urce and sinks

	} else if serviceMode == "schemaNode" {
		// run the framework in the schemaMode
	} else if serviceMode == "persistanceMode" {
		// run the framework in the servicemOde
	} else if serviceMode == "webRTC" {
		// run the framework in the webRTC mode (directly connecting to interface for real time data view)
	}
	s.ServiceModeField = serviceMode
	return s

}

func (s *SchemaValidatorServiceStruct) Source(sourceType string) SchemValidatorServiceInterface {

	if sourceType == "local" {
		// enable kafka consumer flow
		s.ReadSource = sourceType
	} else if sourceType == "rest" {
		// enable rest api workflow
	} else if sourceType == "webRtC" {
		// enable web rtc flow
	} else if sourceType == "financial transaction" {
		// enable financial data
	} else if sourceType == "CDC" {
		// enable cdc workflow
	}

	return s
}

func (s *SchemaValidatorServiceStruct) Sink(sinkType string) SchemValidatorServiceInterface {

	if sinkType == "kafka" {
		// enable kafka consumer flow
	} else if sinkType == "local" {
		// enable rest api workflow
	} else if sinkType == "cmd" {
		// enable web rtc flow
	} else if sinkType == "s3" {
		// enable financial data
	} else if sinkType == "snowflake" {
		// enable cdc workflow
	}

	return s
}

func (s *SchemaValidatorServiceStruct) SechemaValidatorServiceType() {

}

func (s *SchemaValidatorServiceStruct) ErrorWroker() {

}

func (s *SchemaValidatorServiceStruct) NumberOfWorkers(workerCount int) SchemValidatorServiceInterface {
	s.WorkerCount = workerCount
	return s

}

// Check all the input from different methods and build the final struct
func (s *SchemaValidatorServiceStruct) Build() SchemValidatorServiceInterface {
	return s
}
func (s *SchemaValidatorServiceStruct) NumberOfchains() {}
func (s *SchemaValidatorServiceStruct) Nodes()          {}
func (s *SchemaValidatorServiceStruct) Metircs()        {}
func (s *SchemaValidatorServiceStruct) Logging()        {}
func (s *SchemaValidatorServiceStruct) ApplyRegex()     {}

func (s *SchemaValidatorServiceStruct) Metrics() {}
