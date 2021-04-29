package extract

import (
	"../basicdata"
)

type readStrategy interface {
	ReadData(path string,dataindex int) error
	Allocate() error
	PresentData(Menssage string, index int, time int)
	SaveData(i int,Name string) 
}
//func (ie *ImageExtractor)  ReadData(path string, show bool, colorfull bool, i int){

type readerParameters struct{
	Format bool
	Show bool
	Print bool
}

type DataReader struct {
	readerParameters

	Strategy readStrategy
	
	split 		[][]string
	readOrigins []string
	Readinfo 	cartesian.ReadInformation
}