package main

import (
	"./imageprocessing"
	"gocv.io/x/gocv"
	"fmt"

)

/**
 *
 * @brief      { function_description }
 * @return     { description_of_the_return_value }
 */
func main() {

	var images []gocv.Mat
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		images[i] = gocv.NewMat()
	}

	imageprocessing.ReadFolder(&images,"./imageprocessing/Images/danger",true,true,true)


}
