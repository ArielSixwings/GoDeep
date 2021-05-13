package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func PresentProcessedData(){
	var(
		background = gocv.IMRead("background.jpg", gocv.IMReadUnchanged)
		
		theRectangle image.Rectangle

		centergender image.Point
		centerclass image.Point

		rectanglecolor color.RGBA
		colormale color.RGBA
		colorclass color.RGBA
	)

	theRectangle.Min.X = 64 
	theRectangle.Min.Y = 936
	theRectangle.Max.X = 128
	theRectangle.Max.Y = 132 - Age

	centergender.X = 700
	centergender.Y = 750

	centerclass.X = 700
	centerclass.Y = 250

	gocv.Circle(&background, centergender, 120, colormale, 240)
	gocv.Circle(&background, centerclass, 90, colorclass, 190)
	gocv.Rectangle(&background, theRectangle, rectanglecolor, 64)
}


func main(){
	var centergender image.Point
	var centerclass image.Point

	var 

	

	
	window := gocv.NewWindow("the image") //basic window
	
	window.IMShow(Image)               //show the image
	window.WaitKey(0)
	
	
	window.IMShow(Image)               //show the image
	window.WaitKey(0)
}