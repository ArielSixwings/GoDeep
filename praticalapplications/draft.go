package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := make([]int, 5)
	//c := make([]int, 5)

	b = b[:0]

	var j int = 0

	printSlice("a", a)
	
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	for i := 0; i < 5; i++ {
		if a[i]%2 == 0{
			b = b[:1+j]
			b[j] = a[i]
			j++
		}
	}

	a = b
	
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}