package generalizecartesian

import(
	"math"
	//"fmt"
	"sort"
	"errors"
)

func (d ByDist) Len() int { return len(d) }
func (d ByDist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d ByDist) Less(i, j int) bool { return d[i].dist < d[j].dist }

/**
 * [for the ith entry, selected by the i parameter, sort all the distances between the entry and the know group]
 * @param  {*Labelfeatures} lf *Labelfeatures) Sortdist(i int) 	[the data set]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Sortdist(i int, sortflag Groupflag) error{
	
	switch sortflag {
	case Centerdistflag:
		if (*lf).is_sortedbydist[i] {
			return errors.New("the distance set of this dataset are already sorted by the distance to the group center")
		}		
		sort.Sort(ByDist((*lf).result[i].f_point))

		(*lf).result[i].learnedlabel = (*lf).result[i].f_point[0].distlabel

		(*lf).is_sortedbycenter[i] = true
	case Knowflag:
		if (*lf).is_sortedbydist[i] {
			return errors.New("the distance set of this dataset are already sorted")
		}
		sort.Sort(ByDist((*lf).result[i].f_point))
	
		(*lf).is_sortedbydist[i] = true
	}
	return nil
}

/**
 * [get the distance from each entry of the train group to the know group]
 * @param  {*Labelfeatures} lf *Labelfeatures) Calcdistance( 	[the data set]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Calcdistance() error{ 

	var sum float64 = 0.0
	if len((*lf).train) == 0{
		if len((*lf).know) == 0{
			return errors.New("the train and know datasets weren't provided")
		} else{
			return errors.New("the train dataset weren't provided")
		}
	}else{
		if len((*lf).know) == 0{
			return errors.New("the know dataset weren't provided")
		}
	}

	if len((*lf).result) == 0{
		(*lf).Allocate(Resultflag,len((*lf).train),len((*lf).know))
	}
	
	(*lf).is_sortedbydist = make([]bool,len((*lf).train))

	for i := 0; i < len((*lf).train); i++ {

		(*lf).is_sortedbydist[i] = false
		
		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).train[i].features[f] - (*lf).know[j].features[f],2))
			}			
			(*lf).result[i].f_point[j].dist = math.Sqrt(sum)
			(*lf).result[i].f_point[j].distlabel = (*lf).know[j].label
		} 
	}

	return nil

}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) CalcCenterdistance( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) CalcCenterdistance() error{

	var currentdist float64

	if len((*lf).train) == 0{
		if len((*lf).centroid) == 0{
			return errors.New("the train and centroid datasets weren't provided")
		} else{
			return errors.New("the train dataset weren't provided")
		}
	}else{
		if len((*lf).centroid) == 0{
			return errors.New("the centroid dataset weren't provided")
		}
	}

	if len((*lf).result) == 0{
		(*lf).Allocate(Resultflag,len((*lf).train),len((*lf).centroid))
	}

	(*lf).is_sortedbycenter = make([]bool,len((*lf).train))

	for i := 0; i < len((*lf).train); i++ {
		
		(*lf).is_sortedbycenter[i] = false
		
		for j := 0; j < len((*lf).centroid); j++ {

			currentdist = math.Pow(((*lf).train[i].features[0] - (*lf).centroid[j].features[0]),2)
			
			currentdist += math.Pow(((*lf).train[i].features[1] - (*lf).centroid[j].features[1]),2)
			currentdist += math.Pow(((*lf).train[i].features[2] - (*lf).centroid[j].features[2]),2)
			
			(*lf).result[i].f_point[j].dist = currentdist
			(*lf).result[i].f_point[j].distlabel =  (*lf).centroid[j].label
			// fmt.Println(i,"  (*lf).result[i].f_point[j].distlabel : ",(*lf).result[i].f_point[j].distlabel )
		}
	}

	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,secondsize ...int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,secondsize ...int) error{
	switch allflag {
	case Resultflag:
		(*lf).result = make([]labeldist,allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).result[i].f_point = make([]featurepoint,secondsize[0])
			}
		}
		return nil
	case Knowflag:
		(*lf).know = make([]features,allsize)
		return nil
	case Trainflag:
		(*lf).train = make([]features,allsize)
		return nil
	case Interestflag:
		(*lf).interestgroup = make([]interest,allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).interestgroup[i].interestdist = make([]float64,secondsize[0])
				(*lf).interestgroup[i].interestlabel = make([]string,secondsize[0])
			}
		}
		return nil
	case Centroidflag:
		(*lf).centroid = make([]features,3)
		return nil
	case Centerdistflag:
		(*lf).centerdist = make([]featurepoint,3)
		return nil
	default:
		return errors.New("invalid request of Allocate method, unkown allocate flag")
	}
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) SetInterest(t_size int ,k int ,val float64 [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) SetInterest(t_size int ,k int ,val float64) error{
	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*lf).interestgroup[i].interestlabel[j] = "default"
			(*lf).interestgroup[i].interestdist[j] = val
		}
	}
	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) AddInterest(t_size int ,k int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) AddInterest(t_size int ,k int) error{
	
	if len((*lf).interestgroup) == 0{
		return errors.New("interestgroup not allocated")
	}

	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*lf).interestgroup[i].interestlabel[j] = (*lf).result[i].f_point[j].distlabel
			(*lf).interestgroup[i].interestdist[j] = (*lf).result[i].f_point[j].dist
		}
	}
	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) GetGreatestOcorrence( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GetGreatestOcorrence(k int) error{
	
	if len((*lf).result) == 0{
		return errors.New("result data set not provided")
	}

	ocorrence := make(map[string]int)

	for i := 0; i < len((*lf).train); i++ {
		(*lf).result[i].greatestoccurrence = 0
		for j := 0; j < k; j++ {
			ocorrence[(*lf).interestgroup[i].interestlabel[j]] = 0
		}
		for j := 0; j < k; j++ {
			ocorrence[(*lf).interestgroup[i].interestlabel[j]]++
		}
		for j := 0; j < k; j++ {
			if (*lf).result[i].greatestoccurrence < ocorrence[(*lf).interestgroup[i].interestlabel[j]]{
				(*lf).result[i].greatestoccurrence = ocorrence[(*lf).interestgroup[i].interestlabel[j]]
				(*lf).result[i].learnedlabel = (*lf).interestgroup[i].interestlabel[j]
			}
		}		
	}
	return nil
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetKnowstring(i int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf Labelfeatures) GetKnowstring(i int) (string,error){ 

	if len(lf.know) == 0{
		return "invalid", errors.New("know dataset weren't provided")
	}
	return lf.know[i].label,nil
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetTrainstring(i int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf Labelfeatures) GetTrainstring(i int) (string,error){ 
	if len(lf.train ) == 0 {
		return "invalid", errors.New("train dataset weren't provided")
	}
	return lf.train[i].label,nil
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) GetResultint(i int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf Labelfeatures) GetGreatestOccurrence(i int) (int,error){
	if len(lf.result) == 0 {
		return 0, errors.New("results weren't provided")
	}	
	return lf.result[i].greatestoccurrence,nil
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Getlen(lenflag Groupflag [description]
 * @return {error} 	 											[gets errors]
 */
