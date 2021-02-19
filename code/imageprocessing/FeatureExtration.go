package imageprocessing

import (
	"gocv.io/x/gocv"
	// "fmt"

)

/**
 * [Take the GLCM and the means of a group of images]
 * @param {[type]} Images []gocv.Mat [description]
 * @param {[type]} GLCMs  []gocv.Mat [description]
 * @param {[type]} means  []gocv.Mat [description]
 * @param {[type]} show   bool       [description]
 */
func GroupGLCM(Images []gocv.Mat, GLCMs []gocv.Mat, means []gocv.Mat, show bool) {
	
	GLCM := gocv.NewMat()

	mean := gocv.NewMat()
	
	for i := 0; i < 50; i++ {
		gocv.CalcCovarMatrix(Images[i], &GLCM, &mean, gocv.CovarCols, Images[2].Type())

		GLCMs[i] = GLCM
		means[i] = mean

		if show {
			ShowImage("And this is yout image", GLCMs[i], 100)
		}
	}
}
