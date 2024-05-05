package model

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Platform string `json:"platform"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}