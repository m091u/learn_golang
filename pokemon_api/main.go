package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	Name    string `json:"name"`
	Weight  int    `json:"weight"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

type PokemonList struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func contains(slice []int, num int) bool {
	for _, n := range slice {
		if n == num {
			return true
		}
	}
	return false
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	var pokemons []Pokemon

	// read from json file
	file, err := os.ReadFile("database.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON data from the file
	var pokemonList PokemonList
	if err := json.Unmarshal(file, &pokemonList); err != nil {
		log.Fatal(err)
	}

	// Extract IDs from URLs and populate the ids slice
	ids := []int{}

	for _, result := range pokemonList.Results {
		idStr := strings.TrimSuffix((result.URL[34:]), "/")
		id, _ := strconv.Atoi(idStr)
		ids = append(ids, id)
	}

	//random ID generation
	// for i := 0; i < 10; i++ {
	// 	id := rand.Intn(1302)
	// 	for {
	// 		if !contains(ids, id) {
	// 			ids = append(ids, id)
	// 			break
	// 		}
	// 	}
	// }

	// Retrieve Pokémon data from the PokeAPI
	for i, _ := range ids[:10] {
		url := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(ids[i])
		log.Println(url)
		resp, err := http.Get(url)

		if err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			return
		}

		var pokemon Pokemon
		_ = json.NewDecoder(resp.Body).Decode(&pokemon)
		pokemons = append(pokemons, pokemon)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/pokemons", getPokemon).Methods("GET")

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("", router))
}
