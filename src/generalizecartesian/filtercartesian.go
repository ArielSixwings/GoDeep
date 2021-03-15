package generalizecartesian

import(
	"fmt"
	"math"
	"../basicdata"
)

func (lf *Labelfeatures) Filterdataset( rule func(int) bool) error{
	
	temp := make([]cartesian.Features, len((*lf).know))

	temp = temp[:0]

	var j int = 0
	
	for i := 0; i < len((*lf).know); i++ {
		if rule(i){
			fmt.Println("that data was aproved: ",i,"  ",(*lf).know[i])
			temp = temp[:1+j]
			temp[j] = (*lf).know[i]
			j++
		}
	}

	(*lf).know = temp

	return nil
}

func (lf *Labelfeatures) PurplezoneRule(i int) bool{

	if (*lf).insidetworadius(i,0,1) || (*lf).insidetworadius(i,0,2) || (*lf).insidetworadius(i,1,2) {
		return false
	}else{
		return true
	}
}

func (lf *Labelfeatures) RedzoneRule(i int) bool{

	if (*lf).insidetworadius(i,0,1) && (*lf).insidetworadius(i,0,2) && (*lf).insidetworadius(i,1,2) {
		return false
	}else{
		return true
	}
}

// func (lf *Labelfeatures) MaxCaoszoneRule(i int) bool{

// 	if (*lf).RedzoneRule(i,0,1) && (*lf).Maxcentroid(i) {
// 		return false
// 	}else{
// 		return true
// 	}
// }

func (lf *Labelfeatures) MinCaoszoneRule(i int) bool{

	if (*lf).RedzoneRule(i) && ((*lf).allcentroid.Minradius > (*lf).euclidiandistance((*lf).know[i].Features,(*lf).allcentroid.Features))  {
		return false
	}else{
		return true
	}
}

func (lf *Labelfeatures) insidetworadius(i int, groupA int, groupB int) bool{
	
	var insideA bool = (*lf).euclidiandistance((*lf).know[i].Features,(*lf).centroid[groupA].Features) < (*lf).centroid[groupA].Radius
	var insideB bool = (*lf).euclidiandistance((*lf).know[i].Features,(*lf).centroid[groupB].Features) < (*lf).centroid[groupB].Radius
	return insideA && insideB
}

func (lf *Labelfeatures) euclidiandistance(x [3]float64,y [3]float64) float64{
	
	var sum float64 = 0.0

	for f := 0; f < 3; f++ {
		sum += math.Pow(x[f] - y[f],2)
	}

	return math.Sqrt(sum)
}

