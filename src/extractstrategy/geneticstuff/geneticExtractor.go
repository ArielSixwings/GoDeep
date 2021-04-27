package extract

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func (st *GeneticExtractor) GetFathers() []string {
	return (*st).fathers
}
func (st *GeneticExtractor) GetChilds() []string {
	return (*st).childs
}

func (st GeneticExtractor) PrintFather() {
	fmt.Print(st.fathers)
}
func (st GeneticExtractor) PrintChild() {
	fmt.Print(st.childs)
}

func (st GeneticExtractor) writeText(lines []string, pathFile string) error {
	// Creates the text file
	file, err := os.Create(pathFile)
	// If you have found an error return it
	if err != nil {
		return err
	}
	// Ensures that the file will be closed after use
	defer file.Close()

	// Creates a writer responsible for writing each line of the slice in the text file
	Writer := bufio.NewWriter(file)
	for _, lines := range lines {
		fmt.Fprint(Writer, lines)
	}

	// If the flush function returns an error it will be returned here as well
	return Writer.Flush()
}

func (st GeneticExtractor) readText(pathFile string) ([]string, error) {
	// Open the file
	file, err := os.Open(pathFile)
	// If you have found an error when trying to open the file, return the error found
	if err != nil {
		return nil, err
	}
	// Ensures that the file will be closed after use
	defer file.Close()

	// Creates a scanner that reads each line of the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Returns the lines read and an error if an error occurs in the scanner
	return lines, scanner.Err()
}

func (st GeneticExtractor) GenerateDataSet(originString []string, destinationPath string, tam int) error {
	var list []string
	for i := 0; i < len(originString); i++ {
		list = append(list, originString[i])
	}
	file, err := os.Create(destinationPath)
	defer file.Close()
	for j := 0; j < tam; j++ {
		err := st.writeText(list, destinationPath)
		if err != nil {
			log.Fatalf("Erro: Error generating the data set")
			for index, line := range originString {
				fmt.Println(index, line)
			}
		}
	}
	return err
}

func (st *GeneticExtractor) GenerateStringFather() {
	var list []string
	var counter int = 1 // counter that determines the length of the string
	var i int
	for j := 0; j < 250; j++ {
		for i = counter; i < 64000; i++ {
			randonInt := rand.Intn(9)
			if randonInt == 3 {
				list = append(list, "_")
			} else if randonInt == 2 || randonInt == 5 {
				list = append(list, "A")
			} else if randonInt == 1 || randonInt == 8 {
				list = append(list, "C")
			} else if randonInt == 0 || randonInt == 4 {
				list = append(list, "T")
			} else if randonInt == 6 || randonInt == 7 {
				list = append(list, "G")
			}
			if i%256 == 0 && i != 0 {
				list = append(list, "\n")
				counter = counter + 256
				break
			}
		}
	}
	(*st).fathers = list

}

func (st *GeneticExtractor) GenerateStringChild(pathDataSetPai string) error {

	var textRead string

	// open the file
	file, err := os.Open(pathDataSetPai)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	// read line by line
	var replacement int = 1
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), textRead)

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
			if j == 255 && replacement != 1 {
				s[j] = s[j] + "\n"
			}
			(*st).childs = append((*st).childs, s[j])
			replacement = 2
		}
	}
	return err
}