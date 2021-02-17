package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"./imageprocessing"
	"gocv.io/x/gocv"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func main() {
	var files []string
	var name string
	nametemp := []string{"\"./","\""}
	var firtst bool = true
	testcolor := gocv.NewMat()

	root := "./imageprocessing/Images/danger"
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if( firtst){
			firtst = false
			continue
		}
		name = strings.Join(nametemp, file)
		fmt.Println(name)
		testcolor = imageprocessing.ReadImage(testcolor, file, true, false, true)
	}
}