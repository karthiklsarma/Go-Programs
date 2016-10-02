package main

import(
	"fmt"
	"sync"
)

func Gen(nums ...int)<-chan int{
	out := make(chan int)
	go func(){
		for _,num := range nums{
			out<-num
		}
		close(out)
	}()
	return out
}

func Sqrt(in ...<-chan int)<-chan int{
	out := make(chan int)
	for _,inchannel := range in{
		go func(ch <-chan int){
			for num := range ch{
				out<-num*num
			}
		}(inchannel)
		close(out)
	}
	return out
}

func main(){
	genResult := Gen(1,2,3)
	var wg sync.WaitGroup
	out := Sqrt(genResult)
	for ch := range out{
		fmt.Println(ch)
	}
}
