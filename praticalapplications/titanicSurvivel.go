package main

import (
	"../src/ExtractStrategy"
	"../src/LearnStrategy"
	"../src/ProcessStrategy"
	"../src/LearnStrategy/nonparametric"
)

func main() {
	var datasetextractor extract.TextExtractor
	var datatransformer process.StatisticProcessing
	var datalearner learnstrategy.DataLearner

	origins := []string{"../data/StatisticData/titanic"}

	datasetextractor.SetOrigins(origins,&datasetextractor)

	datasetextractor.Read(false,false,false)
	datasetextractor.PrintFile()
	
	datatransformer.Texts = datasetextractor.Texts
	datatransformer.ConvertData()
	datatransformer.PrintFeatures()

	datasetextractor.Readinfo.SizeData = 890
	datalearner.Build(&datatransformer.Information,datasetextractor.Readinfo,445)

	knn := &nonparametric.Knn{}
	datalearner.SetLearnStrategy(knn)
	datalearner.ProcessLearn()
	datalearner.Printresults()	
}