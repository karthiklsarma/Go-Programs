package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')
	if(err!=nil){
		fmt.Println(err.Error())
	}
	fmt.Println(text)
}
