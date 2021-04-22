package main

import (
	"./src/readerstrategy"
	"./src/learnstrategy"
	
	"gocv.io/x/gocv"
	"./src/imagehandler/imageprocessing"
	"./src/imagehandler/computervision"

	
	"./src/learnstrategy/nonparametric"
	

)

func main() {
	var datasetextractor extractstrategy.DataReader
	var datatransformer imageprocessing.ImageProcessing
	var datavision computervision.ComputerVison
	var datalearner learnstrategy.DataLearner

	var normtype gocv.NormType = gocv.NormMinMax
	origins := []string{"./src/imagehandler/Images/danger", 
		"./src/imagehandler/Images/asphalt", 
		"./src/imagehandler/Images/grass"}

	ImageApplication := &extractstrategy.ImageExtractor{}
	datasetextractor.SetReadStrategy(ImageApplication)
	datasetextractor.SetOrigins(origins)
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
	datasetextractor.Read()
	
	datatransformer.GetImages(ImageApplication)
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