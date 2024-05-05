package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"movie_recomm/model"

	"github.com/gorilla/mux"
)

var movies []model.Movie
var users []model.User

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/recommendations", getRecommendations).Methods("GET")
	router.HandleFunc("/register", registerUser).Methods("POST")
	router.HandleFunc("/login", loginUser).Methods("POST")
	router.HandleFunc("/logout", logoutUser).Methods("POST")

	log.Println("Server running...")
	log.Fatal(http.ListenAndServe("", router))
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	//user already exists
	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("User already exists")
			return
		}
	}

	//add user to list
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Check if user exists and password matches
	for _, existingUser := range users {
		if existingUser.Username == user.Username && existingUser.Password == user.Password {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Login successful")
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Invalid username or password")
}

func logoutUser(w http.ResponseWriter, r *http.Request) {

	// Clear the session token by setting an expired cookie
	cookie := http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour), //expiration in the past
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Logout successful")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies = fetchMovies()
	json.NewEncoder(w).Encode(movies)
}

func getRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies = fetchMovies()

	var recommendations []model.Movie

	for _, movie := range movies {
		if movie.Genre == "Action" {
			recommendations = append(recommendations, movie)
		}
	}
	json.NewEncoder(w).Encode(recommendations)
}

func fetchMovies() []model.Movie {
	// Mock data for demonstration (replace with actual API call to JustWatch)
	movies := []model.Movie{
		{ID: "1", Title: "Movie 1", Genre: "Action", Platform: "Netflix"},
		{ID: "2", Title: "Movie 2", Genre: "Comedy", Platform: "Amazon Prime"},
		{ID: "3", Title: "Movie 3", Genre: "Drama", Platform: "Hulu"},
		{ID: "4", Title: "Movie 4", Genre: "Thriller", Platform: "Netflix"},
		{ID: "5", Title: "Movie 5", Genre: "Action", Platform: "Amazon Prime"},
	}
	return movies
}
