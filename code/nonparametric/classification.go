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
	
	auxlen,_ := (*dataset).Getlen(generalizecartesian.Trainflag)

	(*dataset).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,generalizecartesian.Knowflag)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

func Kmeans(dataset *generalizecartesian.Labelfeatures){

	auxlen,_ := (*dataset).Getlen(generalizecartesian.Trainflag)

	(*dataset).Centroid()

	(*dataset).CalcCenterdistance()
	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,generalizecartesian.Centerdistflag)
	}

	(*dataset).GetAccuracy()	
}
