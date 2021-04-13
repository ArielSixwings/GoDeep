package main

import (
	"../src/imagehandler/imageextractor"
	"../src/imagehandler/imageprocessing"
	"../src/imagehandler/computervision"
	"../src/learnstrategy/nonparametric"
	"../src/learnstrategy"
	"../src/basicdata"
	"gocv.io/x/gocv"
)

func main() {
	var normtype gocv.NormType = gocv.NormMinMax
	var datasetextractor imageextractor.ImageExtractor
	var datatransformer imageprocessing.ImageProcessing
	var datavision computervision.ComputerVison
	var datalearner learnstrategy.DataLearner

	labelsize := make([]cartesian.Sizelabel,3)
	
	labelsize[0].Label = "danger"
	labelsize[0].Size_l = 25

	labelsize[1].Label = "asphalt"
	labelsize[1].Size_l = 25

	labelsize[2].Label = "grass"
	labelsize[2].Size_l = 25

	origins := []string{"../src/imagehandler/Images/danger", 
		"../src/imagehandler/Images/asphalt", 
		"../src/imagehandler/Images/grass"}

	datasetextractor.SetOrigins(origins)
	datasetextractor.Read()
	
	datatransformer.GetImages(&datasetextractor)
	datatransformer.GroupGLCM(true, true)
	datatransformer.GroupNormalizedGLCM(0.0, 255.0, normtype,true ,true)	
	
	datavision.GetBaseImages(&datatransformer)
	datavision.GroupFeature(true,computervision.EnergyFeature,computervision.CorrelationFeature,computervision.ContrastFeature)
	//datavision.DefineLabels() 	//+++++++++++++++++++++++++ create that function
	datavision.PrintFeatures()

	datalearner.Build(&datavision,labelsize,75)
	datalearner.Printfeatures()

	knn := &nonparametric.Knn{}
	datalearner.SetLearnStrategy(knn)
	datalearner.ProcessLearn()
	datalearner.Printresults()
}