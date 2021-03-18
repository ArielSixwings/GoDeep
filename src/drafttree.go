package main

import(
	"fmt"
	"./tree"
)

func auxrule(){
	fmt.Println("test rule")
}

func main(){
	var testtree tree.Node
	testtree.Insert(auxrule)
}