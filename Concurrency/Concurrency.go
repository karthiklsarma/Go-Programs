package main

import(
	"fmt"
	"sync"
)

func PrintFirst(wg *sync.WaitGroup){
	fmt.Println("one")
	wg.Done()
}

func PrintTwo(wg *sync.WaitGroup){
	fmt.Println("Two")
	wg.Done()
}

func PrintThree(wg *sync.WaitGroup){
	fmt.Println("Three")
	wg.Done()
}

func main(){
	var vg sync.WaitGroup
	vg.Add(3)
	go PrintFirst(&vg)
	go PrintTwo(&vg)
	go PrintThree(&vg)
	vg.Wait()
}
