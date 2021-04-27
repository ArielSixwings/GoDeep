package main

import (
	"../src/extractstrategy"
	"fmt"
)

func main() {
	var datasetextractor extract.TextExtractor

	origins := []string{"../data/GeneticData"}

	datasetextractor.SetOrigins(origins,&datasetextractor)

	fmt.Println("About to call Read")
	datasetextractor.Read(false,false,true)
	datasetextractor.PrintFile()	
}