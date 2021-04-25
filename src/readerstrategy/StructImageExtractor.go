package extract

import (
	"gocv.io/x/gocv"
)

type ImageExtractor struct {
	DataReader
	Images 		[]gocv.Mat
}