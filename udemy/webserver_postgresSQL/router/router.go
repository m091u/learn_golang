package router

import (
	"github.com/gorilla/mux" // | a reliable router
	"udemy/webserver_postgresSQL/services"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", services.CreatePosts).Methods("POST")
	router.HandleFunc("/posts/{id}", services.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", services.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", services.DeletePost).Methods("DELETE")

	return router
}
