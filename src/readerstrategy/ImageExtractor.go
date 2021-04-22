package extractstrategy

import (
	"path/filepath"
	"gocv.io/x/gocv"
)
func (ie *ImageExtractor) Allocate(datareader DataReader) error{
	
	(*ie).Images = make([]gocv.Mat,datareader.Readinfo.SizeData)
	
	for i := 0; i < datareader.Readinfo.SizeData; i++ {
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
func (ie *ImageExtractor)  ReadData(datareader DataReader, i int) error{

	ImagePath := filepath.Join(datareader.readOrigins[0]) //set path to the base image
	
	if datareader.Format { 	//in that context, verify it should read a collorfull, or a gray image
		(*ie).Images[i] = gocv.IMRead(ImagePath, gocv.IMReadUnchanged) //read the base image as as RGB

	} else {
		(*ie).Images[i] = gocv.IMRead(ImagePath, gocv.IMReadGrayScale) //read the base image in grayscale

	}

	if datareader.Show {
		ShowImage("And this is yout image", (*ie).Images[i], 100)

	}
	return nil
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
 * [SaveImage description: Saves an image]
 * @param {[type]} Name  string   [name to be saved]
 * @param {[type]} Image gocv.Mat [the image to be saved]
 */
func (ie ImageExtractor) SaveImage(i int,Name string) {
	gocv.IMWrite(Name, ie.Images[i]) //save the image
}