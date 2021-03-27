package genericdata

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"../basicdata"
)

/**
 * [get the distance from each entry of the test group to the train group]
 * @struct {[type]} ds *DataSet) Calcdistance( 	[the data set]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) Calcdistance() error {

	var sum float64 = 0.0
	if len((*ds).test) == 0 {
		if len((*ds).train) == 0 {
			return errors.New("the test and train datasets weren't provided")
		} else {
			return errors.New("the test dataset weren't provided")
		}
	} else {
		if len((*ds).train) == 0 {
			return errors.New("the train dataset weren't provided")
		}
	}

	if len((*ds).result) == 0 {
		(*ds).Allocate(Resultflag, len((*ds).test), len((*ds).train))
	}

	(*ds).is_sortedbydist = make([]bool, len((*ds).test))

	for i := 0; i < len((*ds).test); i++ {

		(*ds).is_sortedbydist[i] = false

		for j := 0; j < len((*ds).train); j++ {
			sum = 0.0
			for f := 0; f < 3; f++ {
				sum += (math.Pow((*ds).test[i].Features[f]-(*ds).train[j].Features[f], 2))
			}
			(*ds).result[i].F_point[j].Dist = math.Sqrt(sum)
			(*ds).result[i].F_point[j].Distlabel = (*ds).train[j].Label
		}
	}

	return nil

}

/**
 * [get the distance from each entry of the test group to the centroid of each label of the train group]
 * @struct {[type]} ds *DataSet) CalcCenterdistance( 	[the data set]
 * @return {error} 													[gets errors]
 */
func (ds *DataSet) CalcCenterdistance() error {

	var currentdist float64

	if len((*ds).test) == 0 {
		if len((*ds).centroid) == 0 {
			return errors.New("the test dataset weren't provided and the centroid weren't computed")
		} else {
			return errors.New("the test dataset weren't provided")
		}
	} else {
		if len((*ds).centroid) == 0 {
			return errors.New("the centroid dataset weren't computed")
		}
	}

	if len((*ds).result) == 0 {
		(*ds).Allocate(Resultflag, len((*ds).test), len((*ds).centroid))
	}

	(*ds).is_sortedbycenter = make([]bool, len((*ds).test))

	for i := 0; i < len((*ds).test); i++ {

		(*ds).is_sortedbycenter[i] = false

		for j := 0; j < len((*ds).centroid); j++ {

			currentdist = math.Pow(((*ds).test[i].Features[0] - (*ds).centroid[j].Features[0]), 2)

			currentdist += math.Pow(((*ds).test[i].Features[1] - (*ds).centroid[j].Features[1]), 2)
			currentdist += math.Pow(((*ds).test[i].Features[2] - (*ds).centroid[j].Features[2]), 2)

			(*ds).result[i].F_point[j].Dist = currentdist
			(*ds).result[i].F_point[j].Distlabel = (*ds).centroid[j].Label
		}
	}

	return nil
}

