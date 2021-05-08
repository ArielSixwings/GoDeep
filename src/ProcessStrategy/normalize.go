package process

import(
	//"errors"
	"gocv.io/x/gocv"
)

type Normalize struct{
	alpha 	float64
	beta 	float64
	normtype gocv.NormType
}

func (this Normalize) Process(ip *ImageProcessing,index int) error{
	gocv.Normalize((*ip).GLCMs[index], &(*ip).NormalizedGLCMs[index], this.alpha, this.beta, this.normtype)
	return nil
}

func (this Normalize) Allocate(ip *ImageProcessing) error{
	(*ip).NormalizedGLCMs = make([]gocv.Mat,(*ip).Readinfo.SizeData)
	
	for i := 0; i < (*ip).Readinfo.SizeData; i++ {
		(*ip).NormalizedGLCMs[i] = gocv.NewMat()
	}
	return nil
}

func (this Normalize) Verify(setalpha float64, setbeta float64,setnormtype gocv.NormType) error{
	return nil
}

func (this *Normalize) SetParameters(setalpha float64, setbeta float64,setnormtype gocv.NormType) error{
	err := (*this).Verify(setalpha,setbeta,setnormtype)
	if err == nil{
		(*this).alpha = setalpha
		(*this).beta = setbeta
		(*this).normtype = setnormtype
		return nil
	} else{
		return err
	}
}