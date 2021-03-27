package imageprocessing

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gocv.io/x/gocv"
)
func (ie *ImageExtractor) Allocate(size int){
	(*ie).Images = make([]gocv.Mat,size)
}
/**
 * [ReadImage description: read an image following the parameters]
 * @param {[type]} Image     gocv.Mat [the image]
 * @param {[type]} path      string   [the path to the image]
 * @param {[type]} show      bool     [if it is true show the image]
 * @param {[type]} save      bool     [if it is true save the image]
 * @param {[type]} colorfull bool     [if it is true read the image as a 3 chanel rbg]
 * @return {[type]}   gocv.Mat        [the readed image]
 */
func (ie *ImageExtractor)  ReadImage(path string, show bool, save bool, colorfull bool, i ...int){

	ImagePath := filepath.Join(path) //set path to the base image
	if len((*ie).Images) == 0{
		(*ie).Allocate(75) 	//temporary solution
	}
	if colorfull {
		if len(i) > 0{
			(*ie).Images[i[0]] = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
		} else{
			(*ie).Images[0] = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB
		}
	} else {
		if len(i) > 0{
			(*ie).Images[i[0]] = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
		} else{
			(*ie).Images[0] = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale
		}
	}

	if show {
		if len(i) > 0{
			(*ie).ShowImage(i[0],"And this is yout image", 100)
		} else{
			(*ie).ShowImage(0,"And this is yout image", 100)
		}
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
func (ie *ImageExtractor) ReadFolder(folder string, print bool, show bool, colorfull bool) {
	
	var files []string
	var name string
	var firtst bool = true
	var i int
	nametemp := []string{"\"./", "\""}

	err := filepath.Walk(folder, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		if firtst {
			firtst = false
			i = 0
			continue
		}

		name = strings.Join(nametemp, file)

		if print {
			fmt.Println("geting image:     ", name)

		}

		(*ie).ReadImage(file, show, false, colorfull,i)

		i++
	}
}

/**
 * [ShowImage description: Shows the image]
 * @param {[type]} Menssage string   [menssage in the window]
 * @param {[type]} Image    gocv.Mat [image to be showed]
 * @param {[type]} time     int      [time of the window]
 */
func (ie ImageExtractor) ShowImage(i int, Menssage string, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(ie.Images[i])               //show the image
	window.WaitKey(time)
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
 * [visit description:]
 * @param  {[type]} files *[]string        [array of files names]
 * @return {[type]} filepath.WalkFunc      [parameter used at filepath.Walk()]
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

/**
 * [FolderLength description: get the number of files in the folder]
 * @param {[type]} folder string [name of folder]
 * @return {[type]} int          [lenght of the folder(number of files)]
 */
func FolderLength(folder string) int {
	var files []string

	err := filepath.Walk(folder, visit(&files))

	if err != nil {
		panic(err)
	}
	return (len(files) - 1)
}
