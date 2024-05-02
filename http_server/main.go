package main

import (
	"net/http"
	"src/http_server/api"
)

func main() {

	srv := api.Newserver()
	http.ListenAndServe(":8080", srv)
}
