package genericdata

import(
	"../basicdata"
)
type Groupflag int

const (
	Trainflag Groupflag = 0

	Testflag Groupflag = 1

	Resultflag Groupflag = 2

	GreatestOccurrence Groupflag = 3

	Labelflag Groupflag = 4

	Statusflag Groupflag = 5

	Interestflag Groupflag = 6

	Centroidflag Groupflag = 7

	Centerdistflag Groupflag = 8

	Allcentroidflag Groupflag = 9
)

type DataSet struct {

	test 				[]cartesian.Features
	train 				[]cartesian.Features	
	result 				[]cartesian.Labeldist
	
	centroid 			[]cartesian.Centroidinfo 	//kmeans
	
	centerdist 			[]cartesian.Featurepoint
	
	sizelabel 			[]cartesian.Sizelabel
	allcentroid			cartesian.Allcenter
	
	is_sortedbydist 	[]bool
	is_sortedbycenter 	[]bool 

}

type learnStrategy interface {
	learn(ds *DataSet)
}

type ByDist []cartesian.Featurepoint