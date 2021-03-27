package nonparametric

import (
	"../genericdata"
	"errors"
	"fmt"
	"math"
	"sort"
	"../basicdata"

)

/**
 * [KNN description: K nearest neighbors uses the k smallests distances from the studied group to the know group]
 * @struct {[type]} dataset *dataset.DataSet   [that data set contains the study group, 
 *                                                                the know group and some parameters computed using the know group]
 * @param  {[type]} k             int                            [number of neighbors]
 */
func (thek *Knn) KNN(dataset *genericdata.DataSet,k int){
	
	auxlen,_ := (*dataset).Getlen(genericdata.Testflag)

	(*dataset).Calcdistance()

	for i := 0; i < auxlen; i++ {
		(*dataset).Sortdist(i,genericdata.Trainflag)

		(*dataset).AddInterest(i,k)
	}
	(*dataset).GetGreatestOcorrence(k)
}

// type fifo struct {
//     testA string
//     testint int
// }

// func (l *fifo) evict(c *cache) {
//     (*l).testint = 22
//     (*l).testA = "Evicting by fifo strtegy"
//     fmt.Println((*l).testA)
//     fmt.Println((*l).testint)
// }