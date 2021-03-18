package tree

//type feature float64

type Node struct {
	rule 	func()
	left	*Node
	right	*Node
	
}