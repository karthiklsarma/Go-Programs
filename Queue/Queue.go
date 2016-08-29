package Queue

import(
	"errors"
)
type T_node struct{
	data int
	left *T_node
	right *T_node
}

type Queue []*T_node

func (q *Queue)Push(n *T_node){
	*q = append(*q,n)
}

func (q *Queue)Pop()(*T_node, error){
	if q.Len()<=0{
		return nil,errors.New("")
	}
	n := (*q)[0]
	(*q) = (*q)[1:]
	return n,nil
}

func (q *Queue)Len()int{
	return len(*q)
}

func (node *T_node) SetData(val int){
	node.data = val
}

func (node *T_node) SetLeft(val int){
	node.left = new(T_node)
	node.left.SetData(val)
}

func (node *T_node) SetRight(val int){
	node.right = new(T_node)
	node.right.SetData(val)
}

func (node *T_node) GetData() int{
	return (*node).data
}

func (node *T_node) GetLeft() *T_node{
	return (*node).left
}

func (node *T_node) GetRight() *T_node{
	return (*node).right
}