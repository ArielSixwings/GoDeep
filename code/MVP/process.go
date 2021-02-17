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

	fmt.Println("oi1")
	testcolor := gocv.NewMat() //mat for histogram equalization
	testgray := gocv.NewMat()
	fmt.Println("oi2")

	testcolor = imageprocessing.ReadImage(testcolor, "./imageprocessing/Images/danger/danger_08.png", true, false, true)
	testgray = imageprocessing.ReadImage(testgray, "./imageprocessing/Images/danger/danger_08.png", true, false, false)


}
