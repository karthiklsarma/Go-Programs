package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"io/ioutil"
	"io"
)

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true);
	for _,specs:= range route_list{
		router.
			Methods(specs.Method).
			HandlerFunc(specs.Function).
			Name(specs.Name).
			Path(specs.Path)
	}
	return router
}

type Route struct{
	Name string
	Path string
	Method string
	Function http.HandlerFunc
}

type MyRoutes []Route

var route_list = MyRoutes{
	Route{
		"Employee Records",
		"/employees",
		"GET",
		GetEmployeeRecords,
	},
	Route{
		"Employee detail",
		"/employees/{empId}",
		"GET",
		GetEmployeeDetail,
	},
	Route{
		"Employee Submit",
		"/employees/",
		"POST",
		PostEmployeeDetail,
	},
}

type Employee struct{
	Id int64	`json:"id"`
	Name string	`json:"name"`
	Details string	`json:"details"`
}

type Employees []Employee
var empdata Employees

func GetEmployeeRecords(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err:=json.NewEncoder(w).Encode(empdata); err!=nil{
		panic(err)
	}
}

func PostEmployeeDetail(w http.ResponseWriter, r *http.Request){
	var emp Employee

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil{
		panic(err)
	}
	if err := r.Body.Close(); err!=nil{
		panic(err)
	}
	if err := json.Unmarshal(body, &emp);err!=nil{
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode("{invalid request}")
		panic(err)
	}else {
		empdata = append(empdata, emp)
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(empdata)
	}
}

func GetEmployeeDetail(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	empId := vars["empId"]
	for _,employee := range empdata{
		if employee.Id == empId{

		}
	}
}

func main(){
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080",router))
}
