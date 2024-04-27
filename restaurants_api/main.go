package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Restaurant represents a restaurant data structure
type Restaurant struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Cuisines string `json:"cuisines"`
	Rating   struct {
		StarRating float64 `json:"starRating"`
	} `json:"rating"`
	Address struct {
		FirstLine  string `json:"firstLine"`
		PostalCode string `json:"postalCode"`
	} `json:"address"`
}

func main() {
	http.HandleFunc("/restaurants", getRestaurants)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	// Make request to the API
	resp, err := http.Get("https://uk.api.just-eat.io/discovery/uk/restaurants/enriched/bypostcode/EC4M7RF")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	var data struct {
		Restaurants []Restaurant `json:"restaurants"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Restaurants)
}
const express = require("express");
const axios = require("axios");
const cors = require("cors");
const app = express();
const port = 3000;

app.use(cors());

app.get("/", (req, res) => {
  res.send("Server working");
});

app.get("/restaurants/:postcode", async (req, res) => {
  try {
    const postcode = req.params.postcode;
    const response = await axios.get(
      `https://uk.api.just-eat.io/discovery/uk/restaurants/enriched/bypostcode/${postcode}`
    );
    const restaurants = response.data.restaurants
      .slice(0, 10)
      .map((restaurant) => ({
        name: restaurant.name,
        rating: restaurant.rating.starRating,
        cuisines: restaurant.cuisines.map((cuisine) => cuisine.name),
        address: restaurant.address.firstLine,
      }));
    res.json(restaurants);
  } catch (error) {
    console.log("Error:", error);
    res.status(500).json({ error: "Internal server error" });
  }
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
