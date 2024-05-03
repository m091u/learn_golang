package main

import (
	"fmt"
	"net/http"
	"todo_list/controller"
	"todo_list/model"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mux := controller.Register()

	db, err := model.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Listeneing to port ")
	log.Fatal(http.ListenAndServe(":3000", mux)) //created a server listening on a port and past in a multiplexer/ router
}
