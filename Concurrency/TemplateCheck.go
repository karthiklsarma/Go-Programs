package main

import(
	"fmt"
	"html/template"
)

func main(){
	fmt.Println(template.HTMLEscapeString("<script>alert()</script>"))
}
