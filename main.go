package main

import (
    "fmt"
    "log"
		"net/http"
		"encoding/json"
)

type Restaurant struct {
	Name string `json:"Name"`
	Desc string `json:"Desc"`
}

var restaurantList []Restaurant

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to the HomePage!")
  fmt.Println("Endpoint Hit: homePage")
}

func getAllRestaurants(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: getAllRestaurants")
	json.NewEncoder(w).Encode(restaurantList)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/restaurants", getAllRestaurants)
  log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	restaurantList = []Restaurant{
		Restaurant{Name: "Popeyes", Desc: "Oily af chicken"},
		Restaurant{Name: "Foodie Fruitie", Desc: "Food poisoning"},
	}

  handleRequests()
}