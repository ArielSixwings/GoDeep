package main

import (
	"fmt"
	"bufio"
//	"errors"
	"os"
//	"path/filepath"
	"strings"
	"strconv"
	"../generalizecartesian"
	"../basicdata"
)

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func scanText(caminhoDoArquivo string) ([]string, error) {
	var linhas []string
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return linhas, scanner.Err()
}

func printText(conteudo []string, err error) {
	if err != nil {
		//log.Fatalf("Erro:", err)
		fmt.Println("Erro:", err)
	}

	for indice, linha := range conteudo {
		fmt.Println(indice, linha)
	}
}

func splitTextTrain(conteudo []string) ([]string, []string, []string){
	var split = make([]string, len(conteudo))
	var pSurvived, pClass, pSex, pAge []string
	var folderText string
	var i int = 0

	for i=0; i<len(conteudo); i++ {
		folderText = folderText + conteudo[i]
	}

	for i=0; i<len(conteudo); i++ {
		split = append(strings.Split(folderText, ","))
	}

	// Spliting information
	for i=12; i<len(split)/2; i+=12 {
		pSurvived = append(pSurvived, split[i])
	}

	for i=13; i<len(split)/2; i+=12 {
		pClass = append(pClass, split[i])
	}

	for i=16; i<len(split)/2; i+=12 {
		pSex = append(pSex, split[i])
	}

	for i=17; i<len(split)/2; i+=12 {
		pAge = append(pAge, split[i])
	}

	return pAge, pClass, pSex
}

func splitTextKnow(conteudo []string) ([]string, []string, []string){
	var split = make([]string, len(conteudo))
	var pSurvived, pClass, pSex, pAge []string
	var folderText string
	var i int = 0

	for i=0; i<len(conteudo); i++ {
		folderText = folderText + conteudo[i]
	}

	for i=0; i<len(conteudo); i++ {
		split = append(strings.Split(folderText, ","))
	}

	// Spliting information
	for i=12+(len(split)/2); i<len(split); i+=12 {
		pSurvived = append(pSurvived, split[i])
	}

	for i=13+(len(split)/2); i<len(split); i+=12 {
		pClass = append(pClass, split[i])
	}

	for i=16+(len(split)/2); i<len(split); i+=12 {
		pSex = append(pSex, split[i])
	}

	for i=17+(len(split)/2); i<len(split); i+=12 {
		pAge = append(pAge, split[i])
	}


	return pAge, pClass, pSex
}

func convertData(data_1 []string, data_2 []string, data_3 []string) ([]float64, []float64, []float64) {
	var pData1Float64, pData2Float64,pData3Float64 float64
	var pData1SliceFloat64, pData2SliceFloat64, pData3SliceFloat64 []float64
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

	//fmt.Println(pData1SliceFloat64)
	//fmt.Println(pData2SliceFloat64)
	//fmt.Println(pData3SliceFloat64)

	return pData1SliceFloat64, pData2SliceFloat64, pData3SliceFloat64

}


//###### APAGAR #######
func main() {
	var dataKnow, dataTrain []string
	var tAge, tClass, tSex, kAge, kClass, kSex []string
	var tData1, tData2, tData3, kData1, kData2, kData3 []float64
	var dataset generalizecartesian.Labelfeatures

	// Scan folders
	dataTrain, err := scanText("tempTrain.csv")
	printText(dataTrain, err)
	dataKnow, err = scanText("tempTest.csv")
	printText(dataKnow, err)

	// Convert data from data set train
	tAge, tClass, tSex = splitTextTrain(dataTrain)
	tData1, tData2, tData3 = convertData(tAge, tClass, tSex)

	kAge, kClass, kSex = splitTextKnow(dataKnow)
	kData1, kData2,  kData3 = convertData(kAge, kClass, kSex)

	// Convert data sets
	//dataSetTrain, dataSetKnow = convertDataSet(dataTrain, dataKnow)
	

	fmt.Println(tData1)
	fmt.Println(tData2)
	fmt.Println(tData3)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(kAge)
	fmt.Println(kClass)
	fmt.Println(kSex)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(kData1)
	fmt.Println(kData2)
	fmt.Println(kData3)

	//fmt.Println("Generalizing know data set")
	//generalizecartesian.Generalize_for_nonparametric(&dataset, kData1, kData2, kData3, knowls, generalizecartesian.Knowflag,3*knowsize)
	
	//fmt.Println("Generalizing train data set")
	//generalizecartesian.Generalize_for_nonparametric(&dataset, tData1, tData2, tData3,trainls, generalizecartesian.Trainflag,3*trainsize)


}