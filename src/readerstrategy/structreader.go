package readerstrategy

import (
	"../basicdata"
)

type readStrategy interface {
	ReadData(rd *DataReader) error
	Allocate(rd *DataReader) error
}
//func (ie *ImageExtractor)  ReadData(path string, show bool, colorfull bool, i int){

type readerParameters struct{
	path string
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