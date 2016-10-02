package main

import(
	"fmt"
	"sync"
)

func start(ch chan int, wg *sync.WaitGroup){
	fmt.Println(<-ch)
	wg.Done()
}

func readInput(ch chan int, wg *sync.WaitGroup){
	var i int
	fmt.Scanf("%d",&i)
	ch<-i
	wg.Done()
}

func main(){
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go readInput(ch,&wg)
	go start(ch,&wg)
	wg.Wait()
}