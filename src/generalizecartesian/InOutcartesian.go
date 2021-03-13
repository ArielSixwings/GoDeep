package generalizecartesian

import (
	"fmt"
)

/**
 * [Printresults description: print at the terminal the computed results, with Success rate]
 * @struct {[type]}  lf Labelfeatures [thats a Labelfeatures method]
 * @return {[type]}  error            [gets errors]
 */
func (lf Labelfeatures) Printresults(){

	var total , right float64

	right = 0
	total = float64( len(lf.result))
	
	fmt.Println("These are the results")
	
	for i := 0; i < len(lf.result); i++ {
		
		fmt.Println(lf.result[i].learnedlabel,"label should be: ",lf.train[i].label," and that label status is:", (lf.result[i].learnedlabel == lf.train[i].label))
		
		if lf.result[i].learnedlabel == lf.train[i].label {
			right++ 
		}
	}
	fmt.Println("Total = ",total)
	fmt.Println("Success rate = ",100*(right/total),"%")
}

/**
 * [func description: print at the terminal the computed distances]
 * @struct {[type]}  lf Labelfeatures [thats a Labelfeatures method]
 * @return {[type]}  error            [gets errors]
 */
func (lf Labelfeatures) Printdists(){
	
	fmt.Println("These are the distances between the data set groups")
	
	for i := 0; i < len(lf.result); i++ {
		
		fmt.Println("results: ", i)
		for j  := 0; j  < len(lf.result[i].f_point); j ++ {
			fmt.Println(lf.result[i].f_point[j].dist)	
		}
	}	
}

/**
 * [Printinterest description: print interest group]
 * @struct {[type]}  lf Labelfeatures [thats a Labelfeatures method]
 * @return {[type]}  error            [gets errors]
 */
func (lf Labelfeatures) Printinterest(){
	for i := 0; i < len(lf.interestgroup); i++ {
		fmt.Println(lf.interestgroup[i])
	}
}

/**
 * [Printfeatures description: print groups features]
 * @struct {[type]}  lf Labelfeatures [thats a Labelfeatures method]
 * @return {[type]}  error            [gets errors]
 */
func (lf Labelfeatures) Printfeatures() error{
	
	fmt.Println("These are the know features")
	for i := 0; i < len(lf.know); i++ {
		fmt.Println(lf.know[i])	
	
	}
	
	fmt.Println("These are the train features")
	for i := 0; i < len(lf.train); i++ {
		fmt.Println(lf.train[i])
	}

	return nil
}