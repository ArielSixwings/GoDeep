package textextractor

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (st *TextExtractor) getName() string {
	return (*st).name
}
func (ft *FolderExtractor) getName() string {
	return (*ft).name
}
func (st *TextExtractor) allocate(size int) {
	(*st).texts = make([][]string, size)
	for i := 0; i < size-1; i++ {
		(*st).texts[i] = make([]string, 250)
	}
}
func (st *TextExtractor) GetTexts(position int) []string {
	return (*st).texts[position]
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
func (st TextExtractor) verifycandidate(candidate []string) bool {
	if candidate[0] == ".." || candidate[0] == "." { //"../src/imagehandler/Images/danger" or "./Images/grass_1.png"
		return true
	} else {
		return false
	}
}
func (st *TextExtractor) SetOrigins(origins []string) ([]bool, error) {
	var originsIntegrity bool = true
	path := make([][]string, len(origins))
	statusorigins := make([]bool, len(origins))
	for i := 0; i < len(origins); i++ {
		path[i] = append(strings.Split(origins[i], "/"))
		statusorigins[i] = (*st).verifycandidate(path[i])
		if originsIntegrity {
			originsIntegrity = statusorigins[i]
		}
	}
	if originsIntegrity {
		(*st).readOrigins = origins
		return statusorigins, nil
	} else {
		return statusorigins, errors.New("There was an error to set the origins, path provided is not valid")
	}
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func (ft *FolderExtractor) ScanFolder(folder string, index ...int) {
	var first bool = true
	var i int
	err := filepath.Walk(folder, visit(&ft.files))
	if err != nil {
		panic(err)
	}
	ft.allocate(len(ft.files))
	for _, file := range ft.files {
		if first {
			if len(index) > 0 {
				i = index[0]
			} else {
				i = 0
			}
			first = false
			continue
		}
		(*ft).ScanText(file, i)
		i++
	}
}
func (ft FolderExtractor) PrintNameFolder() {
	for j := 0; j < len(ft.files); j++ {
		fmt.Print((ft).files[j])
	}
}
func (st TextExtractor) PrintNameFiles() {
	for j := 0; j < len(st.name); j++ {
		fmt.Print((st).name[j])
	}
}

// Function that reads the contents of the file and returns a slice of the string with all lines of the file
func (st *TextExtractor) ScanText(filepath string, index int) error {
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
	var temp []string
	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}
	for i := 0; i < len(temp); i++ {
		(*st).texts[index][i] = temp[i]
	}

	temp = temp[:0]
	// Returns the lines read and an error if an error occurs in the scanner
	return scanner.Err()
}
