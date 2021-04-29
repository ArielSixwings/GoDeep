package process

import(
	"../basicdata"
)

type learnStrategy interface {
	Process(dp *DataProcessing)
}

type DataProcessing struct {

}


type ByDist []cartesian.Featurepoint