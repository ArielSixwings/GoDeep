package nonparametric

import (
	"../"
)

type Kmeans struct {
}

/**
 * [Kmeans description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *dataset.DataSet   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 */
func (theKeans *Kmeans) Learn(dataset *learnstrategy.DataSet){

	auxlen,_ := (*dataset).Getlen(learnstrategy.Testflag)

	(*dataset).Centroid()

	(*dataset).CalcCenterdistance()
	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,learnstrategy.Centerdistflag)
	}

	(*dataset).GetAccuracy()	
}