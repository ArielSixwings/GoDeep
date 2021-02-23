package generalizeimage

type labeldist struct{

	dist []float64

	learnedlabel string

	greatestoccurrence int
}

type features struct {

	features 	[3]float64 
	
	label 		string
}


type labelfeatures struct {

	study 		[]features

	know 		[]features
	
	result 	[]labeldist
}

/**
 * [func description]
 * @param  {[type]} lf *labelfeatures) calcdistance( [description]
 * @return {[type]}    [description]
 */

type ByDistance []labeldist

func (d Bydist) Len() int { return len(d) }
func (d Bydist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d Bydist) Less(i, j int) bool { return d.dist[i] < d.dist[j] }

func (lf *labelfeatures) calcdistance() { 

	var sum float64 = 0.0

	for i := 0; i < len((*lf).study); i++ {
		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).know[j].features[f] - (*lf).study[j].features[f],2))
			}			
			(*lf).result[i].dist[j] = math.Sqrt(sum)
			(*lf).result[i].learnedlabel[j] = (*lf).know[j].label
		}
	}

}

type labelsize struct{
	label 	string
	size 	int
}

type groupflag int

const (

	knowflag groupflag = 0

	trainflag groupflag = 1

)

func generalize_for_nonparametric(LabelFeatures *labelfeatures, feature_X []float64, feature_Y []float64, feature_Z []float64,ls []labelsize,group groupflag){

	var size int
	if group == know{
		size len((*LabelFeatures).know)
	} else{
		size len((*LabelFeatures).train)
	}

	for i := 0; i < size; i++ {
		if group == knowflag{
			(*LabelFeatures).know[i] = make([]features,size)
			
			(*LabelFeatures).know[i].features[0] = feature_X[i]
			(*LabelFeatures).know[i].features[1] = feature_Y[i]
			(*LabelFeatures).know[i].features[2] = feature_Z[i]
			
			for j := 0; j < len(ls); j++ {
				for s := 0; s < ls[j].size; s++ {
					(*LabelFeatures).know[i].label = ls[s].label
				}
			}
		} else{
			(*LabelFeatures).train[i] = make([]features,size)

			(*LabelFeatures).train[i].features[0] = feature_X[i]
			(*LabelFeatures).train[i].features[1] = feature_Y[i]
			(*LabelFeatures).train[i].features[2] = feature_Z[i]
			
			for j := 0; j < len(ls); j++ {
				for s := 0; s < ls[j].size; s++ {
					(*LabelFeatures).train[i].label = ls[s].label
				}
			}
		}	
	}
}

