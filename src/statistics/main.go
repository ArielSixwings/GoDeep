package main

import (
	"bufio"
	"fmt"

	//"log"
	"os"
)

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func scanText(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
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

//###### APAGAR #######
func main() {
	var conteudo []string
	conteudo, err := scanText("tempTrain.csv")
	printText(conteudo, err)

	Allocate(allflag Groupflag, allsize int, secondsize ...float64)
}













func convertDataSet(dataTrain []string, dataKnow []string) ([]float64, []float64) {
	var dataTrainFloat64, dataKnowFloat64 float64
	var dataTrainSliceFloat64, dataKnowSliceFloat64 []float64
	var i int

	for i=0; i<len(dataTrain); i++ {
		dataTrainFloat64, _ = strconv.ParseFloat(dataTrain[i], 64)
		dataTrainSliceFloat64 = append(dataTrainSliceFloat64, dataTrainFloat64)
		fmt.Println("inside train loop")
		fmt.Println(dataTrain[i])
	}

	for i=0; i<len(dataKnow); i++ {
		dataKnowFloat64, _ = strconv.ParseFloat(dataKnow[i], 64)
		dataKnowSliceFloat64 = append(dataKnowSliceFloat64, dataKnowFloat64)
		fmt.Println("inside Know loop")
		fmt.Println(dataKnow[i])
	}

	return dataTrainSliceFloat64, dataKnowSliceFloat64
}