package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo_list/views"
)

func ping() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request was received")

		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
