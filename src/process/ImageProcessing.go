package process

import (
	"gocv.io/x/gocv"
	"fmt"
	"errors"
	"../extractstrategy"

)

/**
 * [getGLCM description: compute the grey level co-occurrence matrice]
 * @param  {[type]} Image   gocv.Mat      [image that we will use to compute the glcm]
 * @param  {[type]} GLCM    *gocv.Mat     [glcm that will be computed]
 * @param  {[type]} delta_r int           [step taked in the row diretion]
 * @param  {[type]} delta_c int           [step taked in the col diretion]
 * @return {[type]}                       [error handling]
 */
func (ip *ImageProcessing) getGLCM( index int,delta_r int, delta_c int) error{
	if (*ip).GLCMs[index].Rows() == 0{
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
	

	for r := 0; r < ((*ip).FilteredImages[index].Rows()-delta_r)	; r++ {
		for c := 0; c < ((*ip).FilteredImages[index].Cols()-delta_c); c++ {
			GLCM_row = (*ip).FilteredImages[index].GetUCharAt(r,c)
			GLCM_col = (*ip).FilteredImages[index].GetUCharAt((r+delta_r),(c+delta_c))

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

	for r := 0; r < (*ip).GLCMs[index].Rows()	; r++ {
		for c := 0; c < (*ip).GLCMs[index].Cols(); c++ {
			(*ip).GLCMs[index].SetUCharAt(r,c,uint8(255*(auxGLCM[r][c]/max)))
			// auxGLCM[GLCM_row][GLCM_col] = auxGLCM[GLCM_row][GLCM_col]/sum
		}
	}
	return nil
}

func (ip *ImageProcessing) getNormalizedGLCM(i int,alpha float64, beta float64, typ gocv.NormType) error{
	gocv.Normalize((*ip).GLCMs[i], &(*ip).NormalizedGLCMs[i], alpha, beta, typ)
	return nil
}

func (ip *ImageProcessing) AllocateIpStructs(size int,alliptype AllIpType) error{
	switch alliptype {
	case AllGLCM:
		(*ip).GLCMs = make([]gocv.Mat,size)
	
		for i := 0; i < size; i++ {
			(*ip).GLCMs[i] = gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)
		}
	case AllNormalizedGLCM:
		(*ip).NormalizedGLCMs = make([]gocv.Mat,size)
	
		for i := 0; i < size; i++ {
			(*ip).NormalizedGLCMs[i] = gocv.NewMat()
		}
	default:
		return errors.New("invalid request of AllocateIpStructs method, unkown allocate flag")
	}
	return nil
}

/**
 * [GroupGLCM description: Take the GLCM and the means of a group of images]
 * @param {[type]} Images []gocv.mat      [group of images used to compute the glcm group]
 * @param {[type]} GLCMs  []gocv.Mat      [group of glcms]
 * @param {[type]} show   bool            [if its true, show the computed glcms]
 * @return {[type]}                       [error handling] 
 */
func (ip *ImageProcessing) GroupGLCM(print bool ,show bool) error{

	if len((*ip).FilteredImages) == 0{
		return errors.New("glcm group and images group weren't provided")
	}

	if len((*ip).GLCMs) == 0{
		(*ip).AllocateIpStructs(len((*ip).FilteredImages),AllGLCM) 
	}

	
	for i := 0; i < len((*ip).FilteredImages); i++ {
		if print{
			fmt.Println("Calculating GLCM   ", (i+1), "of ", len((*ip).FilteredImages))

		}

		(*ip).getGLCM(i,0,1)

		if show {
			//extract.ShowImage("GLCMs", (*ip).GLCMs[i], 100)
		}

	}
	return nil
}

func (ip *ImageProcessing) GroupNormalizedGLCM(alpha float64, beta float64, typ gocv.NormType,print bool ,show bool) error{

	if len((*ip).GLCMs) == 0{
		return errors.New("GLCMs werent provided")
	} else{
		if len((*ip).NormalizedGLCMs) == 0{
			(*ip).AllocateIpStructs(len((*ip).GLCMs),AllNormalizedGLCM)
		}
	}
	
	for i := 0; i < len((*ip).GLCMs); i++ {
		if print{
			fmt.Println("Calculating NormalizedGLCM   ", (i+1), "of ", len((*ip).GLCMs))

		}

		(*ip).getNormalizedGLCM(i,0.0,255.0,typ)

		if show {
			//extract.ShowImage("NormalizedGLCMs", (*ip).NormalizedGLCMs[i], 100)
		}

	}
	return nil
}

func (ip *ImageProcessing) GetImages(ie *extract.ImageExtractor){
	(*ip).FilteredImages = (*ie).Images
}