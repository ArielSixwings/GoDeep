package imageprocessing

import (
	"gocv.io/x/gocv"
	"fmt"
	"math"
	"errors"

)

/**
 * [getGLCM description: compute the grey level co-occurrence matrice]
 * @param  {[type]} Image   gocv.Mat      [image that we will use to compute the glcm]
 * @param  {[type]} GLCM    *gocv.Mat     [glcm that will be computed]
 * @param  {[type]} delta_r int           [step taked in the row diretion]
 * @param  {[type]} delta_c int           [step taked in the col diretion]
 * @return {[type]}                       [error handling]
 */
func getGLCM(Image gocv.Mat, GLCM *gocv.Mat, delta_r int, delta_c int) error{
	if (*GLCM).Rows() == 0{
		return errors.New("glcm vector wasn't provided")
	}
	
	auxGLCM			:= make([][]float64,256)
	
	//var sum float64 = 0
	var max float64 = 0.0

	var GLCM_row uint8 = 0
	var GLCM_col uint8 = 0
	for i := 0; i < 256; i++ {

		auxGLCM[i] = make([]float64,256)
		
		for j := 0; j < 256; j++ {
			auxGLCM[i][j] = 0
		}

	}
	

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
				max = auxGLCM[r][c]
			}
			//sum+= auxGLCM[r][c]
		}
	}	

	for r := 0; r < (*GLCM).Rows()	; r++ {
		for c := 0; c < (*GLCM).Cols(); c++ {
			(*GLCM).SetUCharAt(r,c,uint8(255*(auxGLCM[r][c]/max)))
			// auxGLCM[GLCM_row][GLCM_col] = auxGLCM[GLCM_row][GLCM_col]/sum
		}
	}
	return nil
}
	
/**
 * [GroupGLCM description: Take the GLCM and the means of a group of images]
 * @param {[type]} Images []gocv.mat      [description]
 * @param {[type]} GLCMs  []gocv.Mat      [description]
 * @param {[type]} means  []gocv.Mat      [description]
 * @param {[type]} show   bool            [description]
 * @return {[type]}                       [error handling] 
 */
func GroupGLCM(Images []gocv.Mat, GLCMs *[]gocv.Mat, print bool ,show bool) error{

	if len(*GLCMs) == 0{
		if len(Images) == 0{
			return errors.New("glcm group and images group weren't provided")
		}
		return errors.New("glcms group wasn't provided")
	} else{
		if len(Images) == 0{
			return errors.New("images group wasn't provided")
		}
	}
	
	for i := 0; i < len(Images); i++ {
		if print{
			fmt.Println("Calculating GLCM   ", (i+1), "of ", len(Images))

		}

		getGLCM(Images[i], &(*GLCMs)[i], 0,1)

		if show {
			ShowImage("GLCMs", (*GLCMs)[i], 100)
		}

	}
	return nil
}

/**
 * [energy description:]
 * @param  {[type]} GLCM gocv.Mat      [A GLCM gocv.Mat]
 * @return {[type]} float64            [energy of the image that produced the GLCM]
 */
func energy(GLCM gocv.Mat) float64{

	var energy float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			energy += float64(math.Pow(float64(GLCM.GetUCharAt(r,c)),2))
		}
	}
	return energy
}

/**
 * [correlation description:]
 * @param  {[type]} GLCM gocv.Mat      [A GLCM gocv.Mat]
 * @return {[type]} float64            [correlation of the image that produced the GLCM]
 */
func correlation(GLCM gocv.Mat) float64{
	var correlation float64 = 0

	muRow,muCol := getMu(GLCM)
	
	sigmaRow,sigmaCol := getSigma(GLCM,muRow,muCol)

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			correlation += (float64(r)*float64(c))*float64(GLCM.GetUCharAt(r,c)) - (muRow*muCol)
		}
	}
	correlation = correlation/(sigmaRow*sigmaCol)
	return correlation
}

/**
 * [homogeneity description:]
 * @param  {[type]} GLCM gocv.Mat      [A GLCM gocv.Mat]
 * @return {[type]} float64            [homogeneity of the image that produced the GLCM]
 */
func homogeneity(GLCM gocv.Mat) float64{

	var homogeneity float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			homogeneity += (1/(1+math.Pow(float64(r-c),2)))*float64(GLCM.GetUCharAt(r,c))
		}
	}
	return homogeneity
}

/**
 * [contrast description:]
 * @param {[type]} GLCM gocv.Mat [description]
 * @return {[type]}       [description]
 */
func contrast(GLCM gocv.Mat) float64{

	var contrast float64 = 0

	for r := 0; r < GLCM.Rows()	; r++ {
		for c := 0; c < GLCM.Cols(); c++ {
			contrast += math.Pow(float64(r-c),2) * float64(GLCM.GetUCharAt(r,c))
		}
	}

	return contrast
}

/**
 * [getMu description:]
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
 * [getSigma description:]
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

/**
 * [GroupFeature description: Calculate the energy of some group of images]
 * @param {[type]} GLCMs   *[]gocv.Mat    [Group of images GLCMs]
 * @param {[type]} Energys []float64      [Respectives Energys]
 * @param {[type]} print   bool           [if its true, print progress]
 * @return {[type]}                       [error handling]
 */
func GroupFeature(GLCMs *[]gocv.Mat, Features []float64,featuretype FeatureType, print bool) error{
	
	if len(Features) == 0 {
		return errors.New("Features weren't computed")
	}

	for i := 0; i < len(*GLCMs); i++ {

		switch featuretype {
		case EnergyFeature :
			if print{
				fmt.Println("Calculating Energy:  ",(i+1), "of ",len(*GLCMs))
			}
			Features[i] = energy((*GLCMs)[i])	

		case ContrastFeature :
			if print{
				fmt.Println("Calculating Contrast:  ",(i+1), "of ",len(*GLCMs))
			}
			Features[i] = contrast((*GLCMs)[i])

		case CorrelationFeature :
			if print{
				fmt.Println("Calculating Correlation:  ",(i+1), "of ",len(*GLCMs))
			}
			Features[i] = correlation((*GLCMs)[i])
		case HomogeneityFeature :
			if print{
				fmt.Println("Calculating Homogeneity:  ",(i+1), "of ",len(*GLCMs))
			}
			Features[i] = homogeneity((*GLCMs)[i])			
		}
	}
	return nil
}