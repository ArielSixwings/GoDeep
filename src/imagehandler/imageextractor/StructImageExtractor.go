package imageextractor

import (
	"gocv.io/x/gocv"
)

type ImageExtractor struct {
	Images 	[]gocv.Mat
	split 	[]string
}