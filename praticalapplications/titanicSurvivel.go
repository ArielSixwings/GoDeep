package main

import (
	"fmt"
	
	"../src/statistics"
	"../src/generalizecartesian"
	"../src/nonparametric"
)

func main() {
	path := "tempTrain.csv"
	filePath := filepath.Join(path)

	lines, err := scanText(filePath)
	printText(lines, err)
}
