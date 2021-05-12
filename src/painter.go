package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main(){
	var Image = gocv.IMRead("background.jpg", gocv.IMReadUnchanged)
	var center image.Point
	var thecolor color.RGBA
	var thesize image.Rectangle

	thesize.Min.X = 32 
	thesize.Min.Y = 480
	thesize.Max.X = 64
	thesize.Max.Y = 224
	
	thecolor.R = 255
	thecolor.G = 0
	thecolor.B = 0
	thecolor.A = 1
	
	center.X = 428
	center.Y = 128
	
	window := gocv.NewWindow("the image") //basic window
	
	window.IMShow(Image)               //show the image
	window.WaitKey(0)
	
	gocv.Circle(&Image, center, 40, thecolor, 80)
	gocv.Rectangle(&Image, thesize, thecolor, 32)
	
	window.IMShow(Image)               //show the image
	window.WaitKey(0)
}