package learnstrategy

import(
	"fmt"
	"math"
	"../basicdata"
)

func (ds *DataLearner) Filterdataset( rule func(int) bool) error{
	
	temp := make([]cartesian.Features, len((*ds).train))

	temp = temp[:0]

	var j int = 0
	
	for i := 0; i < len((*ds).train); i++ {
		if rule(i){
			fmt.Println("that data was aproved: ",i,"  ",(*ds).train[i])
			temp = temp[:1+j]
			temp[j] = (*ds).train[i]
			j++
		}
	}

	(*ds).train = temp

	return nil
}

func (ds *DataLearner) PurplezoneRule(i int) bool{

	if (*ds).insidetworadius(i,0,1) || (*ds).insidetworadius(i,0,2) || (*ds).insidetworadius(i,1,2) {
		return false
	}else{
		return true
	}
}

func (ds *DataLearner) RedzoneRule(i int) bool{

	if (*ds).insidetworadius(i,0,1) && (*ds).insidetworadius(i,0,2) && (*ds).insidetworadius(i,1,2) {
		return false
	}else{
		return true
	}
}

// func (ds *DataSet) MaxCaoszoneRule(i int) bool{

// 	if (*ds).RedzoneRule(i,0,1) && (*ds).Maxcentroid(i) {
// 		return false
// 	}else{
// 		return true
// 	}
// }

func (ds *DataLearner) MinCaoszoneRule(i int) bool{

	if (*ds).RedzoneRule(i) && ((*ds).allcentroid.Minradius > (*ds).euclidiandistance((*ds).train[i].Features,(*ds).allcentroid.Features))  {
		return false
	}else{
		return true
	}
}

func (ds *DataLearner) insidetworadius(i int, groupA int, groupB int) bool{
	
	var insideA bool = (*ds).euclidiandistance((*ds).train[i].Features,(*ds).centroid[groupA].Features) < (*ds).centroid[groupA].Radius
	var insideB bool = (*ds).euclidiandistance((*ds).train[i].Features,(*ds).centroid[groupB].Features) < (*ds).centroid[groupB].Radius
	return insideA && insideB
}

func (ds *DataLearner) euclidiandistance(x [3]float64,y [3]float64) float64{
	
	var sum float64 = 0.0

	for f := 0; f < 3; f++ {
		sum += math.Pow(x[f] - y[f],2)
	}

	return math.Sqrt(sum)
}

