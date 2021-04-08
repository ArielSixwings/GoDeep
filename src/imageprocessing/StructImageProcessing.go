package imageprocessing

import (
	"gocv.io/x/gocv"
)
type AllIpType int

const (

	AllGLCM AllIpType = 0

	AllNormalizedGLCM AllIpType = 1

)

type ImageProcessing struct {
	FilteredImages 	[]gocv.Mat
	GLCMs 			[]gocv.Mat
	NormalizedGLCMs []gocv.Mat
}