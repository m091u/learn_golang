package router

import (

	"github.com/gorilla/mux" // | a reliable router
	"github.com/jagottsicher/myGoWebserver/services"
)

func CreateRouter() *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	
	return router
}