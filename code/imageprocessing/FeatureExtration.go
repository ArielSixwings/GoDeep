package imageprocessing

import (
	"gocv.io/x/gocv"
	"fmt"
	"math"
	

)

/**
 * [Take the GLCM and the means of a group of images]
 * @param {[type]} Images []gocv.mat [description]
 * @param {[type]} GLCMs  []gocv.Mat [description]
 * @param {[type]} means  []gocv.Mat [description]
 * @param {[type]} show   bool       [description]
 */
func GroupGLCM(Images []gocv.Mat, GLCMs []gocv.Mat, means []gocv.Mat, show bool) {
	
	GLCM := gocv.NewMat()

	mean := gocv.NewMat()
	
	for i := 0; i < len(Images); i++ {
		fmt.Println("Calculating GLCM   ", i, "of ", len(Images))
		gocv.CalcCovarMatrix(Images[i], &GLCM, &mean, gocv.CovarCols, Images[2].Type())

		GLCMs[i] = GLCM
		means[i] = mean
		if show {
			ShowImage("And this is yout image", GLCMs[i], 100)
		}
	}
}

/**
 * [getEnergy description]
 * @param  {[type]} GLCM gocv.Mat      [A GLCM gocv.Mat]
 * @return {[type]} float64 		   [Energy of the image that produced the GLCM]
 */
func Energy(GLCM gocv.Mat) float64{
	var Energy float64 = 0
	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			Energy += float64(math.Pow(float64(GLCM.GetUCharAt(r,c)),2))
		}
	}
	return Energy
}

func GroupEnergy(GLCMs []gocv.Mat, Energys []float64){

	for i := 0; i < len(GLCMs); i++ {
		fmt.Println("Calculating Energy:  ",i, "of ",len(GLCMs))
		Energys[i] = Energy(GLCMs[i])	
	}

}