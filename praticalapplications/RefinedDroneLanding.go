package main

import (
	"../src/readerstrategy"
	"../src/learnstrategy"
	
	"gocv.io/x/gocv"
	"../src/imagehandler/imageprocessing"
	"../src/imagehandler/computervision"

	
	"../src/learnstrategy/nonparametric"
	"fmt"
	

)

func main() {
	var datasetextractor extract.ImageExtractor
	var datatransformer imageprocessing.ImageProcessing
	var datavision computervision.ComputerVison
	var datalearner learnstrategy.DataLearner

	var normtype gocv.NormType = gocv.NormMinMax
	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	datasetextractor.SetOrigins(origins,&datasetextractor)

	fmt.Println("About to call Read")
	datasetextractor.Read(false,false,true)
	
	datatransformer.GetImages(&datasetextractor)
	datatransformer.GroupGLCM(true, true)
	datatransformer.GroupNormalizedGLCM(0.0, 255.0, normtype,true ,true)	
	
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