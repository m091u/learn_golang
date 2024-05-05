package model

type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Cuisine string `json:"cuisine"`
	PreparationTime int `json:"preparation_time"`
	Link string `json:"link"`
}