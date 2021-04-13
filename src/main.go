package main

import (
	"./textextractor"
)

func main() {
	var reader textextractor.TextExtractor
	var reader2 textextractor.TextExtractor
	reader.ScanText("./genetic/tapes/fathers.txt")
	reader2.ScanText("./genetic/tapes/childs.txt")
	reader.PrintFile()
	reader2.PrintFile()

}
