package main

import(
	"fmt"
)

func Player(ch chan int,name string, num int)<-chan int{
	go func(){
		for {
			fmt.Println(name)
			ch <- num
			num++
		}
	}()
	return ch
}

func FanIn(ch1,ch2 <-chan int)<-chan int{
	c := make(chan int)
	go func() {
		for {
			select {
			case n:=<-ch1: c<-n
			case n:=<-ch2: c<-n
			}
		}
	}()
	return c
}
func main(){
	ch := make(chan int)
	m := FanIn(Player(ch,"Bob",0),Player(ch,"Molly",10))
	for i:=0;i<5;i++{
		fmt.Println(<-m)
	}
}
