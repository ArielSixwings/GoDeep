package main

import (
	"testing"

	"../src/extractstrategy"
	//"../src/learnstrategy"

	//"gocv.io/x/gocv"
	//"../src/processStrategy"
	
	//"../src/DataAnalysis"
	//"../src/learnstrategy/nonparametric"
)

var datasetextractor extract.ImageExtractor

func TestSetOrigins(t *testing.T) {
	
	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	_, err := datasetextractor.SetOrigins(origins,&datasetextractor)

	if err != nil {
		t.Error("Unexpected value")
	}
}

func TestRead(t *testing.T) {

	err := datasetextractor.Read(false,false,true)

	if err != nil {
		t.Error("Unexpected value")
	}
}