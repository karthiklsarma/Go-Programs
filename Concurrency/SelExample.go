package main

import(
	"fmt"
	"time"
	"sync"
)

func main(){
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		for {
			select {
			case val := <-ch:
				fmt.Printf("Got value, %d", val)
				wg.Done()
				return;
			case <-time.After(2 * time.Second):
				fmt.Println("2 seconds over. Still waiting")
			}
		}
	}()
	go func(){
		time.Sleep(6*time.Second)
		ch<-7
		wg.Done()
	}()
	wg.Wait()
}
