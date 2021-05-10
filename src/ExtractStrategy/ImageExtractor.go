package extract

import (
	"path/filepath"
	"gocv.io/x/gocv"
)

func (ie *ImageExtractor) Allocate() error{
	
	(*ie).Images = make([]gocv.Mat,(*ie).Readinfo.SizeData)
	
	for i := 0; i < (*ie).Readinfo.SizeData; i++ {
		(*ie).Images[i] = gocv.NewMat()
	}
	return nil
}

func (ie *ImageExtractor) ReadData(path string,dataindex int) error{

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

func (ie ImageExtractor) PresentData(Menssage string, index int, time int) {
	window := gocv.NewWindow(Menssage) //basic window
	window.IMShow(ie.Images[index])               //show the image
	window.WaitKey(time)
}

func (ie ImageExtractor) SaveData(i int,Name string) {
	gocv.IMWrite(Name, ie.Images[i]) //save the image
}