package cartesian

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func PresentProcessedData(data Features){
	
	
	var(
		thebackground = gocv.IMRead("../src/basicdata/background.jpg", gocv.IMReadUnchanged)
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
	theRectangle.Max.Y = 936 - int(data.Features[2]*10.0)

	centergender.X = 700
	centergender.Y = 750

	centerclass.X = 700
	centerclass.Y = 250

	purple(&rectanglecolor)
	
	if data.Features[1] == 1{
		blue(&colorgender)
	} else{
		pink(&colorgender)
	}

	if data.Features[0] == 2{
		gold(&colorclass)
	}else{
		if data.Features[0] == 1{
			silver(&colorclass)
		}else{
			copper(&colorclass)
		}
	}
	
	gocv.Circle(&thebackground, centergender, 120, colorgender, 250)
	gocv.Circle(&thebackground, centerclass, 90, colorclass, 190)
	gocv.Rectangle(&thebackground, theRectangle, rectanglecolor, 64)

	window := gocv.NewWindow("Processed Data")
	window.IMShow(thebackground)
	window.WaitKey(100)

}