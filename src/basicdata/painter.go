package cartesian

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"fmt"
)

func PresentProcessedData(data Features){
	
	var thebackground gocv.Mat
	thebackground = gocv.IMRead("background.jpg", gocv.IMReadUnchanged)
	
	var(
		
		theRectangle image.Rectangle

		centergender image.Point
		centerclass image.Point

		rectanglecolor color.RGBA
		colorgender color.RGBA
		colorclass color.RGBA
	)

	theRectangle.Min.X = 64 
	theRectangle.Min.Y = 936
	theRectangle.Max.X = 128
	theRectangle.Max.Y = 936 //- int(data.Features[1]*2.0)

	centergender.X = 700
	centergender.Y = 750

	centerclass.X = 700
	centerclass.Y = 250


	
	purple(&rectanglecolor)
	pink(&colorgender)
	gold(&colorclass)
	
	//gocv.Circle(&background, centergender, 120, colorgender, 100)
	//gocv.Circle(&background, centerclass, 90, colorclass, 190)
	//gocv.Rectangle(&background, theRectangle, rectanglecolor, 64)

	fmt.Println("antes de new window")
	window := gocv.NewWindow("Processed Data")
	fmt.Println("depois de new window")
	window.IMShow(thebackground)
	window.WaitKey(100)
}