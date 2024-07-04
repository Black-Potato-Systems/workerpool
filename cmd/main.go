package main

import (
	"flag"

	"github.com/blackpotato/workerpool/internal/schemavalidatorservice"

	"github.com/golang/glog"
)

func main() {

	// This main function would be used as an
	// examples file to shoewcase the capabilities of the
	// library.

	// file should initialize interface and its services


	// Always On application...... infinite for loop 
	// use of select and context 

	// Define the pipeline ..... in a confg file ....

	flag.Set("logtostderr", "1")
	flag.Parse()

	fries := schemavalidatorservice.InitFries().
		Source("local").
		ServiceMode("processor").
		NumberOfWorkers(2).
		Sink("kafka").
		Build()

	glog.Infoln(fries)

	// reeads from kafka topic, this should return []byte array
	friesDataFrame := fries.Read("reading local file....")
	// convert csv data into tabular form
	// this process should take a function as a parameter and then... run that function in worker pool
	// Every Process mustleave footprint and should be a part of DAG
	friesDataFrame.Process().Process().Process().LongRunningProces(2)

	// internally Read rea the file ..parallely and further combine the chunks
	// add all the chunks to the friesFrame.

	// this friesFrame gives the ability to read the data in tabular format in rows and columns

	// or It can return friesDataStructure

	// Init Source and Sinks
	//  basically there should be a method  directly coming from fies variable which allows to
	// read the data either from the batch source \
	// or from the streaming source.

	// after reading the data, there should be abstracted process function,
	// whatever written in the process function should be converted into
	// go routine the number of long running workers defined
	// Also the input data should be divied based on the number of workers defined
	// and sent over a channel to process further.....
	// ther could be multiple process functions written

	// There should a geberic method here which abstracts the underlying source here

	// Again the sink should be abstracted from the user/

	// It should just expose couple of methods like writeToConsole

	// Top Level functions ......

	//

	// Draw a design Diagram first

	// What type of worker pool to select?

	// long running worker pool?

	// worker pools which closes after the work has been done

	// chain of worker pools ........

	// Source of data ......

	// type of data

	// protobuf data

	// could be used as a library....

}
