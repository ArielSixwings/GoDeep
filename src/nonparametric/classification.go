package nonparametric

import (
	"../genericdata"

)

/**
 * [KNN description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *dataset.DataSet   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 * @param  {[type]} k             int                            [number of neighbors]
 */
func KNN(dataset *genericdata.DataSet,k int){
	
	auxlen,_ := (*dataset).Getlen(genericdata.Testflag)

	(*dataset).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,genericdata.Trainflag)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

/**
 * [Kmeans description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *dataset.DataSet   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 */
func Kmeans(dataset *genericdata.DataSet){

	auxlen,_ := (*dataset).Getlen(genericdata.Testflag)

	(*dataset).Centroid()

	(*dataset).CalcCenterdistance()
	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,genericdata.Centerdistflag)
	}

	(*dataset).GetAccuracy()	
}
