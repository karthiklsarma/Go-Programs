package main

import(
	"fmt"
)

type Player struct{
}

func (p *Player)Play(name string, table chan int){
	for{
		cnt := <-table
		fmt.Println(name,cnt)
		cnt++
		table<- cnt
	}
}

func main(){
	table := make(chan int)
	player := &Player{}
	go player.Play("Ping", table)
	go player.Play("Pong", table)
	table<-0
}