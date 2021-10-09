package main

import (
	"../src/ExtractStrategy"
	"../src/LearnStrategy"
	
	"gocv.io/x/gocv"
	"../src/ProcessStrategy"
	"../src/DataAnalysis"

	
	"../src/LearnStrategy/nonparametric"
)

func main() {
	var (
		datasetextractor extract.ImageExtractor
		datatransformer process.ImageProcessing
		datavision computervision.ComputerVison
		datalearner learnstrategy.DataLearner

		normtype gocv.NormType = gocv.NormMinMax

		glcm process.GLCM
		normalize process.Normalize
	)

	origins := []string{"../data/ImagesData/danger", 
		"../data/ImagesData/asphalt", 
		"../data/ImagesData/grass"}

	datasetextractor.SetOrigins(origins,&datasetextractor)

	datasetextractor.Read(true,true,true)
	
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