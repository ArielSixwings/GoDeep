package process

import(
	"errors"
	"gocv.io/x/gocv"
)
type GLCM struct{
	delta_r int
	delta_c int
}

func (this GLCM) Process(ip *ImageProcessing,index int) error { 
	if (*ip).GLCMs[index].Rows() == 0{
		return errors.New("glcm vector wasn't provided")
	}
	
	auxGLCM	:= make([][]float64,256)
	
	var max float64 = 0.0

	var GLCM_row uint8 = 0
	var GLCM_col uint8 = 0
	for i := 0; i < 256; i++ {

		auxGLCM[i] = make([]float64,256)
		
		for j := 0; j < 256; j++ {
			auxGLCM[i][j] = 0
		}

	}

	for r := 0; r < ((*ip).FilteredImages[index].Rows()-this.delta_r)	; r++ {
		for c := 0; c < ((*ip).FilteredImages[index].Cols()-this.delta_c); c++ {
			GLCM_row = (*ip).FilteredImages[index].GetUCharAt(r,c)
			GLCM_col = (*ip).FilteredImages[index].GetUCharAt((r+this.delta_r),(c+this.delta_c))

			auxGLCM[GLCM_row][GLCM_col]++
		}
	}
	for r := 0; r < 256	; r++ {
		for c := 0; c < 256; c++ {
			if auxGLCM[r][c] > max{
				max = auxGLCM[r][c]
			}
		}
	}
	for r := 0; r < (*ip).GLCMs[index].Rows()	; r++ {
		for c := 0; c < (*ip).GLCMs[index].Cols(); c++ {
			(*ip).GLCMs[index].SetUCharAt(r,c,uint8(255*(auxGLCM[r][c]/max)))
		}
	}
	return nil
}

func (this GLCM) Allocate(ip *ImageProcessing) error{
	(*ip).GLCMs = make([]gocv.Mat,(*ip).Readinfo.SizeData)
	
	for i := 0; i < (*ip).Readinfo.SizeData; i++ {
		(*ip).GLCMs[i] = gocv.NewMatWithSize(256, 256, gocv.MatTypeCV8U)
	}
	return nil
}

func (this GLCM) Verify(deltar int, deltac int) error{
	return nil
}

func (this *GLCM) SetParameters(deltar int, deltac int) error{
	err := (*this).Verify(deltar,deltac)
	if err == nil{
		(*this).delta_r = deltar
		(*this).delta_c = deltac
		return nil
	} else{
		return err
	}
}
