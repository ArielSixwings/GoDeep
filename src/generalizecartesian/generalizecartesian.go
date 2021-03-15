package generalizecartesian

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"../basicdata"
)

/**
 * [return lenght of the dist of the ByDist sort template]
 * @struct {[type]}   (d ByDist)  [description]
 * @return {[type]}   int         [description]
 */
func (d ByDist) Len() int { return len(d) }

/**
 * [func description]
 * @struct  {[type]}    d ByDist
 * @param i {[type]} int
 * @param j {[type]} int [description]
 * @return  {[type]}   [description]
 */
func (d ByDist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

/**
 * [return the samlest distance betwee the ith and jth entry]
 * @struct {[type]} d ByDist)       Less(i, j int [description]
 * @return {[type]}   [description]
 */
func (d ByDist) Less(i, j int) bool { return d[i].Dist < d[j].Dist }

/**
 * [for the ith entry, selected by the i parameter, sort all the distances between the entry and the know group]
 * @struct {[type]} lf *Labelfeatures)					[the data set]
 * @param  Sortdist(i int)
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Sortdist(i int, sortflag Groupflag) error {

	switch sortflag {
	case Centerdistflag:

		if len((*lf).is_sortedbydist) != 0 {
			if (*lf).is_sortedbydist[i] {
				return errors.New("the distance set of this dataset are already sorted by the distance to the group center")
			} else {
				if len((*lf).result) == 0 {
					return errors.New("result weren't computed")
				}
			}
		}
		sort.Sort(ByDist((*lf).result[i].F_point))
		(*lf).result[i].Learnedlabel = (*lf).result[i].F_point[0].Distlabel

		(*lf).is_sortedbycenter[i] = true
	case Knowflag:
		if (*lf).is_sortedbydist[i] {
			return errors.New("the distance set of this dataset are already sorted")
		} else {
			if len((*lf).result) == 0 {
				return errors.New("result weren't computed")
			}
		}
		sort.Sort(ByDist((*lf).result[i].F_point))

		(*lf).is_sortedbydist[i] = true
	}
	return nil
}

/**
 * [get the distance from each entry of the train group to the know group]
 * @struct {[type]} lf *Labelfeatures) Calcdistance( 	[the data set]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Calcdistance() error {

	var sum float64 = 0.0
	if len((*lf).train) == 0 {
		if len((*lf).know) == 0 {
			return errors.New("the train and know datasets weren't provided")
		} else {
			return errors.New("the train dataset weren't provided")
		}
	} else {
		if len((*lf).know) == 0 {
			return errors.New("the know dataset weren't provided")
		}
	}

	if len((*lf).result) == 0 {
		(*lf).Allocate(Resultflag, len((*lf).train), len((*lf).know))
	}

	(*lf).is_sortedbydist = make([]bool, len((*lf).train))

	for i := 0; i < len((*lf).train); i++ {

		(*lf).is_sortedbydist[i] = false

		for j := 0; j < len((*lf).know); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*lf).train[i].Features[f]-(*lf).know[j].Features[f], 2))
			}
			(*lf).result[i].F_point[j].Dist = math.Sqrt(sum)
			(*lf).result[i].F_point[j].Distlabel = (*lf).know[j].Label
		}
	}

	return nil

}

/**
 * [get the distance from each entry of the train group to the centroid of each label of the know group]
 * @struct {[type]} lf *Labelfeatures) CalcCenterdistance( 	[the data set]
 * @return {error} 													[gets errors]
 */
func (lf *Labelfeatures) CalcCenterdistance() error {

	var currentdist float64

	if len((*lf).train) == 0 {
		if len((*lf).centroid) == 0 {
			return errors.New("the train dataset weren't provided and the centroid weren't computed")
		} else {
			return errors.New("the train dataset weren't provided")
		}
	} else {
		if len((*lf).centroid) == 0 {
			return errors.New("the centroid dataset weren't computed")
		}
	}

	if len((*lf).result) == 0 {
		(*lf).Allocate(Resultflag, len((*lf).train), len((*lf).centroid))
	}

	(*lf).is_sortedbycenter = make([]bool, len((*lf).train))

	for i := 0; i < len((*lf).train); i++ {

		(*lf).is_sortedbycenter[i] = false

		for j := 0; j < len((*lf).centroid); j++ {

			currentdist = math.Pow(((*lf).train[i].Features[0] - (*lf).centroid[j].Features[0]), 2)

			currentdist += math.Pow(((*lf).train[i].Features[1] - (*lf).centroid[j].Features[1]), 2)
			currentdist += math.Pow(((*lf).train[i].Features[2] - (*lf).centroid[j].Features[2]), 2)

			(*lf).result[i].F_point[j].Dist = currentdist
			(*lf).result[i].F_point[j].Distlabel = (*lf).centroid[j].Label
		}
	}

	return nil
}

/**
 * [use make build in fucntion to allocate setioncs of the Labelfeatures based on the allocate flag]
 * @struct {[type]} lf *Labelfeatures) Allocate(allflag Groupflag, allsize int,secondsize ...int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Allocate(allflag Groupflag, allsize int, secondsize ...int) error {
	if allsize == 0 {
		return errors.New("invalid size of length 0, can't allocate")
	} else {
		if allsize < 0 {
			return errors.New("invalid value of size, can't use negative value to allocate")
		}
	}
	switch allflag {
	case Resultflag:
		(*lf).result = make([]cartesian.Labeldist, allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).result[i].F_point = make([]cartesian.Featurepoint, secondsize[0])
			} else {
				return errors.New("invalid sencondsize of length 0 or negative, can't allocate")
			}
		}
		return nil
	case Knowflag:
		(*lf).know = make([]cartesian.Features, allsize)
		return nil
	case Trainflag:
		(*lf).train = make([]cartesian.Features, allsize)
		return nil
	case Interestflag:
		(*lf).interestgroup = make([]cartesian.Interest, allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*lf).interestgroup[i].Interestdist = make([]float64, secondsize[0])
				(*lf).interestgroup[i].Interestlabel = make([]string, secondsize[0])
			} else {
				return errors.New("invalid sencondsize of length 0 or negative, can't allocate")
			}
		}
		return nil
	case Centroidflag:
		(*lf).centroid = make([]cartesian.Centroidinfo, 3)
		return nil
		// case Allcentroidflag:
		// 	(*lf).allcentroid = make([]centroidinfo,3)
		return nil
	case Centerdistflag:
		(*lf).centerdist = make([]cartesian.Featurepoint, 3)
		return nil
	default:
		return errors.New("invalid request of Allocate method, unkown allocate flag")
	}
}

/**
 * [add the interest group based on the k nearest neighbors]
 * @struct {[type]} lf *Labelfeatures) AddInterest(t_size int ,k int [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) AddInterest(t_size int, k int) error {

	if len((*lf).interestgroup) == 0 {
		(*lf).Allocate(Interestflag, len((*lf).train), k)
	}

	if len((*lf).result) == 0 {
		return errors.New("result weren't computed")
	}

	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*lf).interestgroup[i].Interestlabel[j] = (*lf).result[i].F_point[j].Distlabel
			(*lf).interestgroup[i].Interestdist[j] = (*lf).result[i].F_point[j].Dist
		}
	}
	return nil
}

/**
 * [get the greatest ocorrence at the interest group]
 * @struct {[type]} lf *Labelfeatures) GetGreatestOcorrence( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GetGreatestOcorrence(k int) error {

	if len((*lf).result) == 0 {
		return errors.New("result data set not provided")
	}

	ocorrence := make(map[string]int)

	for i := 0; i < len((*lf).train); i++ {
		(*lf).result[i].Greatestoccurrence = 0
		for j := 0; j < k; j++ {
			ocorrence[(*lf).interestgroup[i].Interestlabel[j]] = 0
		}
		for j := 0; j < k; j++ {
			ocorrence[(*lf).interestgroup[i].Interestlabel[j]]++
		}
		for j := 0; j < k; j++ {
			if (*lf).result[i].Greatestoccurrence < ocorrence[(*lf).interestgroup[i].Interestlabel[j]] {
				(*lf).result[i].Greatestoccurrence = ocorrence[(*lf).interestgroup[i].Interestlabel[j]]
				(*lf).result[i].Learnedlabel = (*lf).interestgroup[i].Interestlabel[j]
			}
		}
	}
	return nil
}

/**
 * [get the label of the ith entry at the selected group]
 * @struct {[type]} lf Labelfeatures) GetKnowstring(i int [description]
 * @return {string,error} 	 											[gets errors]
 */
func (lf Labelfeatures) Getlabel(labelflag Groupflag, i int) (string, error) {
	switch Labelflag {
	case Knowflag:
		if len(lf.know) == 0 {
			return "invalid", errors.New("know dataset weren't provided")
		}
		return lf.know[i].Label, nil

	case Trainflag:
		if len(lf.train) == 0 {
			return "invalid", errors.New("train dataset weren't provided")
		}
		return lf.train[i].Label, nil
	default:
		return "error", errors.New("invalid request of Getlabel method, unkown label flag")
	}
}

/**
 * [get the length of the section selected by the flag]
 * @struct {[type]} lf Labelfeatures) Getlen(lenflag Groupflag [description]
 * @return {int,error} 	 											[gets errors]
 */
func (lf Labelfeatures) Getlen(lenflag Groupflag) (int, error) {
	switch lenflag {
	case Knowflag:
		return len(lf.know), nil

	case Trainflag:
		return len(lf.train), nil
	case Centroidflag:
		return len(lf.centroid), nil
	default:
		return 0, errors.New("invalid request of Getlen method, unkown length flag")
	}
}

/**
 * [Generalize_for_nonparametric description]
 * @struct {[type]} lf        *Labelfeatures [description]
 * @param {[type]} feature_X []float64      [description]
 * @param {[type]} feature_Y []float64      [description]
 * @param {[type]} feature_Z []float64      [description]
 * @param {[type]} ls        []Sizelabel    [description]
 * @param {[type]} group     Groupflag      [description]
 * @param {[type]} size      int            [description]
 */
func Generalize_for_nonparametric(lf *Labelfeatures, feature_X []float64, feature_Y []float64, feature_Z []float64, ls []cartesian.Sizelabel, group Groupflag, size int) error {

	if size != len(feature_X) {
		return errors.New("Incompatible length between features and data set")
	}

	var j int = 0

	if group == Knowflag {
		(*lf).Allocate(Knowflag, size)
	} else {
		(*lf).Allocate(Trainflag, size)
	}

	(*lf).sizelabel = ls
	for i := 0; i < size; i++ {
		if group == Knowflag {

			(*lf).know[i].Features[0] = feature_X[i]
			(*lf).know[i].Features[1] = feature_Y[i]
			(*lf).know[i].Features[2] = feature_Z[i]

			if i < (1+j)*ls[j].Size_l {
				(*lf).know[i].Label = ls[j].Label
			} else {
				j++
				(*lf).know[i].Label = ls[j].Label
			}

		} else {

			(*lf).train[i].Features[0] = feature_X[i]
			(*lf).train[i].Features[1] = feature_Y[i]
			(*lf).train[i].Features[2] = feature_Z[i]

			if i < (1+j)*ls[j].Size_l {
				(*lf).train[i].Label = ls[j].Label
			} else {
				j++
				(*lf).train[i].Label = ls[j].Label
			}
		}
	}
	return nil
}

/**
 * [compute the centroid of each group]
 * @struct {[type]} lf *Labelfeatures) Centroid( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) Centroid() error {

	if len((*lf).know) == 0 {
		return errors.New("know dataset weren't provided")
	}

	if len((*lf).centroid) == 0 {
		(*lf).Allocate(Centroidflag, 1)
	}
	var sun [3]float64
	var allsun [3]float64
	var distgroupcentroid [3]float64
	var auxindex int

	allsun[0] = 0.0
	allsun[1] = 0.0
	allsun[2] = 0.0

	for i := 0; i < len((*lf).sizelabel); i++ {
		sun[0] = 0.0
		sun[1] = 0.0
		sun[2] = 0.0
		for j := 0; j < (*lf).sizelabel[i].Size_l; j++ {

			auxindex = j + (i * (*lf).sizelabel[i].Size_l)

			sun[0] += (*lf).know[auxindex].Features[0]
			sun[1] += (*lf).know[auxindex].Features[1]
			sun[2] += (*lf).know[auxindex].Features[2]

			allsun[0] += (*lf).know[auxindex].Features[0]
			allsun[1] += (*lf).know[auxindex].Features[1]
			allsun[2] += (*lf).know[auxindex].Features[2]
		}
		(*lf).centroid[i].Features[0] = (sun[0] / float64((*lf).sizelabel[i].Size_l))
		(*lf).centroid[i].Features[1] = (sun[1] / float64((*lf).sizelabel[i].Size_l))
		(*lf).centroid[i].Features[2] = (sun[2] / float64((*lf).sizelabel[i].Size_l))

		(*lf).centroid[i].Label = (*lf).sizelabel[i].Label

		(*lf).allcentroid.Features[0] = allsun[0] / float64(len((*lf).know))
		(*lf).allcentroid.Features[1] = allsun[1] / float64(len((*lf).know))
		(*lf).allcentroid.Features[2] = allsun[2] / float64(len((*lf).know))
	}
	for i := 0; i < len((*lf).sizelabel); i++ {
		(*lf).allcentroid.Maxradius = allsun[0] / float64(len((*lf).know))
	}

	distgroupcentroid[0] = (*lf).euclidiandistance((*lf).allcentroid.Features, (*lf).centroid[0].Features)
	distgroupcentroid[1] = (*lf).euclidiandistance((*lf).allcentroid.Features, (*lf).centroid[1].Features)
	distgroupcentroid[2] = (*lf).euclidiandistance((*lf).allcentroid.Features, (*lf).centroid[2].Features)

	(*lf).allcentroid.Minradius = distgroupcentroid[0]
	(*lf).allcentroid.Maxradius = distgroupcentroid[1]

	for i := 0; i < len((*lf).sizelabel); i++ {
		if (*lf).allcentroid.Minradius > distgroupcentroid[i] {
			(*lf).allcentroid.Minradius = distgroupcentroid[i]
		}
		if (*lf).allcentroid.Maxradius < distgroupcentroid[i] {
			(*lf).allcentroid.Maxradius = distgroupcentroid[i]
		}
	}
	fmt.Println((*lf).allcentroid)

	return nil
}

/**
 * [comput the distance between each group centroid]
 * @struct {[type]} lf *Labelfeatures) GroupCenterdists( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GroupCenterdists() error {

	if len((*lf).centroid) == 0 {
		return errors.New("centroid weren't provided")
	}
	var sum float64 = 0
	var j = 0

	if len((*lf).centerdist) == 0 {
		(*lf).Allocate(Centerdistflag, 1)
	}

	for i := 0; i < 2; i++ {
		for j = i + 1; j < 3; j++ {
			sum = math.Pow(((*lf).centroid[i].Features[0] - (*lf).centroid[j].Features[0]), 2)
			sum += math.Pow(((*lf).centroid[i].Features[1] - (*lf).centroid[j].Features[1]), 2)
			sum += math.Pow(((*lf).centroid[i].Features[2] - (*lf).centroid[j].Features[2]), 2)

			(*lf).centerdist[i+j-1].Dist = math.Sqrt(sum)
			(*lf).centerdist[i+j-1].Distlabel = (*lf).centroid[i].Label + " to " + (*lf).centroid[j].Label
		}
	}
	return nil
}

/**
 * [check if the learned label provided by some external IA process is corret]
 * @struct {[type]}lf *Labelfeatures) GetAccuracy( [description]
 * @return {error} 	 											[gets errors]
 */
func (lf *Labelfeatures) GetAccuracy() error {

	if len((*lf).result) == 0 {
		return errors.New("results weren't computed")
	}

	for i := 0; i < len((*lf).train); i++ {
		if (*lf).result[i].Learnedlabel == (*lf).train[i].Label {
			(*lf).result[i].Status = true
		} else {
			(*lf).result[i].Status = false
		}
	}

	return nil
}

func (lf *Labelfeatures) Calcradius() error {
	if len((*lf).know) == 0 {
		if len((*lf).centroid) == 0 {
			return errors.New("know dataset and centroid weren't provided")
		} else {
			return errors.New("know dataset weren't provided")
		}
	} else {
		if len((*lf).centroid) == 0 {
			return errors.New("centroid weren't provided")
		}
	}

	var auxradius float64
	var auxindex int

	for i := 0; i < len((*lf).sizelabel); i++ {

		auxradius = 0.0

		(*lf).centroid[i].Radius = auxradius

		for j := 0; j < (*lf).sizelabel[i].Size_l; j++ {

			auxindex = j + (i * (*lf).sizelabel[i].Size_l)

			auxradius = (*lf).euclidiandistance((*lf).know[auxindex].Features, (*lf).centroid[i].Features)

			if (*lf).centroid[i].Radius < auxradius {
				(*lf).centroid[i].Radius = auxradius
			}
		}

	}

	return nil
}
