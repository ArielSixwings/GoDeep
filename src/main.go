package main

//import "./genetic/GeneticInformation"
import "./textextractor"

func main() {
	var reader textextractor.FolderExtractor
	//var process geneticinformation.GeneticInformation
	reader.ScanFolder("./genetic/Extractor/tapes")
	//process.GetResult(reader.GetTexts(0), reader.GetTexts(1), 0)

	reader.PrintFile()
}
