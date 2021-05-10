package learnstrategy

import (
	"errors"
	"../basicdata"
	"fmt"
)
func (ds *DataLearner) SetLearnStrategy(ls learnStrategy) {
	ds.Strategy = ls
}
func (ds *DataLearner) ProcessLearn(){
	ds.Strategy.Learn(ds)
}

func (ds *DataLearner) Allocate(allflag Groupflag, allsize int, secondsize ...int) error {
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

func (ds DataLearner) Getlabel(labelflag Groupflag, i int) (string, error) {
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

func (ds DataLearner) Getlen(lenflag Groupflag) (int, error) {
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

func (ds *DataLearner) Build(features *[]cartesian.Features,ri cartesian.ReadInformation,groupsize int) error {
	var j int = 0
	var k int = 0
	var proportion int
	proportion = ri.SizeData/groupsize
	for i := 0; i < ri.SizeData; i++ {
		if i%2 == 0 {
			(*ds).train = append((*ds).train,(*features)[i])
		} else {
			(*ds).test = append((*ds).test,(*features)[i])	
		}
	}

	fmt.Println(k)
	fmt.Println((*features)[k].Label)
	for i := 0; i < groupsize; i++ {
		if len(ri.Labelsize) > 1{
			if i < (1+j)*ri.Labelsize[j].Size_l/proportion {
				(*ds).train[i].Label = ri.Labelsize[j].Label
				(*ds).test[i].Label = ri.Labelsize[j].Label
		} else {
			j++
				(*ds).train[i].Label = ri.Labelsize[j].Label
				(*ds).test[i].Label = ri.Labelsize[j].Label
		}

		} else {
			if i < groupsize/2 {
				(*ds).train[i].Label = (*features)[k].Label
				(*ds).test[i].Label = (*features)[k].Label
		} 	else {
				k = len(*features) - 1
				(*ds).train[i].Label = (*features)[k].Label
				(*ds).test[i].Label = (*features)[k].Label
				fmt.Println((*ds).train[i].Label)
			}		
		}
	}
	return nil
}

func (ds *DataLearner) GetAccuracy() error {

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