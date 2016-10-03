package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Boring(msg string)<-chan string{
	c := make(chan string)
	go func(){
		for i:=0;;i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}
	}()
	return c
}

func Combine(a,b <-chan string)<-chan string{
	c := make(chan string)
	go func(){
		for{
			select{
			case x:=<-a: c<-x
			case x:=<-b: c<-x
			}
		}
	}()
	return c
}

func main(){
	incoming := Combine(Boring("Bob"),Boring("Julie"))
	for i:=0;i<10;i++{
		fmt.Println(<-incoming)
	}
}
