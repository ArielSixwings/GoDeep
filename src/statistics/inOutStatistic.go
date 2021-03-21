package main

import (
	"fmt"
	"bufio"
//	"errors"
	"os"
//	"path/filepath"
	"strings"
	"strconv"
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

func splitText(conteudo []string) {
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

	convertData(pAge)
}


func convertData(data []string) {
	var pDataFloat64 float64
	var pDataSliceFloat64 []float64
	var i int

	for i=0; i<len(data); i++ {
		pDataFloat64, _ = strconv.ParseFloat(data[i], 64)
		pDataSliceFloat64 = append(pDataSliceFloat64, pDataFloat64)
	}

	
	
	fmt.Println(pDataSliceFloat64)

}


//###### APAGAR #######
func main() {
	var conteudo []string
	conteudo, err := scanText("tempTrain.csv")
	printText(conteudo, err)
	splitText(conteudo)
	
}