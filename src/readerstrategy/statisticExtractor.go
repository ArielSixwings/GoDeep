package extract

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func scanText(filePath string) ([]string, error) {
	var lines []string
	
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func printText(fileContents []string, err error) {
	if err != nil {
		//log.Fatalf("Erro:", err)
		fmt.Println("Erro:", err)
	}

	for index, line := range fileContents {
		fmt.Println(index, line)
	}
}

func splitText(fileContents []string) ([]string, []string, []string, []string){
	var split = make([]string, len(fileContents))
	var pSurvived, pClass, pSex, pAge []string
	var folderText string
	var i int = 0

	for i=0; i<len(fileContents); i++ {
		folderText = folderText + fileContents[i]
	}

	for i=0; i<len(fileContents); i++ {
		split = append(strings.Split(folderText, ","))
	}

	// Spliting information
	for i=12; i<len(split); i+=12 {
		pSurvived = append(pSurvived, split[i])
	}

	for i=13; i<len(split); i+=12 {
		pClass = append(pClass, split[i])
	}

	for i=16; i<len(split); i+=12 {
		pSex = append(pSex, split[i])
	}

	for i=17; i<len(split); i+=12 {
		pAge = append(pAge, split[i])
	}

	return pAge, pClass, pSex, pSurvived
}

func convertData(data_1 []string, data_2 []string, data_3 []string, data_4 []string) ([]float64, []float64, []float64, []float64) {
	var pData1Float64, pData2Float64,pData3Float64, pData4Float64 float64
	var pData1SliceFloat64, pData2SliceFloat64, pData3SliceFloat64, pData4SliceFloat64 []float64
	var i int

	for i=0; i<len(data_1); i++ {
		pData1Float64, _ = strconv.ParseFloat(data_1[i], 64)
		pData1SliceFloat64 = append(pData1SliceFloat64, pData1Float64)
	}

	for i=0;i<len(data_2); i++ {
		pData2Float64, _ = strconv.ParseFloat(data_2[i], 64)
		switch pData2Float64 {
			case 1:
				pData2Float64 = 0
			case 2: 
				pData2Float64 = 1
			case 3:
				pData2Float64 = 2
		}
		pData2SliceFloat64 = append(pData2SliceFloat64, pData2Float64)
	}

	for i=0;i<len(data_3); i++ {
		pData3Float64, _ = strconv.ParseFloat(data_3[i], 64)

		if data_3[i] == "male" {
			pData3Float64 = 0
		} else {
			pData3Float64 = 1
		}

		pData3SliceFloat64 = append(pData3SliceFloat64, pData3Float64)
	}

	for i=0; i<len(data_4); i++ {
		pData4Float64, _ = strconv.ParseFloat(data_4[i], 64)
		pData4SliceFloat64 = append(pData4SliceFloat64, pData4Float64)
	}

	return pData1SliceFloat64, pData2SliceFloat64, pData3SliceFloat64, pData4SliceFloat64

}

func sortData(data_4 []string, tamData int) ([]string) {
	var sortedData = make([]string, tamData)
	lowerVectorIndex := 0
	upperVectorIndex := 0

	for i := 0; i < tamData; i++ {
		if data_4[i] == "0" {
			sortedData[lowerVectorIndex] = data_4[i]
			lowerVectorIndex++
		} else {
			sortedData[tamData - 1 - upperVectorIndex] = data_4[i]
			upperVectorIndex++
		}
	}

	return sortedData
}