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
	// const var name string
	// name = "./imageprocessing/Images/danger"
	
	//left := make([]int, leftLength)
	
	var size int
	
	size  = imageprocessing.FolderLength("./imageprocessing/Images/danger")

	Images := make([]gocv.Mat,size)

	imageprocessing.ReadFolder(Images,"./imageprocessing/Images/danger",true,false,true)

	for i := 0; i < size; i++ {
		window := gocv.NewWindow("Images")
		window.IMShow(Images[i])               
		window.WaitKey(100)
	}

}
