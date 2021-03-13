package statistics

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func ScanText(filePath string) ([]string, error) {
	var lines []string
	// Abre o arquivo
	arquivo, err := os.Open(filePath)

	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		defer arquivo.Close()
		fmt.Println(err)
		return lines, errors.New("Ocurred a ploblem during scan the path")
	}

	// Cria um scanner que le cada linha do arquivo
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()
	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	return lines, scanner.Err()
}

func PrintText(lines []string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}