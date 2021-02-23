package nonparametric

import (
	"fmt"
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

	for i := 0; i < (*LabelFeatures).Getlen(generalizeimage.Trainflag); i++ {
		auxresult[(*LabelFeatures).GetKnowstring(i)] = 0
	}

	for i := 0; i < (*LabelFeatures).Getlen(generalizeimage.Trainflag); i++ {
		
		(*LabelFeatures).Sortdist(i)
		
		for j := 0; j < k; j++ {
			fmt.Println("current label", (*LabelFeatures).GetKnowstring(j))
			auxresult[(*LabelFeatures).GetKnowstring(j)]++
		}
		aux_ocurrence = auxresult[(*LabelFeatures).GetKnowstring(0)]
		aux_label = (*LabelFeatures).GetKnowstring(0)
				
		(*LabelFeatures).SetResult(i,aux_label,aux_ocurrence)

		for g := 0; g < k; g++ {
			if (*LabelFeatures).GetResultint(i) < auxresult[(*LabelFeatures).GetKnowstring(g)]{
				
				aux_ocurrence = auxresult[(*LabelFeatures).GetKnowstring(g)]
				aux_label = (*LabelFeatures).GetKnowstring(g)
				
				(*LabelFeatures).SetResult(i,aux_label,aux_ocurrence)
			}
		}
	}
}