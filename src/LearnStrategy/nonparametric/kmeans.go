package nonparametric

import (
	"../"
)

type Kmeans struct {
}

/**
 * [Kmeans description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} DataLearner *DataLearner.DataLearner   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 */
func (theKeans *Kmeans) Learn(DataLearner *learnstrategy.DataLearner){

	auxlen,_ := (*DataLearner).Getlen(learnstrategy.Testflag)

	(*DataLearner).Centroid()

	(*DataLearner).CalcCenterdistance()
	for i := 0; i < auxlen; i++ {
		(*DataLearner).Sortdist(i,learnstrategy.Centerdistflag)
	}

	(*DataLearner).GetAccuracy()	
}