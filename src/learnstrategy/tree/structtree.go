package tree

type StatusKeys int 

const (
	IsRoot StatusKeys = 0

	IsNode StatusKeys = 1

	IsLeaf StatusKeys = 2

)

type Node struct {

	status 	StatusKeys

	
	
	rule 	func() bool
	left	[]Node
	right	[]Node
	
}