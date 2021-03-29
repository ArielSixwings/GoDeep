package learnstrategy

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

type learnStrategy interface {
	Learn(ds *DataSet)
}

type DataLearner struct {

	test 				[]cartesian.Features
	train 				[]cartesian.Features	
	result 				[]cartesian.Labeldist
	
	sizelabel 			[]cartesian.Sizelabel

	interestgroup 		[]cartesian.Interest 		//knn
	
	centroid 			[]cartesian.Centroidinfo 	//kmeans
	centerdist 			[]cartesian.Featurepoint
	allcentroid			cartesian.Allcenter
	
	
	is_sortedbydist 	[]bool
	is_sortedbycenter 	[]bool

	Strategy learnStrategy
}


type ByDist []cartesian.Featurepoint