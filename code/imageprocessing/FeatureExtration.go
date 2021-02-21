package imageprocessing

import (
	"gocv.io/x/gocv"
	"fmt"
	"math"
	

)

func getGLCM(Image gocv.Mat, GLCM *gocv.Mat, delta_r int, delta_c int){
	
	auxGLCM			:= make([][]float64,256)
	
	for i := 0; i < 256; i++ {

		auxGLCM[i] = make([]float64,256)
		
		for j := 0; j < 256; j++ {
			auxGLCM[i][j] = 0
		}

	}
	
	var max float64 = 0.0

	var GLCM_row uint8 = 0
	var GLCM_col uint8 = 0

	/*Percorre a imagem calculando a ocorrência de cada padrão
	Não é necessário borda pois Delta_y e Delta_x limitam até que pixel se acessa a imagem*/

	for r := 0; r < (Image.Rows()-delta_r)	; r++ {
		for c := 0; c < (Image.Cols()-delta_c); c++ {
			GLCM_row = Image.GetUCharAt(r,c)
			GLCM_col = Image.GetUCharAt((r+delta_r),(c+delta_c))

			auxGLCM[GLCM_row][GLCM_col]++
		}
	}
	for r := 0; r < 256	; r++ {
		for c := 0; c < 256; c++ {
			if auxGLCM[r][c] > max{
				max =auxGLCM[r][c]
			}
		}
	}	

	for r := 0; r < (*GLCM).Rows()	; r++ {
		for c := 0; c < (*GLCM).Cols(); c++ {
			(*GLCM).SetUCharAt(r,c,uint8(255*(auxGLCM[r][c]/max)))
			//fmt.Print(auxGLCM[r][c],"  ")
		}
		//fmt.Println(" ")
	}
	//fmt.Println("max:                                                          ",max)
}
	
/**
 * [Take the GLCM and the means of a group of images]
 * @param {[type]} Images []gocv.mat [description]
 * @param {[type]} GLCMs  []gocv.Mat [description]
 * @param {[type]} means  []gocv.Mat [description]
 * @param {[type]} show   bool       [description]
 */
func GroupGLCM(Images []gocv.Mat, GLCMs *[]gocv.Mat, print bool ,show bool) {
	
	for i := 0; i < len(Images); i++ {
		if print{
			fmt.Println("Calculating GLCM   ", (i+1), "of ", len(Images))

		}

		getGLCM(Images[i], &(*GLCMs)[i], 1,0)

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

/**
 * [GroupCorrelation description]
 * @param {[type]} GLCMs        *[]gocv.Mat [description]
 * @param {[type]} Correlations []float64   [description]
 * @param {[type]} print        bool        [description]
 */
func GroupCorrelation(GLCMs *[]gocv.Mat, Correlations []float64, print bool){

	for i := 0; i < len(*GLCMs); i++ {
		
		if print{
			fmt.Println("Calculating Correlation:  ",(i+1), "of ",len(*GLCMs))
		}

		Correlations[i] = Correlation((*GLCMs)[i])	
	}

}

/**
 * [Correlation description]
 * @param {[type]} GLCM gocv.Mat) (float64 [description]
 */
func Correlation(GLCM gocv.Mat) float64{
	var Correlation float64 = 0

	muRow,muCol := getMu(GLCM)
	
	sigmaRow,sigmaCol := getSigma(GLCM,muRow,muCol)

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			Correlation += (float64(r)*float64(c))*float64(GLCM.GetUCharAt(r,c)) - (muRow*muCol)
		}
	}
	Correlation = Correlation/(sigmaRow*sigmaCol)
	return Correlation
}

/**
 * [getMu description]
 * @param  {[type]} GLCM gocv.Mat)     (float64,float64 [description]
 * @return {[type]}      [description]
 */
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

/**
 * [getSigma description]
 * @param  {[type]} GLCM  gocv.Mat      [description]
 * @param  {[type]} muRow float64       [description]
 * @param  {[type]} muCol float64)      (float64,float64 [description]
 * @return {[type]}       [description]
 */
func getSigma(GLCM gocv.Mat, muRow float64, muCol float64) (float64,float64){
	
	var sigmaRow float64 = 0
	var sigmaCol float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			sigmaRow += math.Pow(float64(r) - muRow,2) * float64(GLCM.GetUCharAt(r,c))
			sigmaCol += math.Pow(float64(c) - muCol,2) * float64(GLCM.GetUCharAt(r,c))
		}
	}
	return sigmaRow,sigmaCol
}

func Contrast(GLCM gocv.Mat) float64{

	var Contrast float64 = 0

	for g := 0; g < 256; g++ {
		for r := 0; r < GLCM.Rows()	; r++ {
			for c := 0; c < GLCM.Cols(); c++ {
				Contrast += float64(GLCM.GetUCharAt(r,c))
			}
		}
	}
	return Contrast
}