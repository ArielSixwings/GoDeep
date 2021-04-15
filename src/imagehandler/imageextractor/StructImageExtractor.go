package imageextractor

import (
	"gocv.io/x/gocv"
	"../../basicdata"
)

type ImageExtractor struct {
	Images 		[]gocv.Mat
	split 		[][]string
	readOrigins []string
	Labelsize 	[]cartesian.Sizelabel
}