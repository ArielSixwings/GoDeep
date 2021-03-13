//package main
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a [5120]string
	var c int = 0

	for j := 0; j < 20; j++ {

		for i := c; i < 5120; i++ {
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
			if i == 256 {
				c += 256
				continue
			}
		}
	}
	var chamada int = 0
	for j := 0; j < 5120; j++ {
		fmt.Print(a[j])
		if j%256 == 0 && j != 0 {
			fmt.Println("chamada:", chamada)
			chamada = chamada + 1
		}
	}
}
