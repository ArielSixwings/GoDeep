package main

import (
	"../src/extractstrategy"
	"../src/learnstrategy"
	"../src/process"
	"../src/learnstrategy/nonparametric"
	//"fmt"
)

func main() {
	var datasetextractor extract.TextExtractor
	var datatransformer process.StatisticProcessing
	var datalearner learnstrategy.DataLearner

	origins := []string{"../data/StatisticData/titanic"}

	datasetextractor.SetOrigins(origins,&datasetextractor)

	datasetextractor.Read(false,false,false)
	//datasetextractor.PrintFile()
	
	datatransformer.Texts = datasetextractor.Texts
	datatransformer.ConvertData()
	datatransformer.PrintFeatures()

	datalearner.Build(&datatransformer.Information,datasetextractor.Readinfo,445)

	knn := &nonparametric.Knn{}
	datalearner.SetLearnStrategy(knn)
	datalearner.ProcessLearn()
	datalearner.Printresults()	
}