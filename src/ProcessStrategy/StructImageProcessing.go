package process

import (
	"gocv.io/x/gocv"
)
type AllIpType int

const (

	AllGLCM AllIpType = 0

	AllNormalizedGLCM AllIpType = 1

)
type processStrategy interface {
	Process(ip *ImageProcessing,index int) error
	Allocate(ip *ImageProcessing) error
	//SaveProcessedData() error 
	//PresentProcessing(Menssage string, index int, time int) error
}

type ImageProcessing struct {
	DataProcessing

	Strategy processStrategy

	FilteredImages 	[]gocv.Mat
	GLCMs 			[]gocv.Mat
	NormalizedGLCMs []gocv.Mat
}