package generalizeimage

import(
	"math"
	"fmt"
	"sort"
)
type Groupflag int

const (

	Knowflag Groupflag = 0

	Trainflag Groupflag = 1

	Resultflag Groupflag = 2

	GreatestOccurrence Groupflag = 3

	Labelflag Groupflag = 4

	Statusflag Groupflag = 5

)

type labeldist struct{

	dist 				[]float64
	learnedlabel 		string
	status 				string
	greatestoccurrence 	int
}

type features struct {

	features 	[3]float64 
	label 		string
}


type Labelfeatures struct {

	train 		[]features
	know 		[]features	
	result 		[]labeldist
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

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) Calcdistance( [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) Calcdistance() { 

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

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,sizedist ...int [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,sizedist ...int){
	switch allflag {
	case Resultflag:
		(*lf).result = make([]labeldist,allsize)
		for i := 0; i < allsize; i++ {
			if len(sizedist) > 0 {
				(*lf).result[i].dist = make([]float64,sizedist[0])
			}
		}	

	case Knowflag:
		(*lf).know = make([]features,allsize)

	case Trainflag:
		(*lf).train = make([]features,allsize)			
	}
}

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) SetResult(i int, l_label string, g_ocurrence int [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) SetResult(i int, l_label string, g_ocurrence int){
	(*lf).result[i].learnedlabel = l_label
	(*lf).result[i].greatestoccurrence = g_ocurrence
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetResultstring(i int, getflag Groupflag [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) GetResultstring(i int, getflag Groupflag) string{
	switch getflag {
	case Statusflag:
		return lf.result[i].learnedlabel

	case Labelflag:
		return lf.result[i].status			
	}
	return "default"
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetResultint(i int [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) GetResultint(i int) int{return lf.result[i].greatestoccurrence}
/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printfeatures( [description]
 * @return {[type]}    [description]
 */
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

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printresults( [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Printresults(){

	fmt.Println("These are the results")
	for i := 0; i < len(lf.result); i++ {
		fmt.Println(lf.know[i])	
	}
}
/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Getlen(lenflag Groupflag [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Getlen(lenflag Groupflag) int{ 
	switch lenflag {
	case Knowflag:
		return len(lf.know)

	case Trainflag:
		return len(lf.train)			
	}
	return 0

}

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) Sortdist(i int [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) Sortdist(i int){ 
	sort.Sort(ByDist((*lf).result[i].dist))
}

type Sizelabel struct{
	Label 	string
	Size_l 	int
}

/**
 * [Generalize_for_nonparametric description]
 * @param {[type]} LabelFeatures *Labelfeatures [description]
 * @param {[type]} feature_X     []float64      [description]
 * @param {[type]} feature_Y     []float64      [description]
 * @param {[type]} feature_Z     []float64      [description]
 * @param {[type]} ls            []Sizelabel    [description]
 * @param {[type]} group         Groupflag      [description]
 * @param {[type]} size          int            [description]
 */
func Generalize_for_nonparametric(lf *Labelfeatures, feature_X []float64, feature_Y []float64, feature_Z []float64,ls []Sizelabel,group Groupflag,size int){

	if group == Knowflag{
		(*lf).Allocate(Knowflag,size)
	} else{
		(*lf).Allocate(Trainflag,size)
	}

	var j int = 0 
	for i := 0; i < size; i++ {
		if group == Knowflag{
			
			(*lf).know[i].features[0] = feature_X[i]
			(*lf).know[i].features[1] = feature_Y[i]
			(*lf).know[i].features[2] = feature_Z[i]
			
			if i < (1+j)*ls[j].Size_l{
				(*lf).know[i].label = ls[j].Label
			} else {
				j++
				(*lf).know[i].label = ls[j].Label
			}

		} else{

			(*lf).train[i].features[0] = feature_X[i]
			(*lf).train[i].features[1] = feature_Y[i]
			(*lf).train[i].features[2] = feature_Z[i]
			
			if i < (1+j)*ls[j].Size_l{
				(*lf).train[i].label = ls[j].Label
			} else {
				j++
				(*lf).train[i].label = ls[j].Label
			}
		}	
	}
}

