package learnstrategy

import (
	"errors"
	"../basicdata"
)
func (ds *DataSet) SetLearnStrategy(ls learnStrategy) {
    ds.Strategy = ls
}
func (ds *DataSet) ProcessLearn(){
	ds.Strategy.Learn(ds)
}
/**
 * [use make build in fucntion to allocate setioncs of the DataSet based on the allocate flag]
 * @struct {[type]} ds *DataSet) Allocate(allflag Groupflag, allsize int,secondsize ...int [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) Allocate(allflag Groupflag, allsize int, secondsize ...int) error {
	if allsize == 0 {
		return errors.New("invalid size of length 0, can't allocate")
	} else {
		if allsize < 0 {
			return errors.New("invalid value of size, can't use negative value to allocate")
		}
	}
	switch allflag {
	case Resultflag:
		(*ds).result = make([]cartesian.Labeldist, allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*ds).result[i].F_point = make([]cartesian.Featurepoint, secondsize[0])
			} else {
				return errors.New("invalid sencondsize of length 0 or negative, can't allocate")
			}
		}
		return nil
	case Trainflag:
		(*ds).train = make([]cartesian.Features, allsize)
		return nil
	case Testflag:
		(*ds).test = make([]cartesian.Features, allsize)
		return nil
	case Interestflag:
		(*ds).interestgroup = make([]cartesian.Interest, allsize)
		for i := 0; i < allsize; i++ {
			if len(secondsize) > 0 {
				(*ds).interestgroup[i].Interestdist = make([]float64, secondsize[0])
				(*ds).interestgroup[i].Interestlabel = make([]string, secondsize[0])
			} else {
				return errors.New("invalid sencondsize of length 0 or negative, can't allocate")
			}
		}
		return nil
	case Centroidflag:
		(*ds).centroid = make([]cartesian.Centroidinfo, 3)
		return nil
		// case Allcentroidflag:
		// 	(*ds).allcentroid = make([]centroidinfo,3)
		//return nil
	case Centerdistflag:
		(*ds).centerdist = make([]cartesian.Featurepoint, 3)
		return nil
	default:
		return errors.New("invalid request of Allocate method, unkown allocate flag")
	}
}

/**
 * [get the label of the ith entry at the selected group]
 * @struct {[type]} ds DataSet) Gettrainstring(i int [description]
 * @return {string,error} 	 											[gets errors]
 */
func (ds DataSet) Getlabel(labelflag Groupflag, i int) (string, error) {
	switch Labelflag {
	case Trainflag:
		if len(ds.train) == 0 {
			return "invalid", errors.New("train dataset weren't provided")
		}
		return ds.train[i].Label, nil

	case Testflag:
		if len(ds.test) == 0 {
			return "invalid", errors.New("test dataset weren't provided")
		}
		return ds.test[i].Label, nil
	default:
		return "error", errors.New("invalid request of Getlabel method, unkown label flag")
	}
}

/**
 * [get the length of the section selected by the flag]
 * @struct {[type]} ds DataSet) Getlen(lenflag Groupflag [description]
 * @return {int,error} 	 											[gets errors]
 */
func (ds DataSet) Getlen(lenflag Groupflag) (int, error) {
	switch lenflag {
	case Trainflag:
		return len(ds.train), nil

	case Testflag:
		return len(ds.test), nil
	case Centroidflag:
		return len(ds.centroid), nil
	default:
		return 0, errors.New("invalid request of Getlen method, unkown length flag")
	}
}

/**
 * [Build description]
 * @struct {[type]} ds        *DataSet [description]
 * @param {[type]} feature_X []float64      [description]
 * @param {[type]} feature_Y []float64      [description]
 * @param {[type]} feature_Z []float64      [description]
 * @param {[type]} ls        []Sizelabel    [description]
 * @param {[type]} group     Groupflag      [description]
 * @param {[type]} size      int            [description]
 */
func (ds *DataSet) Build(feature_X []float64, feature_Y []float64, feature_Z []float64, ls []cartesian.Sizelabel, group Groupflag, size int) error {

	if size != len(feature_X) {
		return errors.New("Incompatible length between features and data set")
	}

	var j int = 0

	if group == Trainflag {
		(*ds).Allocate(Trainflag, size)
	} else {
		(*ds).Allocate(Testflag, size)
	}

	(*ds).sizelabel = ls
	for i := 0; i < size; i++ {
		if group == Trainflag {

			(*ds).train[i].Features[0] = feature_X[i]
			(*ds).train[i].Features[1] = feature_Y[i]
			(*ds).train[i].Features[2] = feature_Z[i]

			if i < (1+j)*ls[j].Size_l {
				(*ds).train[i].Label = ls[j].Label
			} else {
				j++
				(*ds).train[i].Label = ls[j].Label
			}

		} else {

			(*ds).test[i].Features[0] = feature_X[i]
			(*ds).test[i].Features[1] = feature_Y[i]
			(*ds).test[i].Features[2] = feature_Z[i]

			if i < (1+j)*ls[j].Size_l {
				(*ds).test[i].Label = ls[j].Label
			} else {
				j++
				(*ds).test[i].Label = ls[j].Label
			}
		}
	}
	return nil
}

/**
 * [check if the learned label provided by some external IA process is corret]
 * @struct {[type]}ds *DataSet) GetAccuracy( [description]
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) GetAccuracy() error {

	if len((*ds).result) == 0 {
		return errors.New("results weren't computed")
	}

	for i := 0; i < len((*ds).test); i++ {
		if (*ds).result[i].Learnedlabel == (*ds).test[i].Label {
			(*ds).result[i].Status = true
		} else {
			(*ds).result[i].Status = false
		}
	}

	return nil
}