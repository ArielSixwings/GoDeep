package generalizeimage

import (
	"fmt"
)
/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printresults( [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Printresults(){

	var total , right float64

	right = 0
	total = float64( len(lf.result))
	fmt.Println("These are the results")
	for i := 0; i < len(lf.result); i++ {
		fmt.Println(lf.result[i].learnedlabel,"and that label status is:", (lf.result[i].learnedlabel == lf.train[i].label))
		if lf.result[i].learnedlabel == lf.train[i].label {
			right++ 
		}
	}
	fmt.Println("Success rate = ",100*(right/total),"%")
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printdists( [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Printdists(){
	fmt.Println("These are the results")
	for i := 0; i < len(lf.result); i++ {
		fmt.Println("results: ", i)
		fmt.Println(lf.result[i])	
	}	
}

/**
 * [func description]
 * @param  {[type]} lf Labelfeatures) Printinterest( [description]
 * @return {[type]}    [description]
 */
func (lf Labelfeatures) Printinterest(){
	for i := 0; i < len(lf.interestgroup); i++ {
		fmt.Println(lf.interestgroup[i])
	}
}