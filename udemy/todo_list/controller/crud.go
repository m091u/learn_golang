package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo_list/model"
	"todo_list/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)

			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

			// } else if r.Method == http.MethodGet {   //get omly method
			// 	data, err := model.ReadAll()
			// 	if err != nil {
			// 		w.Write([]byte(err.Error()))
			// 		return
			// 	}
			// 	w.WriteHeader(http.StatusOK)
			// 	json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet { //get & filter by query
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]

			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{"Item deleted"})
		}
	}
}
