package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"recipes/model"

	"github.com/gorilla/mux"
)

var recipes []model.Recipe

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {

	var recipe model.Recipe
	// _ discards the error coming from decoding
	//The Decode() function returns two values: the decoded data (in this case, the movie variable) and an error (if any).
	// json.NewDecoder(r.Body) creates a new JSON decoder that reads from the request body (r.Body).
	// .Decode(&movie) decodes the JSON data from the decoder and stores it in the movie variable.
	// The & symbol before movie is used to pass the memory address of the movie variable to the Decode() function.
	_ = json.NewDecoder(r.Body).Decode(&recipe)
	recipe.ID = strconv.Itoa(rand.Intn(10000))
	recipes = append(recipes, recipe)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, recipe := range recipes {

		if recipe.ID == params["id"] {
			recipes = append(recipes[:i], recipes[i+1:]...)
			fmt.Fprintf(w, "Recipe with id %s deleted", params["id"])
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Recipe with id %s not found", params["id"])
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, recipe := range recipes {
		if recipe.ID == params["id"] {
			json.NewEncoder(w).Encode(recipe)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "recipe with id %s not found", params["id"])
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var updatedRecipe model.Recipe

	_ = json.NewDecoder(r.Body).Decode(&updatedRecipe)

	for i, recipe := range recipes {
		if  recipe.ID == id {
			recipes[i]= updatedRecipe
			fmt.Fprintf(w, "Recipe with id %s was updated", params["id"])
			json.NewEncoder(w).Encode(updatedRecipe)
		}
	}
}

func main() {
	router := mux.NewRouter()

	recipes = append(recipes, model.Recipe{
		ID:              "1",
		Name:            "Spaghetti Carbonara",
		Cuisine:         "Italian",
		PreparationTime: 30,
		Link:            "https://www.example.com/spaghetti-carbonara",
	})
	recipes = append(recipes, model.Recipe{
		ID:              "2",
		Name:            "Chicken Tikka Masala",
		Cuisine:         "Indian",
		PreparationTime: 45,
		Link:            "https://www.example.com/chicken-tikka-masala",
	})
	recipes = append(recipes, model.Recipe{
		ID:              "3",
		Name:            "Guacamole",
		Cuisine:         "Mexican",
		PreparationTime: 15,
		Link:            "https://www.example.com/guacamole",
	})

	router.HandleFunc("/recipes", getRecipes).Methods("GET")
	router.HandleFunc("/recipes", createRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", getRecipe).Methods("GET")
	router.HandleFunc("/recipes/{id}", deleteRecipe).Methods("DELETE")
	router.HandleFunc("/recipes/{id}", updateRecipe).Methods("PUT")

	fmt.Println("Server running...")
	log.Fatal(http.ListenAndServe("", router))
}
