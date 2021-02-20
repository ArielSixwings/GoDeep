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
	
	//var normtype gocv.NormType = gocv.NormMinMax

	size  = imageprocessing.FolderLength("./imageprocessing/Images/danger")

	Images := make([]gocv.Mat,size)

	GLCMs := make([]gocv.Mat,size)

	//normalizedGLCMs := make([]gocv.Mat,size)

	means := make([]gocv.Mat,size)

	imageprocessing.ReadFolder(Images,"./imageprocessing/Images/danger",true,false,false)
	
	//GroupGLCM(Images []gocv.Mat, GLCMs []gocv.Mat, means []gocv.Mat, show bool)
	imageprocessing.GroupGLCM(Images, GLCMs, means, false)
	// var Energy float64 =  imageprocessing.Energy(GLCMs[1]) 
	// fmt.Println("Energy:    ", Energy)
	
	Energys := make([]float64,size)
	imageprocessing.GroupEnergy(GLCMs,Energys)
	for i := 0; i < size; i++ {
		fmt.Println("Energy:   ", Energys[i])
	}

	//func Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType)
	//min value of dst is alpha and max value of dst is beta
	// for i := 0; i < size; i++ {
	// 	gocv.Normalize(GLCMs[i], &normalizedGLCMs[i], 0.0, 255.0, normtype )
	// 	//imageprocessing.ShowImage("normalizedGLCMs", normalizedGLCMs[i], 100)
	// }
	
	 
	
}
