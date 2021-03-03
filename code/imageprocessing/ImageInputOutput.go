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
 * [ReadImage description: read an image following the parameters]
 * @param {[type]} Image     gocv.Mat [the image]
 * @param {[type]} path      string   [the path to the image]
 * @param {[type]} show      bool     [if it is true show the image]
 * @param {[type]} save      bool     [if it is true save the image]
 * @param {[type]} colorfull bool     [if it is true read the image as a 3 chanel rbg]
 * @return {[type]}   gocv.Mat        [the readed image]
 */
func ReadImage(Image gocv.Mat, path string, show bool, save bool, colorfull bool) gocv.Mat{

	ImagePath := filepath.Join(path) //set path to the base image

	if colorfull {
		Image = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
	} else {
		Image = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
	}

	if show {
		ShowImage("And this is yout image", Image, 100)
	}

	return Image
}

/**
 * [SaveImage description: Saves an image]
 * @param {[type]} Name  string   [name to be saved]
 * @param {[type]} Image gocv.Mat [the image to be saved]
 */
func SaveImage(Name string, Image gocv.Mat) {
	gocv.IMWrite(Name, Image) //save the image
}

/**
 * [ShowImage description: Shows the image]
 * @param {[type]} Menssage string   [menssage in the window]
 * @param {[type]} Image    gocv.Mat [image to be showed]
 * @param {[type]} time     int      [time of the window]
 */
func ShowImage(Menssage string, Image gocv.Mat, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(Image)               //show the image
	window.WaitKey(time)
}

/**
 * [visit description:]
 * @param  {[type]} files *[]string        [array of files names]
 * @return {[type]} filepath.WalkFunc      [parameter used at filepath.Walk()]
 */
func visit(files *[]string)  filepath.WalkFunc{
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

/**
 * [ReadFolder description: read all images at some folder]
 * @param {[type]} Images    *[]gocv.Mat [An Array of gocv.Mat that will be used to contain the images of the folder]
 * @param {[type]} folder    string      [folder name]
 * @param {[type]} print     bool        [if its true, print the names]
 * @param {[type]} show      bool        [if its true, show the images]
 * @param {[type]} colorfull bool        [if its is true take a 3 chanel rbg image]
 */
func ReadFolder(Images []gocv.Mat, folder string, print bool, show bool, colorfull bool) {
	var files []string
	var name string
	var firtst bool = true
	var i int
	nametemp := []string{"\"./","\""}
	
	
	tempimage := gocv.NewMat()


	err := filepath.Walk(folder, visit(&files))
	if err != nil {
		panic(err)
	}


	for _, file := range files {
		
		if( firtst){
			firtst = false
			i = 0 
			continue
		}
		
		name = strings.Join(nametemp, file)
		
		if(print){
			fmt.Println("geting image:     ", name)

		}
		
		tempimage = ReadImage(tempimage, file, show, false, colorfull)
		Images[i] = tempimage
		i++
	}
}

/**
 * [FolderLength description: get the number of files in the folder]
 * @param {[type]} folder string [name of folder]
 * @return {[type]} int          [lenght of the folder(number of files)]
 */
func FolderLength(folder string) int{
	var files []string

	err := filepath.Walk(folder, visit(&files))

	if err != nil {
		panic(err)
	}
	return 	((len(files) - 1)) 
}