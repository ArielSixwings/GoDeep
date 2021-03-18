package tree

import(
	"fmt"
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

func (nd *Node) Insert(newrule func()) error {
	if nd.rule == nil {
		fmt.Println("nd is nil",nd.rule)
		nd.rule = newrule
		fmt.Println("the setted rule: ")
		nd.rule()
		return nil
	} else{
		fmt.Println("nd is	NOT nil",nd)
		return nil		
	}
}