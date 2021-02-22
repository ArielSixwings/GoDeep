package main

import (
	"./imageprocessing"
	"./nonparametric"
	"gocv.io/x/gocv"
	"fmt"
	//"math"
)

func main() {
	
	var size int
	
	var normtype gocv.NormType = gocv.NormMinMax

	size  = imageprocessing.FolderLength(".././code/imageprocessing/Images/danger")

	Images 			:= make([]gocv.Mat,size)

	GLCMs 			:= make([]gocv.Mat,size)
	
	normalizedGLCMs	:= make([]gocv.Mat,size)
	
	Energys			:= make([]float64,size)

	Correlations	:= make([]float64,size)
	
	for i := 0; i < size; i++ {
		GLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedGLCMs[i]	= gocv.NewMat()
	}

	imageprocessing.ReadFolder(Images,".././code/imageprocessing/Images/danger",true,false,false)
	
	imageprocessing.GroupGLCM(Images, &GLCMs, true, true)

	for i := 0; i < GLCMs[0].Rows(); i++ {
		fmt.Println(GLCMs[0].GetUCharAt(i,0))
	}
	//func Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType)
	//min value of dst is alpha and max value of dst is beta
	for i := 0; i < size; i++ {
		gocv.Normalize(GLCMs[i], &normalizedGLCMs[i], 0.0, 255.0, normtype )

	}

	imageprocessing.GroupFeature(&normalizedGLCMs,Energys,imageprocessing.EnergyFeature, true)

	imageprocessing.GroupFeature(&normalizedGLCMs,Correlations,imageprocessing.CorrelationFeature, true)

	for i := 0; i < size; i++ {
		fmt.Println("Energy:   ", Energys[i])
	}

	
	for i := 0; i < size; i++ {
		fmt.Println("Correlation:   ", Correlations[i])
	}	 
	
}
