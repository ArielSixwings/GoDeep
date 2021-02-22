package main
import (
	"fmt"
)

func main(){
	var size int
	var knowsize int
	var trainsize int
	
	size = 50
	trainsize = 20 //int(size/2.5)
	knowsize = size - trainsize

	for i := 0; i < size; i++ {
		if i < trainsize{
			fmt.Println(i)
		} else{
			fmt.Println(i-trainsize)
		}
	}
	fmt.Println("end of firts 50")
	for i := 0; i < size; i++ {
		if i < trainsize{
			fmt.Println(i+trainsize)
		} else{
			fmt.Println(i+(knowsize-trainsize))
		}
	}
	fmt.Println("end of 50 to 99")
	for i := 0; i < size; i++ {
		if i < trainsize{
			fmt.Println(i+(2*trainsize))
		} else{
			fmt.Println(i+((2*knowsize)-trainsize))
		}
	}
	fmt.Println("end of 100 to 149")
}