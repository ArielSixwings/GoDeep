//package main
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func escreverTexto(linhas [256000]string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprint(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}
func escreverTexto2(linhas []string, caminhoDoArquivo string) error {
	// Cria o arquivo de texto
	arquivo, err := os.Create(caminhoDoArquivo)
	// Caso tenha encontrado algum erro retornar ele
	if err != nil {
		return err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um escritor responsavel por escrever cada linha do slice no arquivo de texto
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha)
	}

	// Caso a funcao flush retorne um erro ele sera retornado aqui tambem
	return escritor.Flush()
}

func main() {

	var a [256000]string
	var c int = 1
	var caminho []string
	caminho = append(caminho, "-----------------------------------------------")

	for j := 0; j < 257; j++ {

		for i := c; i < 256000; i++ {
			randonInt := rand.Intn(9)
			if randonInt == 3 {
				a[i] = "_"
			} else if randonInt == 2 || randonInt == 5 {
				a[i] = "A"
			} else if randonInt == 1 || randonInt == 8 {
				a[i] = "C"
			} else if randonInt == 0 || randonInt == 4 {
				a[i] = "T"
			} else if randonInt == 6 || randonInt == 7 {
				a[i] = "G"
			}
			if i%257 == 0 && i != 0 {
				a[i] = ","
				c += 257
				continue
			}
		}
	}

	/*var chamada int = 0
	for j := 0; j < 256000; j++ {
		fmt.Print(a[j])
		if j%256 == 0 && j != 0 {
			fmt.Println()
			fmt.Println("chamada:", chamada)
			chamada = chamada + 1
		}*/

	for j := 0; j < 256000; j++ {
		err := escreverTexto(a, "teste.txt")

		if err != nil {
			log.Fatalf("Erro:", err)
		}
	}
}
