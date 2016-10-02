package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"
)

func givRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _,route:=range routes{
		router.Methods(route.Method).Name(route.Name).Path(route.Path).HandlerFunc(route.funName)
	}
	return router
}

type Route struct{
	Name string
	Path string
	Method string
	funName http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"GET",
		"/login",
		handleLogin,
	},
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	if(r.Method=="GET"){
		json.NewEncoder(w).Encode("Reached Login")
	}else if(r.Method=="POST"){

	}
}
func main(){

}


