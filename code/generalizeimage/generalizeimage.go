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

	Interestflag Groupflag = 6

)

type interest struct{

	interestlabel 		[]string
	interestocurrence 	[]float64
}

type labeldist struct{

	dist 				[]float64
	learnedlabel 		[]string
	status 				string
	greatestoccurrence 	int
}

type features struct {

	features 	[3]float64 
	label 		string
}


type Labelfeatures struct {

	train 			[]features
	know 			[]features	
	result 			[]labeldist
	interestgroup 	[]interest
	is_sorted 		[]bool
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
	
	(*lf).is_sorted = make([]bool,len((*lf).train))

	for i := 0; i < len((*lf).train); i++ {

		//(*lf).is_sorted[i] = false
		
		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).train[i].features[f] - (*lf).know[j].features[f],2))
			}			
			(*lf).result[i].dist[j] = math.Sqrt(sum)
			(*lf).result[i].learnedlabel[j] = (*lf).know[j].label
		}
	}

}

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,secondsize ...int [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,secondsize ...int){
	switch allflag {
	case Resultflag:
		(*lf).result = make([]labeldist,allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).result[i].dist = make([]float64,secondsize[0])
				(*lf).result[i].learnedlabel = make([]string,secondsize[0])
			}
		}
	case Knowflag:
		(*lf).know = make([]features,allsize)

	case Trainflag:
		(*lf).train = make([]features,allsize)
	case Interestflag:
		(*lf).interestgroup = make([]interest,allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).interestgroup[i].interestocurrence = make([]float64,secondsize[0])
				(*lf).interestgroup[i].interestlabel = make([]string,secondsize[0])
			}
		}					
	}
}

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) SetInterest(t_size int ,k int ,val float64 [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) SetInterest(t_size int ,k int ,val float64){
	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*lf).interestgroup[i].interestlabel[k] = "default"
			(*lf).interestgroup[i].interestocurrence[k] = val
		}
	}
}

/**
 * [func description]
 * @param  {[type]} lf *Labelfeatures) AddInterest(t_size int ,k int [description]
 * @return {[type]}    [description]
 */
func (lf *Labelfeatures) AddInterest(t_size int ,k int){
	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*lf).interestgroup[i].interestlabel[k] = (*lf).result[i].learnedlabel[k]
			(*lf).interestgroup[i].interestocurrence[k] = (*lf).result[i].dist[k]
		}
	}
}
/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetResultstring(i int, getflag Groupflag [description]
 * @return {[type]}    [description]
 */
// func (lf Labelfeatures) GetResultstring(i int, getflag Groupflag) string{
// 	switch getflag {
// 	case Statusflag:
// 		return lf.result[i].status

// 	case Labelflag:
// 		return lf.result[i].learnedlabel			
// 	}
// 	return "default"
// }

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetKnowstring(i int [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) GetKnowstring(i int) string{ return lf.know[i].label}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetTrainstring(i int [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) GetTrainstring(i int) string{ return lf.know[i].label}

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
		fmt.Println(lf.result[i].learnedlabel)	
	}
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printdists( [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Printdists(){
	fmt.Println("These are the results")
	for i := 0; i < len(lf.result); i++ {
		fmt.Println("results: ", i)
		fmt.Println(lf.result[i])	
	}	
}

func (lf Labelfeatures) Printinterest(){
	for i := 0; i < len(lf.interestgroup); i++ {
		fmt.Println(lf.interestgroup[i])
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
	if (*lf).is_sorted[i] {
		fmt.Println("the distance set of this images are already sorted")
		return
	} 
	sort.Sort(ByDist((*lf).result[i].dist))
	(*lf).is_sorted[i] = true
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
		fmt.Println("using Allocate at know with size:   ", size)
	} else{
		(*lf).Allocate(Trainflag,size)
		fmt.Println("using Allocate at train with size:   ", size)
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