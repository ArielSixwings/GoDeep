package cartesian


type Interest struct {
	Interestlabel []string
	Interestdist  []float64
}

type Featurepoint struct {
	Dist      float64
	Distlabel string
}

type Labeldist struct {
	F_point            []Featurepoint
	Learnedlabel       string
	Status             bool
	Greatestoccurrence int
}

type Features struct {
	Features [3]float64
	Label    string
}

type Sizelabel struct {
	Label  string
	Size_l int
}

type ReadInformation struct {
	SizeData int
	Labelsize 	[]cartesian.Sizelabel
}

type Centroidinfo struct {

	Features [3]float64
	Radius   float64
	Label    string
}

type Allcenter struct {
	Features  [3]float64
	Maxradius float64
	Minradius float64
}