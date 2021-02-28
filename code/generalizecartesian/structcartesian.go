package generalizecartesian

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

)

type interest struct{

	interestlabel 		[]string
	interestdist 	[]float64
}

type featurepoint struct{
	dist float64
	distlabel string
}

type labeldist struct{

	f_point				[]featurepoint
	learnedlabel 		string
	status 				string
	greatestoccurrence 	int
}

type features struct {

	features 	[3]float64 
	label 		string
}

type Sizelabel struct{
	Label 	string
	Size_l 	int
}

type Labelfeatures struct {

	train 			[]features
	know 			[]featuresS
	result 			[]labeldist
	interestgroup 	[]interest
	is_sortedbydist []bool
	centroid 		[]features
	centerdist 		[]featurepoint
	sizelabel 		[]Sizelabel

}

type ByDist []featurepoint

