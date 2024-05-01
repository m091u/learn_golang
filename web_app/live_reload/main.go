package main

import (
	"net/http"
	"text/template"
)

func handleFunc(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Hello again, Go! It's me.")
	var fileName = "helloWorld.html"
	t, _ := template.ParseFiles(fileName)
	t.ExecuteTemplate(w, fileName, nil)
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3001", nil)
}
