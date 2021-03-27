package tree

import(
	"fmt"
	"math"
)
/*
if root == null
	insert at that point

	else
		if rule()
			insert(at left)
		else
			insert(at right) 

 */
func auxsecondrule() bool{
	fmt.Println("internal node, a leaf!!!")
	return false
}

func (nd *Node) Insert(newrule func()bool) error {
	if nd.rule == nil {
		nd.rule = newrule
		return nil
	} else{
		if nd.rule(){
			(*nd).left = make([]Node, 1)
			(*nd).left[0].Insert(auxsecondrule)
			return nil
		}else{
			(*nd).right = make([]Node, 1)
			(*nd).right[0].Insert(auxsecondrule)
			return nil
		}

	}
}

func (nd *Node) ginileafimpurity() error {

}