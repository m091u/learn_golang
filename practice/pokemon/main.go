package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Pokemon app")
}

func about(w http.ResponseWriter, r *http.Request) {
	var fileName = "intro.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, nil)
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Moves          []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	}
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	//make request to API
	pokemonNumber := strconv.Itoa(rand.IntN(50))
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonNumber + "/")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//read response body
	var pokemons Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemons); err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
		return
	}

	// Serve the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func main() {
	fmt.Println("This is my attempt at creating a pokemon app")

	//set the route
	http.HandleFunc("/", home)
	http.HandleFunc("/intro", about)
	http.HandleFunc("/pokemon", GetPokemon)

	//start the server
	http.ListenAndServe("", nil)
}
