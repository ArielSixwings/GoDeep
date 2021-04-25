package main

import (
	geneticextractor "./folder"
)

func main() {
	var geneticCreator geneticextractor.GeneticExtractor
	//geneticCreator.GenerateStringFather()
	//geneticCreator.PrintFather()
	//geneticCreator.GenerateDataSet(geneticCreator.GetFathers(), "./folder/teste.txt", 64000)
	geneticCreator.GenerateStringChild("./folder/teste.txt")
	geneticCreator.PrintChild()
	geneticCreator.GenerateDataSet(geneticCreator.GetChilds(), "./folder/teste2.txt", 64000)
}
