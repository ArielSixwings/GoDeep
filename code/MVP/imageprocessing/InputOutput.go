package imageprocessing

import (
	"gocv.io/x/gocv"
	"path/filepath"
	"log"
	"strings"
	"os"
	"fmt"
)

/**
 * @brief      Reads an image. colorfull or not
 *
 * @param      Image      The image
 * @param      path       The path
 * @param      show       To show the image
 * @param      save       To save the image
 * @param      colorfull  To read as an rgb image
 *
 * @return     { An image os the type gocv.Mat }
 */
func ReadImage(Image gocv.Mat, path string, show bool, save bool, colorfull bool) gocv.Mat {

	ImagePath := filepath.Join(path) //set path to the base image

	if colorfull {
		Image = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
	} else {
		Image = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
	}

	if show {
		ShowImage("And this is yout image", Image, 0)
	}

	return Image
}

/**
 * @brief      Saves an image.
 *
 * @param      Name   The name
 * @param      Image  The image
 *
 * @return     { An image os the type gocv.Ma }
 */
func SaveImage(Name string, Image gocv.Mat) {
	gocv.IMWrite(Name, Image) //save the image
}

/**
 * @brief      Shows the image.
 *
 * @param      Menssage  The menssage in the window
 * @param      Image     The image
 * @param      time      The time of the window
 *
 * @return     { An image os the type gocv.Ma }
 */
func ShowImage(Menssage string, Image gocv.Mat, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(Image)               //show the image
	window.WaitKey(time)
}

/**
 * @brief     Visit files at some folder.
 *
 * @param      Files  []String of files
 *
 * @return     { some file paths }
 */
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func ReadFolder(Images *[]gocv.Mat, folder string, print bool, show bool, colorfull bool) {
	var files []string
	var name string
	var firtst bool = true
		
	nametemp := []string{"\"./","\""}
	
	
	tempimage := gocv.NewMat()


	err := filepath.Walk(folder, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		
		if( firtst){
			firtst = false
			continue
		}
		
		name = strings.Join(nametemp, file)
		
		if(print){
			fmt.Println("geting image:     ", name)

		}
		
		tempimage = ReadImage(tempimage, file, show, false, colorfull)
	}
}