package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _,eachRoute := range routes{
		router.
			Methods(eachRoute.Method).
			Name(eachRoute.Name).
			Path(eachRoute.Pattern).
			HandlerFunc(eachRoute.Handler)
	}
	return router
}

type Route struct{
	Method string
	Name string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route
var detail map[string]Model

var routes = Routes{
	Route{
		"POST",
		"Store data",
		"/data",
		StoreData,
	},
	Route{
		"GET",
		"find data",
		"/data/{name}",
		FindData,
	},
}

type Model struct{
	Name string	`json:"name"`
	Phone int	`json:"phone"`
}

func StoreData(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("error")
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(422)
	}
	var model Model
	json.Unmarshal(body,&model)
	if detail == nil{
		detail = make(map[string]Model)
	}
	detail[model.Name]=model
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model)
}

func FindData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]
	if val,ok := detail[name];ok{
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(val)
	}else{
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("{Not Found}")
	}
}

func main(){
	router := NewRouter()
	http.ListenAndServe(":8080",router)
}
