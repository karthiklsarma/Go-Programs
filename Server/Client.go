package main

import(
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Info struct{
	Name string	`json:"name"`
	Phone int	`json:"Phone"`
}
func main(){
	var inf = Info{"Joe",2345}
	body,err := json.Marshal(inf)
	req,_ := http.NewRequest("POST", "http://localhost:8080/data", bytes.NewBuffer(body))
	client:=&http.Client{}
	resp,err := client.Do(req)
	if err==nil{
		bytes,_ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bytes))
	}else{
		fmt.Println(err.Error())
	}
}
