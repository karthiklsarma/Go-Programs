package main

import(
	"fmt"
	"container/heap"
)

type MyHeap []int

func (pq MyHeap)Less(i,j int)bool	{return pq[i]<pq[j]}
func (pq MyHeap)Swap(i,j int)	{pq[i],pq[j]=pq[j],pq[i]}
func (pq MyHeap)Len() int 	{return len(pq)}

func (p *MyHeap)Push(item interface{}){
	*p = append(*p,item.(int))
}

func (p *MyHeap)Pop()interface{}{
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0:n-1]
	return x
}

func main(){
	h := &MyHeap{2,3,4}
	heap.Init(h)
	heap.Push(h,3)
	fmt.Printf("Minimum = %d",(*h)[0])
	for h.Len()>0{
		fmt.Printf("%d ",heap.Pop(h))
	}
}