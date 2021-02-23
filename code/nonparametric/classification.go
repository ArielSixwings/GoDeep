package nonparametric

import (
	"fmt"
	"math"
	"sort"

)

func KNN(LabelFeatures *labelfeatures, auxresult *map[string]int,k int){

	(*LabelFeatures).calcdistance()

	for i := 0; i < len((*LabelFeatures).know); i++ {
		(*auxresult)[(*LabelFeatures).result[i].label] = 0
	}

	for i := 0; i < len((*LabelFeatures).result); i++ {
		
		sort.Sort(Bydist((*LabelFeatures).result[i]))
		
		for j := 0; j < k; j++ {
			(*auxresult)[(*LabelFeatures).result[j].label]++
		}
		
		(*LabelFeatures).result[i].greatestoccurrence = (*auxresult)[(*LabelFeatures).result[0]
		(*LabelFeatures).result[i].learnedlabel = (*LabelFeatures).result[0].label

		for j := 0; j < k; j++ {
			if (*LabelFeatures).result[i].greatestoccurrence < (*auxresult)[(*LabelFeatures).result[j].label]{
				(*LabelFeatures).result[i].greatestoccurrence = (*auxresult)[(*LabelFeatures).result[j].label]
				(*LabelFeatures).result[i].learnedlabel = (*LabelFeatures).result[j].label
			}
		}
		
	}

}