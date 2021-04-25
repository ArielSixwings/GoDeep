package extract

import (
	"bufio"
	"fmt"
	"os"
)
func (st *TextExtractor) GetData(index int) []string {
	return (*st).texts[index]
}

func (st TextExtractor) PrintFile() {
	for j := 0; j < len(st.texts)-1; j++ {
		for i := 0; i < 250; i++ {
			if len(st.texts[j][i]) == 0 {
				fmt.Println("posicao zerada")
				break
			}
			fmt.Println("arquivo:", j, "linha:", i, "  ", st.texts[j][i])
		}
	}
}

// Function that reads the contents of the file and returns a slice of the string with all lines of the file
func (st *TextExtractor) ReadData(path string,dataindex int) error{
	// Open the file
	file, err := os.Open(path)
	// If you have found an error when trying to open the file, return the error found
	if err != nil {
		return err
	}
	// Ensures that the file will be closed after use
	defer file.Close()
	// Creates a scanner that reads each line of the file
	scanner := bufio.NewScanner(file)
	var temp []string
	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}
	for i := 0; i < len(temp); i++ {
		(*st).texts[dataindex][i] = temp[i]
	}

	temp = temp[:0]
	// Returns the lines read and an error if an error occurs in the scanner
	return scanner.Err()
}

func (st *TextExtractor) Allocate() error{
	
	(*st).texts = make([][]string,(*st).Readinfo.SizeData)
	
	for i := 0; i < (*st).Readinfo.SizeData; i++ {
		(*st).texts[i] = make([]string, 1000)
	}
	return nil
}

func (st TextExtractor)  PresentData(Menssage string, index int, time int) {
	fmt.Println("Thats a mission for Peter and Moreno!: ", Menssage,index,time)
}
/**
 * [SaveData description: Saves an image]
 * @param {[type]} Name  string   [name to be saved]
 * @param {[type]} Image gocv.Mat [the image to be saved]
 */
func (st TextExtractor)  SaveData(i int,Name string) {
	fmt.Println("Thats a mission for Peter and Moreno!: ", i,Name)
}