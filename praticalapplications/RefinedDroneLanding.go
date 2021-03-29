package main

import (
	"../src/imageprocessing"
	//"../src/learnstrategy/nonparametric"
	//"../src/learnstrategy"
	//"../src/basicdata"
	"gocv.io/x/gocv"
	//"fmt"
	//"math"
)

func main() {
	var normtype gocv.NormType = gocv.NormMinMax
	var datasetextractor imageprocessing.ImageExtractor
	var datatransformer imageprocessing.ImageProcessing
	var datavision imageprocessing.ComputerVison
	var datalearner learnstrategy.DataLearner

	labelsize := make([]cartesian.Sizelabel,3)
	
	labelsize[0].Label = "danger"
	labelsize[0].Size_l = 50

	labelsize[1].Label = "asphalt"
	labelsize[1].Size_l = 50

	labelsize[2].Label = "grass"
	labelsize[2].Size_l = 50

	datasetextractor.ReadFolder("../src/imageprocessing/Images/danger",true,true,false)
	datasetextractor.ReadFolder("../src/imageprocessing/Images/asphalt",true,true,false,50)
	datasetextractor.ReadFolder("../src/imageprocessing/Images/grass",true,true,false,100)
	
	datatransformer.GetImages(&datasetextractor)
	datatransformer.GroupGLCM(true, true)
	datatransformer.GroupNormalizedGLCM(0.0, 255.0, normtype,true ,true)	
	
	datavision.GetBaseImages(&datatransformer)
	datavision.GroupFeature(true,imageprocessing.EnergyFeature,imageprocessing.CorrelationFeature,imageprocessing.ContrastFeature)
	datavision.DefineLabels()
	datavision.PrintFeatures()
	
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
	datalearner.Build(&datavision,labelsize,75)

	fmt.Println("Calling KNN")
	knn := &nonparametric.Knn{}
	dataset.SetLearnStrategy(knn)
	dataset.ProcessLearn()
	dataset.Printresults()
}