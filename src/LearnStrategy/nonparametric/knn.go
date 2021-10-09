package nonparametric

import (
	"../"
)

type Knn struct {
}

/**
 * [KNN description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} DataLearner *DataLearner.DataLearner   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 * @param  {[type]} k             int                            [number of neighbors]
 */
func (thek *Knn) Learn(DataLearner *learnstrategy.DataLearner){
	
	k := 3

	auxlen,_ := (*DataLearner).Getlen(learnstrategy.Testflag)

	(*DataLearner).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*DataLearner).Sortdist(i,learnstrategy.Trainflag)

		(*DataLearner).AddInterest(i,k)
	}
	(*DataLearner).GetGreatestOcorrence(k)
}