/**
 * [add the interest group based on the k nearest neighbors]
 * @struct {[type]} ds *DataSet) AddInterest(t_size int ,k int [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) AddInterest(t_size int, k int) error {

	if len((*ds).interestgroup) == 0 {
		(*ds).Allocate(Interestflag, len((*ds).test), k)
	}

	if len((*ds).result) == 0 {
		return errors.New("result weren't computed")
	}

	for i := 0; i < t_size; i++ {
		for j := 0; j < k; j++ {
			(*ds).interestgroup[i].Interestlabel[j] = (*ds).result[i].F_point[j].Distlabel
			(*ds).interestgroup[i].Interestdist[j] = (*ds).result[i].F_point[j].Dist
		}
	}
	return nil
}

/**
 * [get the greatest ocorrence at the interest group]
 * @struct {[type]} ds *DataSet) GetGreatestOcorrence( [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) GetGreatestOcorrence(k int) error {

	if len((*ds).result) == 0 {
		return errors.New("result data set not provided")
	}

	ocorrence := make(map[string]int)

	for i := 0; i < len((*ds).test); i++ {
		(*ds).result[i].Greatestoccurrence = 0
		for j := 0; j < k; j++ {
			ocorrence[(*ds).interestgroup[i].Interestlabel[j]] = 0
		}
		for j := 0; j < k; j++ {
			ocorrence[(*ds).interestgroup[i].Interestlabel[j]]++
		}
		for j := 0; j < k; j++ {
			if (*ds).result[i].Greatestoccurrence < ocorrence[(*ds).interestgroup[i].Interestlabel[j]] {
				(*ds).result[i].Greatestoccurrence = ocorrence[(*ds).interestgroup[i].Interestlabel[j]]
				(*ds).result[i].Learnedlabel = (*ds).interestgroup[i].Interestlabel[j]
			}
		}
	}
	return nil
}

/**
 * [compute the centroid of each group]
 * @struct {[type]} ds *DataSet) Centroid( [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) Centroid() error {

	if len((*ds).train) == 0 {
		return errors.New("train dataset weren't provided")
	}

	if len((*ds).centroid) == 0 {
		(*ds).Allocate(Centroidflag, 1)
	}
	var sun [3]float64
	var allsun [3]float64
	var distgroupcentroid [3]float64
	var auxindex int

	allsun[0] = 0.0
	allsun[1] = 0.0
	allsun[2] = 0.0

	for i := 0; i < len((*ds).sizelabel); i++ {
		sun[0] = 0.0
		sun[1] = 0.0
		sun[2] = 0.0
		for j := 0; j < (*ds).sizelabel[i].Size_l; j++ {

			auxindex = j + (i * (*ds).sizelabel[i].Size_l)

			sun[0] += (*ds).train[auxindex].Features[0]
			sun[1] += (*ds).train[auxindex].Features[1]
			sun[2] += (*ds).train[auxindex].Features[2]

			allsun[0] += (*ds).train[auxindex].Features[0]
			allsun[1] += (*ds).train[auxindex].Features[1]
			allsun[2] += (*ds).train[auxindex].Features[2]
		}
		(*ds).centroid[i].Features[0] = (sun[0] / float64((*ds).sizelabel[i].Size_l))
		(*ds).centroid[i].Features[1] = (sun[1] / float64((*ds).sizelabel[i].Size_l))
		(*ds).centroid[i].Features[2] = (sun[2] / float64((*ds).sizelabel[i].Size_l))

		(*ds).centroid[i].Label = (*ds).sizelabel[i].Label

		(*ds).allcentroid.Features[0] = allsun[0] / float64(len((*ds).train))
		(*ds).allcentroid.Features[1] = allsun[1] / float64(len((*ds).train))
		(*ds).allcentroid.Features[2] = allsun[2] / float64(len((*ds).train))
	}
	for i := 0; i < len((*ds).sizelabel); i++ {
		(*ds).allcentroid.Maxradius = allsun[0] / float64(len((*ds).train))
	}

	distgroupcentroid[0] = (*ds).euclidiandistance((*ds).allcentroid.Features, (*ds).centroid[0].Features)
	distgroupcentroid[1] = (*ds).euclidiandistance((*ds).allcentroid.Features, (*ds).centroid[1].Features)
	distgroupcentroid[2] = (*ds).euclidiandistance((*ds).allcentroid.Features, (*ds).centroid[2].Features)

	(*ds).allcentroid.Minradius = distgroupcentroid[0]
	(*ds).allcentroid.Maxradius = distgroupcentroid[1]

	for i := 0; i < len((*ds).sizelabel); i++ {
		if (*ds).allcentroid.Minradius > distgroupcentroid[i] {
			(*ds).allcentroid.Minradius = distgroupcentroid[i]
		}
		if (*ds).allcentroid.Maxradius < distgroupcentroid[i] {
			(*ds).allcentroid.Maxradius = distgroupcentroid[i]
		}
	}
	fmt.Println((*ds).allcentroid)

	return nil
}

/**
 * [comput the distance between each group centroid]
 * @struct {[type]} ds *DataSet) GroupCenterdists( [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) GroupCenterdists() error {

	if len((*ds).centroid) == 0 {
		return errors.New("centroid weren't provided")
	}
	var sum float64 = 0
	var j = 0

	if len((*ds).centerdist) == 0 {
		(*ds).Allocate(Centerdistflag, 1)
	}

	for i := 0; i < 2; i++ {
		for j = i + 1; j < 3; j++ {
			sum = math.Pow(((*ds).centroid[i].Features[0] - (*ds).centroid[j].Features[0]), 2)
			sum += math.Pow(((*ds).centroid[i].Features[1] - (*ds).centroid[j].Features[1]), 2)
			sum += math.Pow(((*ds).centroid[i].Features[2] - (*ds).centroid[j].Features[2]), 2)

			(*ds).centerdist[i+j-1].Dist = math.Sqrt(sum)
			(*ds).centerdist[i+j-1].Distlabel = (*ds).centroid[i].Label + " to " + (*ds).centroid[j].Label
		}
	}
	return nil
}

func (ds *DataSet) Calcradius() error {
	if len((*ds).train) == 0 {
		if len((*ds).centroid) == 0 {
			return errors.New("train dataset and centroid weren't provided")
		} else {
			return errors.New("train dataset weren't provided")
		}
	} else {
		if len((*ds).centroid) == 0 {
			return errors.New("centroid weren't provided")
		}
	}

	var auxradius float64
	var auxindex int

	for i := 0; i < len((*ds).sizelabel); i++ {

		auxradius = 0.0

		(*ds).centroid[i].Radius = auxradius

		for j := 0; j < (*ds).sizelabel[i].Size_l; j++ {

			auxindex = j + (i * (*ds).sizelabel[i].Size_l)

			auxradius = (*ds).euclidiandistance((*ds).train[auxindex].Features, (*ds).centroid[i].Features)

			if (*ds).centroid[i].Radius < auxradius {
				(*ds).centroid[i].Radius = auxradius
			}
		}

	}

	return nil
}