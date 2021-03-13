package nonparametric

import (
	"../generalizecartesian"

)

/**
 * [KNN description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *generalizecartesian.Labelfeatures   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 * @param  {[type]} k             int                            [number of neighbors]
 */
func KNN(dataset *generalizecartesian.Labelfeatures,k int){
	
	auxlen,_ := (*dataset).Getlen(generalizecartesian.Trainflag)

	(*dataset).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,generalizecartesian.Knowflag)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

/**
 * [Kmeans description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *generalizecartesian.Labelfeatures   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 */
func Kmeans(dataset *generalizecartesian.Labelfeatures){

	auxlen,_ := (*dataset).Getlen(generalizecartesian.Trainflag)

	(*dataset).Centroid()

	(*dataset).CalcCenterdistance()
	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,generalizecartesian.Centerdistflag)
	}

	(*dataset).GetAccuracy()	
}
