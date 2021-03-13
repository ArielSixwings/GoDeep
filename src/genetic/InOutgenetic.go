//package main
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a [10000]string
	var b [50]string

	for j := 0; j < 2; j++ {

		for i := 0; i < 10000; i++ {
			randonInt := rand.Intn(9)
			if randonInt == 3 {
				a[i] = "_"
			} else if randonInt == 2 || randonInt == 5 {
				a[i] = "A"
			} else if randonInt == 1 || randonInt == 8 {
				a[i] = "C"
			} else if randonInt == 0 || randonInt == 4 {
				a[i] = "T"
			} else if randonInt == 6 || randonInt == 7 {
				a[i] = "G"
			}
			if i == 50 {
				continue
				i = 51
			}
		}
	}

	fmt.Println("\nA:")
	for j := 0; j < 1000; j++ {
		fmt.Print(a[j])
	}