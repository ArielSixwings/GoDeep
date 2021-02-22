//package nonparametric
package main

import (
	//"gocv.io/x/gocv"
	"fmt"
	"math"
	"sort"

)

type labelfeatures struct {

	studyfeatures [3]float64 

	knowfeatures [3]float64
	
	label string

	distance float64
}

func (lf *labelfeatures) calcdistance() { 

	var sum float64 = 0.0

	for i := 0; i < 3; i++ {
		sum += (math.Pow(lf.knowfeatures[i] - lf.studyfeatures[i],2))
	}
	fmt.Println(sum)
	fmt.Println(math.Sqrt(sum))
	(*lf).distance = math.Sqrt(sum)

}

type ByDistance []labelfeatures

func (d ByDistance) Len() int { return len(d) }

func (d ByDistance) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d ByDistance) Less(i, j int) bool {
	return d[i].distance < d[j].distance
}

func KNN(LabelFeatures []labelfeatures, k int){
	fmt.Println("just a tempalte", math.Pow(float64(k),2))
}

func main(){

	knowfeatures := [3]float64{1,2,3}

	studyfeatures := [3]float64{3,3,3}

	knowfeatures2 := [3]float64{1,3,3}

	test := []labelfeatures{
		{studyfeatures,knowfeatures,"danger",5},
		{studyfeatures,knowfeatures2,"safe",0},
	}
	
	for i := 0; i < 2; i++ {
		
		test[i].calcdistance()
		
	}

	sort.Sort(ByDistance(test))

	fmt.Println(test)
}