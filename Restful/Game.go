package main

import(
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"fmt"
)

type Team_Info struct{
	Team string `json:"team"`
	Players []Out_Player_Info `json:"players"`
}

type Out_Player_Info struct{
	Name string	`json:"name"`
	Number	int	`json:"number"`
}

type In_Player_Info struct{
	First string `json:"first_name"`
	Last string `json:"last_name"`
	Team string `json:"team"`
	Number int `json:"number"`
	HomeTown string `json:"hometown"`
	Position string `json:"position"`
	Height string `json:"height"`
	Weight string `json:"weight"`
}

type In_Team_Infor struct{
	Name string `json:"name"`
	League string `json:"league"`
	Players []string `json:"players"`
	Games []string`json:"games"`
}

type Route struct{
	Method string
	Name	string
	Pattern string
	Handler	http.HandlerFunc
}

func GivRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _,routes := range routes{
		router.
		Methods(routes.Method).
		Name(routes.Name).
		Path(routes.Pattern).
		HandlerFunc(routes.Handler)
	}
	return router
}

type Rts []Route

var routes = Rts{
	{
		"GET",
		"Get Team Info",
		"/stats",
		GetStats,
	},
}

func GetStats(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	url := "http://madness-sim.btrll.com/teams/"+GetTeamName()
	request,err := http.NewRequest("GET",url, nil)
	var inTeam In_Team_Infor

	if err==nil{
		resp,_ := client.Do(request)
		body,_ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &inTeam)
		team_info := &Team_Info{}
		team_info.Team = GetTeamName()
		team_info.Players = make([]Out_Player_Info,0)
		for _,playerurl := range inTeam.Players{
			var in_player In_Player_Info
			fmt.Println(string(playerurl))
			inrequest,_ := http.NewRequest("GET",string(playerurl), nil)
			player_response,_ := client.Do(inrequest)
			body,_ := ioutil.ReadAll(player_response.Body)
			json.Unmarshal(body, &in_player)
			out_player := Out_Player_Info{}
			out_player.Name = in_player.First + " " + in_player.Last
			out_player.Number = in_player.Number
			fmt.Println(out_player.Name)
			team_info.Players = append(team_info.Players,out_player)
			fmt.Println((playerurl))
		}
		json.NewEncoder(w).Encode(team_info)
	}

}

func GetTeamName() string{
	Team_Name := "san_francisco"
	return Team_Name
}

func main(){
	router := GivRouter()
	http.ListenAndServe(":8080",router)
}