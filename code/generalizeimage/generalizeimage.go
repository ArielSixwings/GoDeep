package generalizeimage

import(
	"math"
	"fmt"
)
type labeldist struct{

	dist []float64

	learnedlabel string

	greatestoccurrence int
}

type features struct {

	features 	[3]float64 
	
	label 		string
}


type Labelfeatures struct {

	train 		[]features

	know 		[]features
	
	result 	[]labeldist
}

/**
 * [func description]
 * @param  {[type]} lf *labelfeatures) calcdistance( [description]
 * @return {[type]}    [description]
 */

type ByDist []float64

func (d ByDist) Len() int { return len(d) }
func (d ByDist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d ByDist) Less(i, j int) bool { return d[i] < d[j] }

func (lf *Labelfeatures) calcdistance() { 

	var sum float64 = 0.0

	for i := 0; i < len((*lf).train); i++ {
		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).know[j].features[f] - (*lf).train[j].features[f],2))
			}			
			(*lf).result[i].dist[j] = math.Sqrt(sum)
			(*lf).result[i].learnedlabel = (*lf).know[j].label
		}
	}

}

func (lf Labelfeatures) Printfeatures(){
	
	fmt.Println("These are the know features")
	for i := 0; i < len(lf.know); i++ {
	
		fmt.Println(lf.know[i])	
	
	}
	
	fmt.Println("These are the train features")
	for i := 0; i < len(lf.train); i++ {
		fmt.Println(lf.train[i])
	}
}

type Sizelabel struct{
	Label 	string
	Size_l 	int
}

type Groupflag int

const (

	Knowflag Groupflag = 0

	Trainflag Groupflag = 1

)

func Generalize_for_nonparametric(LabelFeatures *Labelfeatures, feature_X []float64, feature_Y []float64, feature_Z []float64,ls []Sizelabel,group Groupflag,size int){

	if group == Knowflag{
		(*LabelFeatures).know = make([]features,size)
	} else{
		(*LabelFeatures).train = make([]features,size)
	}

	var j int = 0 
	for i := 0; i < size; i++ {
		if group == Knowflag{
			
			(*LabelFeatures).know[i].features[0] = feature_X[i]
			(*LabelFeatures).know[i].features[1] = feature_Y[i]
			(*LabelFeatures).know[i].features[2] = feature_Z[i]
			
			if i < (1+j)*ls[j].Size_l{
				(*LabelFeatures).know[i].label = ls[j].Label
			} else {
				j++
				(*LabelFeatures).know[i].label = ls[j].Label
			}

		} else{

			(*LabelFeatures).train[i].features[0] = feature_X[i]
			(*LabelFeatures).train[i].features[1] = feature_Y[i]
			(*LabelFeatures).train[i].features[2] = feature_Z[i]
			
			if i < (1+j)*ls[j].Size_l{
				(*LabelFeatures).train[i].label = ls[j].Label
			} else {
				j++
				(*LabelFeatures).train[i].label = ls[j].Label
			}
		}	
	}
}

