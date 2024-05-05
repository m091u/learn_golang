package main

import (
	"encoding/json"
	"log"
	"net/http"
) 

// Movie represents a movie data structure
type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func main() {
	http.HandleFunc("/movies", getMovies)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Make request to the API
	resp, err := http.Get("http://www.omdbapi.com/?s=batman&apikey=YOUR_API_KEY")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	var data struct {
		Search []Movie `json:"Search"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Search)
}
