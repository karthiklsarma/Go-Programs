package main

import(
	"fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
)

func First(res http.ResponseWriter, req *http.Request){
	fmt.Println("Got request")
	http.SetCookie(res, &http.Cookie{
		Name:"My-Cookie",
		Value:"My-Value",
	})
}

func main(){
	http.HandleFunc("/", First)
	http.ListenAndServe(":8080",nil)
}