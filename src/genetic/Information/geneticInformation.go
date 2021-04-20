package geneticinformation

import (
	"fmt"
	"strings"
)

/*// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
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
}*/

func (st GeneticInformation) score(s1 []string, s2 []string, match int, mismatch int, gap int, i int, j int) int {
	var score = 0
	if s1[i-1] != "_" && s2[j-1] != "_" {
		if s1[i-1] == s2[j-1] {
			score = match
		} else {
			score = mismatch
		}
	} else {
		score = gap
	}
	return score
}
func (st GeneticInformation) maximum3(first int, second int, third int) int {
	if second > third {
		if second > first {
			return second
		}
		return first
	} else {
		if third > first {
			return third
		}
		return first
	}
}
func (st GeneticInformation) getMatrix(firstSequence string, secondSequence string, match int,
	mismatch int, gap int) [][]int {
	s := strings.Split(firstSequence, "")
	t := strings.Split(secondSequence, "")
	d := gap

	matrix := make([][]int, len(firstSequence)+10)
	for i := 0; i < len(firstSequence)+10; i++ {
		matrix[i] = make([]int, len(secondSequence)+10)
	}

	matrix[0][0] = 0
	if len(firstSequence) == len(secondSequence) {
		for i := 0; i <= len(firstSequence); i++ {
			matrix[i][0] = d * i
		}
		for j := 0; j <= len(secondSequence)+1; j++ {
			matrix[0][j] = d * j
		}
	}

	var c = 0
	for i := 1; i <= len(secondSequence); i++ {
		c++
		for j := 1; j <= len(firstSequence); j++ {
			Choice1 := matrix[i-1][j-1] + st.score(s, t, 1, -1, -1, j, c)
			Choice2 := matrix[i-1][j] + d
			Choice3 := matrix[i][j-1] + d
			matrix[i][j] = st.maximum3(Choice1, Choice2, Choice3)
		}
	}
	return matrix
}

func (st GeneticInformation) getScore(firstString string, secondString string, match int, mismatch int, gap int) int {
	s := strings.Split(firstString, "")
	t := strings.Split(secondString, "")
	d := gap

	F := make([][]int, len(firstString)+10)
	for i := 0; i < len(firstString)+10; i++ {
		F[i] = make([]int, len(secondString)+10)
	}

	F[0][0] = 0
	if len(firstString) == len(secondString) {
		for i := 0; i <= len(firstString); i++ {
			F[i][0] = d * i
		}
		for j := 0; j <= len(secondString)+1; j++ {
			F[0][j] = d * j
		}
	}

	var c = 0
	for i := 1; i <= len(secondString); i++ {
		c++
		for j := 1; j <= len(firstString); j++ {
			Choice1 := F[i-1][j-1] + st.score(s, t, 1, -1, -1, j, c)
			Choice2 := F[i-1][j] + d
			Choice3 := F[i][j-1] + d
			F[i][j] = st.maximum3(Choice1, Choice2, Choice3)
		}
	}
	/*for i := 0; i <= len(r2); i++ {
		for j := 0; j <= len(r1); j++ {
			fmt.Printf("%d ", F[i][j])
		}
		fmt.Printf(" \n")
	}*/

	if firstString == secondString {
		return F[len(firstString)][len(secondString)]
	} else if secondString > firstString {
		return F[len(secondString)][len(firstString)]
	} else {
		return F[len(firstString)][len(secondString)]
	}
}

func (st GeneticInformation) Alimented(A string, B string, gap int) (string, string) {
	var AlignmentA = ""
	var AlignmentB = ""
	var i int = len(A)
	var j int = len(B)
	var r1 = A
	var r2 = B
	s := strings.Split(r1, "")
	t := strings.Split(r2, "")
	d := gap
	var Matrix [][]int = st.getMatrix(A, B, 1, -1, -2)
	for i > 0 && j > 0 {
		var Score int = Matrix[i][j]
		var ScoreDiag int = Matrix[i-1][j-1]
		var ScoreUp int = Matrix[i][j-1]
		var ScoreLeft int = Matrix[i-1][j]
		if Score == ScoreDiag+st.score(s, t, 1, -1, -1, i, j) {
			AlignmentA = s[i-1] + AlignmentA
			AlignmentB = t[j-1] + AlignmentB
			i = i - 1
			j = j - 1
		} else if Score == ScoreLeft+d {
			AlignmentA = s[i-1] + AlignmentA
			AlignmentB = "-" + AlignmentB
			i = i - 1
		} else if Score == ScoreUp+d {
			AlignmentA = "-" + AlignmentA
			AlignmentB = t[j-1] + AlignmentB
			j = j - 1
		}
	}
	for i > 0 {
		AlignmentA = s[i-1] + AlignmentA
		AlignmentB = "-" + AlignmentB
		i = i - 1
	}
	for j > 0 {
		AlignmentA = "-" + AlignmentA
		AlignmentB = t[j-1] + AlignmentB
		j = j - 1
	}
	return AlignmentA, AlignmentB
}

func (st GeneticInformation) GetResult(firstSequence []string, secondSequence []string, numberFather int) {
	for i := 0; i < len(firstSequence); i++ {
		st.resultScore = st.getScore(firstSequence[numberFather], secondSequence[i], 1, -1, -2)
		if st.resultScore > 0 {
			fmt.Printf("O score do resultado foi:%d, posicao: %d\n", st.resultScore, i)
		}
	}
}

//Apagar
/*func main() {
	c, err := scanText("../Extractor/tapes/fathers.txt")
	d, err := scanText("../Extractor/tapes/childs.txt")
	//fmt.Printf("%s ", c[0], err)
	fmt.Println(err)
	var J int
	for i := 0; i < len(c); i++ {
		J = getScore(c[45], d[i], 1, -1, -2)
		if J > 0 {
			fmt.Printf("%d, posicao: %d\n", J, i)
		}
	}
}*/
