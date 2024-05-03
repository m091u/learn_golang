package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	//router defines and end point and associates with it the function to be executed
	// reference to the request is required
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", crud())
	return mux
}
