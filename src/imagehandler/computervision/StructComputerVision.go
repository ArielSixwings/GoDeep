package computervision

import(
	"../../basicdata"
	"gocv.io/x/gocv"
)

type FeatureType int

const (

	EnergyFeature FeatureType = 0

	ContrastFeature FeatureType = 1

	CorrelationFeature FeatureType = 2

	HomogeneityFeature FeatureType = 3

)

type ComputerVison struct {
	BaseImages 	[]gocv.Mat
	Information []cartesian.Features
}