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
func GroupGLCM(Images []gocv.Mat, GLCMs *[]gocv.Mat, means *[]gocv.Mat, print bool ,show bool) {
	
	for i := 0; i < len(Images); i++ {
		if print{
			fmt.Println("Calculating GLCM   ", (i+1), "of ", len(Images))

		}

		gocv.CalcCovarMatrix(Images[i], &(*GLCMs)[i], &(*means)[i], gocv.CovarCols, Images[2].Type())

		if show {
			fmt.Println((i+1))
			ShowImage("GLCMs", (*GLCMs)[i], 100)
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

/**
 * [Calculate the energy of some group of images]
 * @param {[type]} GLCMs   *[]gocv.Mat [Group of images GLCMs]
 * @param {[type]} Energys []float64   [Respectives Energys]
 * @param {[type]} print   bool        [if its true, print progress]
 */
func GroupEnergy(GLCMs *[]gocv.Mat, Energys []float64, print bool){

	for i := 0; i < len(*GLCMs); i++ {
		
		if print{
			fmt.Println("Calculating Energy:  ",(i+1), "of ",len(*GLCMs))
		}

		Energys[i] = Energy((*GLCMs)[i])	
	}

}

func Contrast(GLCM gocv.Mat) float64{
	var Contrast float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			Contrast += float64(GLCM.GetUCharAt(r,c))
		}
	}
	return Contrast
}

func Correlation(GLCM gocv.Mat) float64{
	var Correlation float64 = 0

	muRow,muCol := getMu(GLCM)
	
	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			Correlation += float64((r*c)*GLCM.GetUCharAt(r,c)) - (muRow*muCol)
		}
	}
	Correlation = Correlation/()
	return Correlation
}

func getMu(GLCM gocv.Mat) (float64,float64){
	
	var muRow float64 = 0
	var muCol float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			muRow += float64(r) * float64(GLCM.GetUCharAt(r,c))
			muCol += float64(c) * float64(GLCM.GetUCharAt(r,c))
		}
	}
	return muRow,muCol
}

func getx(GLCM gocv.Mat) (float64,float64){
	
	var muRow float64 = 0
	var muCol float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			muRow += float64(r) * float64(GLCM.GetUCharAt(r,c))
			muCol += float64(c) * float64(GLCM.GetUCharAt(r,c))
		}
	}
	return muRow,muCol
}