package main

import (
	"../src/extractstrategy"
	"../src/learnstrategy"
	
	"gocv.io/x/gocv"
	"../src/processStrategy"
	"../src/DataAnalysis"

	
	"../src/learnstrategy/nonparametric"
	"fmt"
	

)

func main() {
	var datasetextractor extract.ImageExtractor
	var datatransformer process.ImageProcessing
	var datavision computervision.ComputerVison
	var datalearner learnstrategy.DataLearner

	var normtype gocv.NormType = gocv.NormMinMax
	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	var glcm process.GLCM
	var normalize process.Normalize

	datasetextractor.SetOrigins(origins,&datasetextractor)

	fmt.Println("About to call Read")
	datasetextractor.Read(false,false,true)
	
	datatransformer.GetImages(&datasetextractor)
	
	glcm.SetParameters(1,0)
	datatransformer.SetProcessStrategy(glcm)
	datatransformer.ProcessGroup(true)
	
	normalize.SetParameters(0.0, 255.0, normtype)
	datatransformer.SetProcessStrategy(normalize)
	datatransformer.ProcessGroup(true)
	
	datavision.GetBaseImages(&datatransformer)
	datavision.GroupFeature(true,computervision.EnergyFeature,computervision.CorrelationFeature,computervision.ContrastFeature)
	datavision.PrintFeatures()

	datalearner.Build(&datavision.Information,datasetextractor.Readinfo,75)
	datalearner.Printfeatures()

	knn := &nonparametric.Knn{}
	datalearner.SetLearnStrategy(knn)
	datalearner.ProcessLearn()
	datalearner.Printresults()
}