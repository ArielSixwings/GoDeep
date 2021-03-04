package generalizecartesian

func (lf *Labelfeatures) Filterdataset( rule func(...float64) bool) error{
	
	temp := make([]features, len((*lf).know))

	temp = temp[:0]

	var j int = 0
	
	for _ ,i := range (*lf).know {
		if rule(radius, currentdist){
			temp = temp[:1+j]
			temp[j] = (*lf).know[i]
			j++
		}
	}

	(*lf).know = temp
}
