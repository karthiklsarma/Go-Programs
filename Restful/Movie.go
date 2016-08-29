package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"io"
	"log"
)


func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		router.
		Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			HandlerFunc(route.Handler)
	}
	return router
}

type Route struct{
	Name string
	Method string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
Route{
	"New movie",
	"PUT",
	"/movies",
	AddMovie,
},
Route{
	"Get Movie Detail",
	"GET",
	"/movies/{movieId}",
	GetMovieDetails,
},
Route{
	"Delete Movie",
	"DELETE",
	"/movies/{movieId}",
	DeleteMovieDetail,
},
}

type Movie struct{
	MovieId string	`json:"movieid"`
	MovieName string `json:"moviename"`
	Rating	int32 `json:"rating"`
}

var movieMap map[string]*Movie

func DeleteMovieDetail(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	movieid := vars["movieId"]
	if _,ok := movieMap[movieid];ok{
		delete(movieMap,movieid)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		json.NewEncoder(w).Encode("{Deleted}")
	}else{
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		http.Error(w,"Not Found",404)
	} 
}

func GetMovieDetails(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	movieid := vars["movieId"]
	log.Println(movieid)
	if movieDetails,ok := movieMap[movieid];ok{
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application-json; charset=UTF-8")
		json.NewEncoder(w).Encode(movieDetails)
	}else{
		w.Header().Set("Content-Type","application-json; charset=UTF-8")
		http.Error(w,"Not Found",404)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request){
	body,err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err!=nil{
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err);err!=nil{
			panic(err)
		}
	}
	var movie Movie
	if json.Unmarshal(body, &movie);err!=nil{
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(err)
	}else{
		w.Header().Set("Content-Type","application-json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if movieMap==nil {
			movieMap = make(map[string]*Movie)
		}
		movieMap[movie.MovieId]=&movie
		json.NewEncoder(w).Encode(movie)
	}
}

func main(){
	router := NewRouter()
	http.ListenAndServe(":8080",router)
}
