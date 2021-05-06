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

func TestSetOrigins(t *testing.T) {
	var datasetextractor extract.ImageExtractor
	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	_, err := datasetextractor.SetOrigins(origins,&datasetextractor)

	if err != nil {
		t.Error("Unexpected value")
	}
}