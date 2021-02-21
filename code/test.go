package main

import (
	"./imageprocessing"
	"gocv.io/x/gocv"
	"fmt"
	//"math"
)
/** CovarFlags
	// CovarScrambled indicates to scramble the results.
	CovarScrambled CovarFlags = 0

	// CovarNormal indicates to use normal covariation.
	CovarNormal CovarFlags = 1

	// CovarUseAvg indicates to use average covariation.
	CovarUseAvg CovarFlags = 2

	// CovarScale indicates to use scaled covariation.
	CovarScale CovarFlags = 4

	// CovarRows indicates to use covariation on rows.
	CovarRows CovarFlags = 8

	// CovarCols indicates to use covariation on columns.
	CovarCols CovarFlags = 16
**/

func main() {
	
	var size int
	
	var normtype gocv.NormType = gocv.NormMinMax

	size  = imageprocessing.FolderLength("./imageprocessing/Images/danger")

	Images 			:= make([]gocv.Mat,size)

	GLCMs 			:= make([]gocv.Mat,size)
	
	normalizedGLCMs	:= make([]gocv.Mat,size)
	
	Energys			:= make([]float64,size)

	Correlations	:= make([]float64,size)
	
	for i := 0; i < size; i++ {
		GLCMs[i]			= gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)	
		normalizedGLCMs[i]	= gocv.NewMat()
	}




	imageprocessing.ReadFolder(Images,"./imageprocessing/Images/danger",true,false,false)
	
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