func (lf Labelfeatures) Getlen(lenflag Groupflag) (int,error){ 
	switch lenflag {
	case Knowflag:
		return len(lf.know),nil

	case Trainflag:
		return len(lf.train),nil
	case Centroidflag:
		return len(lf.centroid),nil
	default:
		return 0,errors.New("invalid request of Getlen method, unkown length flag")
	}	
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
func Generalize_for_nonparametric(lf *Labelfeatures, feature_X []float64, feature_Y []float64, feature_Z []float64,ls []Sizelabel,group Groupflag,size int) error{

	if size != len(feature_X) {
		return errors.New("Incompatible length between features and data set")
	}

	var j int = 0 

	if group == Knowflag{
		(*lf).Allocate(Knowflag,size)
	} else{
		(*lf).Allocate(Trainflag,size)
	}

	(*lf).sizelabel = ls
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
	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) Centroid( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Centroid() error{

	if len((*lf).know) == 0{
		return errors.New("know dataset weren't provided")
	}

	if len((*lf).centroid) == 0{
		(*lf).Allocate(Centroidflag,1)
	}
	var sun [3]float64
	var auxindex int

	for i := 0; i < len((*lf).sizelabel); i++ {
		sun[0] = 0.0
		sun[1] = 0.0
		sun[2] = 0.0
		for j := 0; j < (*lf).sizelabel[i].Size_l; j++ {
			
			auxindex = j+(i*(*lf).sizelabel[i].Size_l)
			
			sun[0]+= (*lf).know[auxindex].features[0]
			sun[1]+= (*lf).know[auxindex].features[1]
			sun[2]+= (*lf).know[auxindex].features[2]
		}
		(*lf).centroid[i].features[0] = (sun[0]/float64((*lf).sizelabel[i].Size_l))
		(*lf).centroid[i].features[1] = (sun[1]/float64((*lf).sizelabel[i].Size_l))
		(*lf).centroid[i].features[2] = (sun[2]/float64((*lf).sizelabel[i].Size_l))

		(*lf).centroid[i].label = (*lf).sizelabel[i].Label
	}

	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) GroupCenterdists( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GroupCenterdists() error{
	
	if len((*lf).centroid) == 0 {
		return errors.New("centroid weren't provided")
	}
	var sum float64 = 0
	var j = 0

	if len((*lf).centerdist) == 0{
		(*lf).Allocate(Centerdistflag,1)
	}
	
	for i := 0; i < 2; i++ {
		for j = i+1; j < 3; j++ {
			sum = math.Pow(((*lf).centroid[i].features[0] - (*lf).centroid[j].features[0]),2)
			sum += math.Pow(((*lf).centroid[i].features[1] - (*lf).centroid[j].features[1]),2)
			sum += math.Pow(((*lf).centroid[i].features[2] - (*lf).centroid[j].features[2]),2)		
		
			(*lf).centerdist[i+j-1].dist = math.Sqrt(sum)
			(*lf).centerdist[i+j-1].distlabel = (*lf).centroid[i].label + " to " + (*lf).centroid[j].label
		}
	}
	return nil
}

/**
 * [func description]
 * @param  {*Labelfeatures} lf *Labelfeatures) GetAccuracy( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GetAccuracy() error{

	if len((*lf).result) == 0{
		return errors.New("results weren't computed")
	}

	for i := 0; i < len((*lf).train); i++ {
		if (*lf).result[i].learnedlabel == (*lf).train[i].label{
			(*lf).result[i].status = true
		} else {
			(*lf).result[i].status = false
		}
	}

	return nil
}