package main

import (
	"testing"
	"../src/extractstrategy"
	"../src/learnstrategy"

	"gocv.io/x/gocv"
	"../src/processStrategy"
	"../src/DataAnalysis"

	"../src/learnstrategy/nonparametric"
)

func TestGetData(t *testing.T) {
	var datasetextractor extract.ImageExtractor
	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	if !datasetextractor.SetOrigins(origins,&datasetextractor) {
		t.Error("Value unexpected")
	}

	
}