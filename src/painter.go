package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func PaintCircle(Image gocv.Mat, center image.Point, radius int, color color.RGBA, thickness int) {
	gocv.Circle(&Image, center, radius, color, thickness)

	//func Circle(img *Mat, center image.Point, radius int, c color.RGBA := color.RGBA, thickness int)


func main(){
	var Image = gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)
	PaintCircle(Image.GetUCharAt(128,128),50 Image.)
}