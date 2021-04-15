package main

import (
	"./textextractor"
)

func main() {
	var leitor [][]string
	var reader textextractor.TextExtractor
	leitor = reader.ScanFolder("./genetic/tapes")
	reader.PrintFile(leitor)
}
