package main

import (
	"./imageprocessing"
	"gocv.io/x/gocv"
	// "fmt"

)

/**
 *
 * @brief      { function_description }
 * @return     { description_of_the_return_value }
 */
func main() {

	var images []gocv.Mat
	imageprocessing.ReadFolder(&images,"./imageprocessing/Images/danger",true,false,true)


}
