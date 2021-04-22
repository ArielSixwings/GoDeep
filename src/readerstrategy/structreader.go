package extractstrategy

import (
	"../basicdata"
)

type readStrategy interface {
	ReadData(rd DataReader,i int) error
	Allocate(rd DataReader) error
}
//func (ie *ImageExtractor)  ReadData(path string, show bool, colorfull bool, i int){

type readerParameters struct{
	index int
	Format bool
	Show bool
	print bool
}

type DataReader struct {
	readerParameters

	Strategy readStrategy
	
	split 		[][]string
	readOrigins []string
	Readinfo 	cartesian.ReadInformation
}