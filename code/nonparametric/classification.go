package nonparametric

import (
	"../generalizeimage"

)

/**
 * [KNN description]
 * @param {[type]} dataset *generalizeimage.Labelfeatures [description]
 * @param {[type]} k             int                            [description]
 */
func KNN(dataset *generalizeimage.Labelfeatures,k int){

	(*dataset).Allocate(generalizeimage.Resultflag,(*dataset).Getlen(generalizeimage.Trainflag),(*dataset).Getlen(generalizeimage.Knowflag))

	(*dataset).Allocate(generalizeimage.Interestflag,(*dataset).Getlen(generalizeimage.Trainflag),k)

	(*dataset).Calcdistance()

	for i := 0; i < (*dataset).Getlen(generalizeimage.Trainflag); i++ {
		(*dataset).Sortdist(i)

		(*dataset).SetInterest(i,k,0)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

