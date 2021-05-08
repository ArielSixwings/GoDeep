package computervision

import (
	"fmt"
	"math"
	"../basicdata"
	"../processStrategy"
	"errors"

)

func (cv *ComputerVison) GetBaseImages(ip *process.ImageProcessing){
	(*cv).BaseImages = (*ip).NormalizedGLCMs
}

func (cv *ComputerVison) AllocateCvStructs(size int){
	(*cv).Information = make([]cartesian.Features, size)
}

func (cv *ComputerVison) correlation(i int) float64{
	var correlation float64 = 0

	muRow,muCol := (*cv).getMu(i)
	
	sigmaRow,sigmaCol := (*cv).getSigma(i,muRow,muCol)

	for r := 0; r < (*cv).BaseImages[i].Rows()	; r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			correlation += (float64(r)*float64(c))*float64((*cv).BaseImages[i].GetUCharAt(r,c)) - (muRow*muCol)
		}
	}
	correlation = correlation/(sigmaRow*sigmaCol)
	return correlation
}

func (cv *ComputerVison) homogeneity(i int) float64{

	var homogeneity float64 = 0

	for r := 0; r < (*cv).BaseImages[i].Rows(); r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			homogeneity += (1/(1+math.Pow(float64(r-c),2)))*float64((*cv).BaseImages[i].GetUCharAt(r,c))
		}
	}
	return homogeneity
}

func (cv *ComputerVison) contrast(i int) float64{

	var contrast float64 = 0

	for r := 0; r < (*cv).BaseImages[i].Rows()	; r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			contrast += math.Pow(float64(r-c),2) * float64((*cv).BaseImages[i].GetUCharAt(r,c))
		}
	}

	return contrast
}

func (cv *ComputerVison) energy(i int) float64{

	var energy float64 = 0

	for r := 0; r < (*cv).BaseImages[i].Rows()	; r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			energy += float64(math.Pow(float64((*cv).BaseImages[i].GetUCharAt(r,c)),2))
		}
	}
	return energy
}

func (cv *ComputerVison) getMu(i int) (float64,float64){
	
	var muRow float64 = 0
	var muCol float64 = 0

	for r := 0; r < (*cv).BaseImages[i].Rows()	; r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			muRow += float64(r) * float64((*cv).BaseImages[i].GetUCharAt(r,c))
			muCol += float64(c) * float64((*cv).BaseImages[i].GetUCharAt(r,c))
		}
	}
	return muRow,muCol
}

func (cv *ComputerVison) getSigma(i int, muRow float64, muCol float64) (float64,float64){
	
	var sigmaRow float64 = 0
	var sigmaCol float64 = 0

	for r := 0; r < (*cv).BaseImages[i].Rows()	; r++ {
		for c := 0; c < (*cv).BaseImages[i].Cols(); c++ {
			sigmaRow += math.Pow(float64(r) - muRow,2) * float64((*cv).BaseImages[i].GetUCharAt(r,c))
			sigmaCol += math.Pow(float64(c) - muCol,2) * float64((*cv).BaseImages[i].GetUCharAt(r,c))
		}
	}
	return sigmaRow,sigmaCol
}

func (cv *ComputerVison) selectfeature(i int,featuretype FeatureType) (error,float64){
		switch featuretype {
		case EnergyFeature :
			return nil,(*cv).energy(i)	

		case ContrastFeature :
			return nil,(*cv).contrast(i)

		case CorrelationFeature :
			return nil,(*cv).correlation(i)
		case HomogeneityFeature :
			return nil,(*cv).homogeneity(i)			
		default:
			return errors.New("invalid request of AllocateIpStructs method, unkown allocate flag"),0.0
		}
}

func (cv *ComputerVison) GroupFeature(print bool,featuretype ...FeatureType) error{
	
	if len((*cv).Information) == 0 {
		(*cv).AllocateCvStructs(len((*cv).BaseImages))
	}

	for i := 0; i < len((*cv).BaseImages); i++ {
		if print{
			fmt.Println("Calculating Features:  ",(i+1), "of ",len((*cv).BaseImages))
		}
		for j := 0; j < len(featuretype); j++ {
			_,(*cv).Information[i].Features[j] = (*cv).selectfeature(i,featuretype[j])
		}
	}
	return nil
}

func (cv ComputerVison) PrintFeatures(){
	for i := 0; i < len(cv.Information); i++ {
		fmt.Println(cv.Information[i])
	}
}