/**
 * [return lenght of the dist of the ByDist sort template]
 * @struct {[type]}   (d ByDist)  [description]
 * @return {[type]}   int         [description]
 */
func (d ByDist) Len() int { return len(d) }

/**
 * [func description]
 * @struct  {[type]}    d ByDist
 * @param i {[type]} int
 * @param j {[type]} int [description]
 * @return  {[type]}   [description]
 */
func (d ByDist) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

/**
 * [return the samlest distance betwee the ith and jth entry]
 * @struct {[type]} d ByDist)       Less(i, j int [description]
 * @return {[type]}   [description]
 */
func (d ByDist) Less(i, j int) bool { return d[i].Dist < d[j].Dist }

/**
 * [for the ith entry, selected by the i parameter, sort all the distances between the entry and the train group]
 * @struct {[type]} ds *DataSet)					[the data set]
 * @param  Sortdist(i int)
 * @return {error} 	 											[gets errors]
 */
func (ds *DataSet) Sortdist(i int, sortflag Groupflag) error {

	switch sortflag {
	case Centerdistflag:

		if len((*ds).is_sortedbydist) != 0 {
			if (*ds).is_sortedbydist[i] {
				return errors.New("the distance set of this dataset are already sorted by the distance to the group center")
			} else {
				if len((*ds).result) == 0 {
					return errors.New("result weren't computed")
				}
			}
		}
		sort.Sort(ByDist((*ds).result[i].F_point))
		(*ds).result[i].Learnedlabel = (*ds).result[i].F_point[0].Distlabel

		(*ds).is_sortedbycenter[i] = true
	case Trainflag:
		if (*ds).is_sortedbydist[i] {
			return errors.New("the distance set of this dataset are already sorted")
		} else {
			if len((*ds).result) == 0 {
				return errors.New("result weren't computed")
			}
		}
		sort.Sort(ByDist((*ds).result[i].F_point))

		(*ds).is_sortedbydist[i] = true
	}
	return nil
}