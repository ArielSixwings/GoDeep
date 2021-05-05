package extract

import (
	"path/filepath"
	"gocv.io/x/gocv"
	"fmt"
)

func (ie ImageExtractor) Sifufu() {
	fmt.Println("que bagulho chato")
}

func (ie *ImageExtractor) Allocate() error{
	
	(*ie).Images = make([]gocv.Mat,(*ie).Readinfo.SizeData)
	
	for i := 0; i < (*ie).Readinfo.SizeData; i++ {
		(*ie).Images[i] = gocv.NewMat()
	}
	return nil
}
/**
 * [ReadImage description: read an image following the parameters]
 * @param {[type]} Image     gocv.Mat [the image]
 * @param {[type]} path      string   [the path to the image]
 * @param {[type]} show      bool     [if it is true show the image]
 * @param {[type]} colorfull bool     [if it is true read the image as a 3 chanel rbg]
 * @return {[type]}   gocv.Mat        [the readed image]
 */

//func (thek *Knn) Learn(DataLearner *learnstrategy.DataLearner){

//func (ie *ImageExtractor)  ReadData(path string, show bool, colorfull bool, i int){
func (ie *ImageExtractor)  ReadData(path string,dataindex int) error{

	ImagePath := filepath.Join(path) //set path to the base image
	
	
	if (*ie).Format { 	//in that context, verify it should read a collorfull, or a gray image
		(*ie).Images[dataindex] = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB

	} else {
		(*ie).Images[dataindex] = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale

	}

	if (*ie).Show {
		(*ie).PresentData("And this is yout image", dataindex, 100)

	}
	return nil
}

/**
 * [PresentData description: Shows the image]
 * @param {[type]} Menssage string   [menssage in the window]
 * @param {[type]} Image    gocv.Mat [image to be showed]
 * @param {[type]} time     int      [time of the window]
 */
func (ie ImageExtractor) PresentData(Menssage string, index int, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(ie.Images[index])               //show the image
	window.WaitKey(time)
}
/**
 * [SaveData description: Saves an image]
 * @param {[type]} Name  string   [name to be saved]
 * @param {[type]} Image gocv.Mat [the image to be saved]
 */
func (ie ImageExtractor) SaveData(i int,Name string) {
	gocv.IMWrite(Name, ie.Images[i]) //save the image
}