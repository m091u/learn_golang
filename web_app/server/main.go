package main

import (
"fmt"
"net/http"
"encoding/json")

func handleFunc( w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
	case "/url":
		url(w, r)
		//http://localhost/url?name=Wallace
	case "/body":
		body(w, r)
		//http://localhost/body
		//with body in JSON: { "name": "Wallace"}
	default:
		w.Write([]byte("hello, Go API!"))
		//http://localhost
	}
}

func url( w http.ResponseWriter, r *http.Request){
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

type Person struct{
	Name string `json:"name"`
}

func body(w http.ResponseWriter, r *http.Request){
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name from JSON request body: %s\n", person.Name)

}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Server is running on default port...")
	http.ListenAndServe("", nil)
    
}