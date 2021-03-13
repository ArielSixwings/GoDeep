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

	Allcentroidflag Groupflag = 9
)

type interest struct {
	interestlabel []string
	interestdist  []float64
}

type featurepoint struct {
	dist      float64
	distlabel string
}

type labeldist struct {
	f_point            []featurepoint
	learnedlabel       string
	status             bool
	greatestoccurrence int
}

type features struct {
	features [3]float64
	label    string
}

type Sizelabel struct {
	Label  string
	Size_l int
}
type centroidinfo struct {
	features [3]float64
	radius   float64
	label    string
}

type allcenter struct {
	features  [3]float64
	maxradius float64
	minradius float64
}

type Labelfeatures struct {
	train             []features
	know              []features
	result            []labeldist
	interestgroup     []interest
	is_sortedbydist   []bool
	is_sortedbycenter []bool
	centroid          []centroidinfo
	centerdist        []featurepoint
	sizelabel         []Sizelabel
	allcentroid       allcenter
}

type ByDist []featurepoint
