package learnstrategy

import (
	"fmt"
)

/**
 * [Printresults description: print at the terminal the computed results, with Success rate]
 * @struct {[type]}  ds DataSet [thats a DataSet method]
 * @return {[type]}  error            [gets errors]
 */
func (ds DataLearner) Printresults(){

	var total , right float64

	right = 0
	total = float64( len(ds.result))
	
	fmt.Println("These are the results")
	
	for i := 0; i < len(ds.result); i++ {
		
		fmt.Println(ds.result[i].Learnedlabel,"Label should be: ",ds.test[i].Label," and that Label status is:", (ds.result[i].Learnedlabel == ds.test[i].Label))
		
		if ds.result[i].Learnedlabel == ds.test[i].Label {
			right++ 
		}
	}
	fmt.Println("Total = ",total)
	fmt.Println("Success rate = ",100*(right/total),"%")
}

/**
 * [func description: print at the terminal the computed distances]
 * @struct {[type]}  ds DataSet [thats a DataSet method]
 * @return {[type]}  error            [gets errors]
 */
func (ds DataLearner) Printdists(){
	
	fmt.Println("These are the distances between the data set groups")
	
	for i := 0; i < len(ds.result); i++ {
		
		fmt.Println("results: ", i)
		for j  := 0; j  < len(ds.result[i].F_point); j ++ {
			fmt.Println(ds.result[i].F_point[j].Dist)	
		}
	}	
}

/**
 * [Printinterest description: print interest group]
 * @struct {[type]}  ds DataSet [thats a DataSet method]
 * @return {[type]}  error            [gets errors]
 */
func (ds DataLearner) Printinterest(){
	for i := 0; i < len(ds.interestgroup); i++ {
		fmt.Println(ds.interestgroup[i])
	}
}

/**
 * [Printfeatures description: print groups features]
 * @struct {[type]}  ds DataSet [thats a DataSet method]
 * @return {[type]}  error            [gets errors]
 */
func (ds DataLearner) Printfeatures() error{
	
	fmt.Println("These are the train features")
	for i := 0; i < len(ds.train); i++ {
		fmt.Println(ds.train[i])	
	
	}
	
	fmt.Println("These are the test features")
	for i := 0; i < len(ds.test); i++ {
		fmt.Println(ds.test[i])
	}

	return nil
}