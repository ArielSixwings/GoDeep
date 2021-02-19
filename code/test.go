package main

import (
	"./imageprocessing"
	"gocv.io/x/gocv"
	// "fmt"

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

/**
 *
 * @brief      { function_description }
 * @return     { description_of_the_return_value }
 */
func main() {
	
	var size int
	
	size  = imageprocessing.FolderLength("./imageprocessing/Images/danger")

	Images := make([]gocv.Mat,size)

	GLCMs := make([]gocv.Mat,size)

	means := make([]gocv.Mat,size)

	imageprocessing.ReadFolder(Images,"./imageprocessing/Images/danger",true,false,false)
	
	//GroupGLCM(Images []gocv.Mat, GLCMs []gocv.Mat, means []gocv.Mat, show bool)
	imageprocessing.GroupGLCM(Images, GLCMs, means, true)
}
