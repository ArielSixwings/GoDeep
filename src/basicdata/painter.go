package cartesian

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"fmt"
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

func PresentGeneticData(tape string, index int){
	var(
		thebackground = gocv.IMRead("../src/basicdata/dnabackground.png", gocv.IMReadUnchanged)
		
		theRectangle image.Rectangle

		thecolor color.RGBA
	)

	theRectangle.Min.X = 76 
	theRectangle.Max.X = 140
	theRectangle.Max.Y = 60 

	for i := 0; i < 26; i++ {
		
		theRectangle.Min.Y = theRectangle.Max.Y+10
		theRectangle.Max.Y = theRectangle.Min.Y+5
		
		switch tape[index+i] {
		case 'A':
			adenina(&thecolor)
		case 'C':
			citosina(&thecolor)
		case 'T':
			timina(&thecolor)
		case 'G':
			guanina(&thecolor)	
		default:
			fmt.Println("Gap at genetic read")
		}
		gocv.Rectangle(&thebackground, theRectangle, thecolor, 5)
	}

	window := gocv.NewWindow("Genetic Data")
	window.IMShow(thebackground)
	window.WaitKey(100)
}