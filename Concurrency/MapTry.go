package main

import(
	"fmt"
)

type Node struct{
	Name string
	Left *Node
	Right *Node
}

func DFS(root *Node){
	if(root!=nil){
		fmt.Println(root.Name)
		DFS(root.Left)
		DFS(root.Right)
	}
}

func BFS(root *Node){
	q := make([]*Node,0)
	q = append(q,root)
	for len(q)!=0{
		node := q[0]
		q = q[1:]
		fmt.Println(node.Name)
		if node.Left!=nil{
			q = append(q,node.Left)
		}
		if node.Right!=nil{
			q = append(q,node.Right)
		}
	}
}

func main(){
	myMap := make(map[string]*Node)
	a := &Node{Name:"first",Left:nil,Right:nil}
	myMap[a.Name] = a
	b := &Node{Name:"second",Left:nil,Right:nil}
	myMap[b.Name] = b
	c := &Node{Name:"third",Left:nil,Right:nil}
	myMap[c.Name] = c
	a.Left = b
	a.Right = c
	BFS(a)
}
