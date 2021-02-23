package nonparametric

import (
	//"fmt"
	"../generalizeimage"

)

/**
 * [KNN description]
 * @param {[type]} LabelFeatures *generalizeimage.Labelfeatures [description]
 * @param {[type]} k             int                            [description]
 */
func KNN(LabelFeatures *generalizeimage.Labelfeatures,k int){

	auxresult := make(map[string]int)
	
	var aux_ocurrence int

	var aux_label string

	(*LabelFeatures).Allocate(generalizeimage.Resultflag,(*LabelFeatures).Getlen(generalizeimage.Trainflag),(*LabelFeatures).Getlen(generalizeimage.Knowflag))

	(*LabelFeatures).Calcdistance()

	for i := 0; i < (*LabelFeatures).Getlen(generalizeimage.Knowflag); i++ {
		auxresult[(*LabelFeatures).GetResultstring(i,generalizeimage.Labelflag)] = 0
	}

	for i := 0; i < (*LabelFeatures).Getlen(generalizeimage.Resultflag); i++ {
		
		(*LabelFeatures).Sortdist(i)
		
		for j := 0; j < k; j++ {
			auxresult[(*LabelFeatures).GetResultstring(j,generalizeimage.Labelflag)]++
		}
		aux_ocurrence = auxresult[(*LabelFeatures).GetResultstring(0,generalizeimage.Labelflag)]
		aux_label = (*LabelFeatures).GetResultstring(0,generalizeimage.Labelflag)
				
		(*LabelFeatures).SetResult(i,aux_label,aux_ocurrence)

		for g := 0; g < k; g++ {
			if (*LabelFeatures).GetResultint(i) < auxresult[(*LabelFeatures).GetResultstring(g,generalizeimage.Labelflag)]{
				
				aux_ocurrence = auxresult[(*LabelFeatures).GetResultstring(g,generalizeimage.Labelflag)]
				aux_label = (*LabelFeatures).GetResultstring(g,generalizeimage.Labelflag)
				
				(*LabelFeatures).SetResult(i,aux_label,aux_ocurrence)
			}
		}
	}
}