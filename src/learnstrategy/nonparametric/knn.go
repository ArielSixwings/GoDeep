package nonparametric

import (
	"../"

)

type Knn struct {
}

/**
 * [KNN description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *dataset.DataSet   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 * @param  {[type]} k             int                            [number of neighbors]
 */
func (thek *Knn) Learn(dataset *learnstrategy.DataSet){
	
	k := 3

	auxlen,_ := (*dataset).Getlen(learnstrategy.Testflag)

	(*dataset).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,learnstrategy.Trainflag)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}