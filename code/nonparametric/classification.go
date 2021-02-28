package nonparametric

import (
	"../generalizecartesian"

)

/**
 * [KNN description]
 * @param {[type]} dataset *generalizecartesian.Labelfeatures [description]
 * @param {[type]} k             int                            [description]
 */
func KNN(dataset *generalizecartesian.Labelfeatures,k int){

	(*dataset).Allocate(generalizecartesian.Interestflag,(*dataset).Getlen(generalizecartesian.Trainflag),k)

	(*dataset).Calcdistance()

	for i := 0; i < (*dataset).Getlen(generalizecartesian.Trainflag); i++ {
		(*dataset).Sortdist(i)

		(*dataset).SetInterest(i,k,0)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

func Kmeans(dataset *generalizecartesian.Labelfeatures){

	(*dataset).Centroid()

	(*dataset).CalcCenterdistance()
	for i := 0; i < (*dataset).Getlen(generalizecartesian.Trainflag); i++ {
		(*dataset).SortCenterdist(i)
	}

	(*dataset).GetAccuracy()	
}
