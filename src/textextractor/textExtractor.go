package textextractor

import (
	"bufio"
	"fmt"
	"os"
)

func (st *TextExtractor) allocate(size int) {
	(*st).texts = make([][]string, size+1)
	for i := 0; i < size+1; i++ {
		(*st).texts[i] = make([]string, size+1)
	}
}
func (st TextExtractor) PrintFile() {
	for j := 0; j < 1; j++ {
		for i := 0; i < len((st).texts)-1; i++ {
			fmt.Println(i, "  ", st.texts[j][i])
		}
	}
}

// Function that reads the contents of the file and returns a slice of the string with all lines of the file
func (st *TextExtractor) ScanText(filepath string) error {
	// Open the file
	file, err := os.Open(filepath)
	// If you have found an error when trying to open the file, return the error found
	if err != nil {
		return err
	}
	// Ensures that the file will be closed after use
	defer file.Close()
	// Creates a scanner that reads each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		(st).temp = append((st).temp, scanner.Text())
	}
	(*st).allocate(len((st).temp))
	fmt.Printf("cont1:%d\n", st.cont)

	for i := 0; i < len((st).temp); i++ {
		(*st).texts[(st).cont][i] = (*st).temp[i]
	}
	st.cont++

	fmt.Printf("cont2:%d\n", st.cont)
	(*st).temp = (*st).temp[:0]
	// Returns the lines read and an error if an error occurs in the scanner
	return scanner.Err()
}
