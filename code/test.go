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

// func getEnergy(GLCM gocv.Mat) {
// 	var Energy float64 = 0
// 	for r := 0; r < GLCM.Rows()	; r++ {
// 		for c := 0; c < GLCM.Cols(); c++ {
// 			Energy += float64(math.Pow(float64(GLCM.GetUCharAt(r,c)),2))
// 		}
// 	}
// 	fmt.Println(Energy)
// 	//return Energy
// }

func main() {
	
	var size int
	
	var normtype gocv.NormType = gocv.NormMinMax

	size  = imageprocessing.FolderLength("./imageprocessing/Images/danger")

	Images 			:= make([]gocv.Mat,size)

	GLCMs 			:= make([]gocv.Mat,size)

	means			:= make([]gocv.Mat,size)
	
	normalizedGLCMs	:= make([]gocv.Mat,size)
	
	Energys			:= make([]float64,size)
	
	for i := 0; i < size; i++ {
		GLCMs[i]			= gocv.NewMat()	
		means[i]			= gocv.NewMat()
		normalizedGLCMs[i]	= gocv.NewMat()
	}




	imageprocessing.ReadFolder(Images,"./imageprocessing/Images/danger",true,false,false)
	
	imageprocessing.GroupGLCM(Images, &GLCMs, &means, true, true)

	//func Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType)
	//min value of dst is alpha and max value of dst is beta
	for i := 0; i < size; i++ {
		gocv.Normalize(GLCMs[i], &normalizedGLCMs[i], 0.0, 255.0, normtype )

	}

	imageprocessing.GroupEnergy(&normalizedGLCMs,Energys,true)

	for i := 0; i < size; i++ {
		fmt.Println("Energy:   ", Energys[i])
	}

	
	 
	
}
