package generalizecartesian
import(
	"../basicdata"
)
type Groupflag int

const (
	Knowflag Groupflag = 0

	Trainflag Groupflag = 1

	Resultflag Groupflag = 2

	GreatestOccurrence Groupflag = 3

	Labelflag Groupflag = 4

	Statusflag Groupflag = 5

	Interestflag Groupflag = 6

	Centroidflag Groupflag = 7

	Centerdistflag Groupflag = 8

	Allcentroidflag Groupflag = 9
)

type Labelfeatures struct {

	train 				[]cartesian.Features
	know 				[]cartesian.Features	
	result 				[]cartesian.Labeldist
	interestgroup 		[]cartesian.Interest
	centroid 			[]cartesian.Centroidinfo
	centerdist 			[]cartesian.Featurepoint
	sizelabel 			[]cartesian.Sizelabel
	allcentroid			cartesian.Allcenter
	is_sortedbydist 	[]bool
	is_sortedbycenter 	[]bool 

}

type ByDist []cartesian.Featurepoint
