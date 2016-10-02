package main

import(
	"fmt"
)

func Send(ch chan int){
	for i:=0;i<10;i++{
		ch<-i
	}
}

func Receive(ch chan int){
	for{
		fmt.Println(<-ch)
	}
}

func main(){
	ch := make(chan int)
	for i:=0;i<10;i++{
		go Receive(ch)
	}
	Send(ch)
}
