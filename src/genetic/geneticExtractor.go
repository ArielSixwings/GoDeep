package geneticextractor

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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

func escreverTexto(linhas []string, caminhoDoArquivo string) error {
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

func lerTexto(caminhoDoArquivo string) ([]string, error) {
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

func gerarDataSet(str []string, caminho string, tam int) error {
	var lista2 []string
	for i := 0; i < len(str); i++ {
		lista2 = append(lista2, str[i])
	}
	arquivo, err := os.Create(caminho)
	defer arquivo.Close()
	for j := 0; j < tam; /*64000*/ j++ {
		err := escreverTexto(lista2, caminho)
		if err != nil {
			log.Fatalf("Erro:", err)
			for indice, linha := range str {
				fmt.Println(indice, linha)
			}
		}
	}
	return err
}

func gerarStringPai(str []string) []string {
	var lista []string

	var c int = 1
	var i int
	for j := 0; j < 250; j++ {
		for i = c; i < 64000; i++ {
			randonInt := rand.Intn(9)
			if randonInt == 3 {
				lista = append(lista, "_")
			} else if randonInt == 2 || randonInt == 5 {
				lista = append(lista, "A")
			} else if randonInt == 1 || randonInt == 8 {
				lista = append(lista, "C")
			} else if randonInt == 0 || randonInt == 4 {
				lista = append(lista, "T")
			} else if randonInt == 6 || randonInt == 7 {
				lista = append(lista, "G")
			}
			if i%256 == 0 && i != 0 {
				lista = append(lista, "\n")
				c = c + 256
				break
			}
		}
	}
	return lista
}

func gerarStringFilho(dtPai string) ([]string, error) {
	var lista []string
	var textoLido string

	// open the file
	file, err := os.Open(dtPai)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	// read line by line
	var troca int = 1
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), textoLido)

		for j := 0; j < len(s); j++ {
			randonInt := rand.Intn(30)
			if randonInt == 3 || randonInt == 15 {
				if s[j] != "G" && s[j] != "," {
					s[j] = "G"
				}
			}
			if randonInt == 0 || randonInt == 25 {
				if s[j] != "C" && s[j] != "," {
					s[j] = "C"
				}
			}
			if randonInt == 7 || randonInt == 21 {
				if s[j] != "T" && s[j] != "," {
					s[j] = "T"
				}
			}
			if randonInt == 6 || randonInt == 28 {
				if s[j] != "A" && s[j] != "," {
					s[j] = "A"
				}
			}
			if randonInt == 9 {
				if s[j] != "_" && s[j] != "," {
					s[j] = "_"
				}
			}
			if j == 255 && troca != 1 {
				s[j] = s[j] + "\n"
			}
			lista = append(lista, s[j])
			troca = 2
		}
	}
	return lista, err
}

/*func main() {

	//var a []string
	//var b []string
	var c []string
	//fmt.Printf("%s", a)
	//b = gerarStringPai(a)
	//fmt.Printf("%s", b)
	c, err := gerarStringFilho("teste.txt")

	//gerarStringPai(&a)
	fmt.Printf("%s %s", c, err)
	gerarDataSet(c, "teste2.txt", 64000)
}
