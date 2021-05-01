package process

import(
	"../basicdata"
)

type processStrategy interface {
	Process() error
	Allocate() error
	SaveProcessedData() error 
	PresentProcessing(Menssage string, index int, time int) error
}

type DataProcessing struct{
	Strategy processStrategy

	Readinfo 	cartesian.ReadInformation
}