package main

import (
	"fmt"
	"net/http"
	"time"
) 

func helloGoPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path{
	case "/":
		fmt.Fprint(w, "<h1>Hello, Go!</h1>")
	case "/about":
		fmt.Fprint(w, "This is the about page.")
	default:
		fmt.Fprint(w, "404 Page not found.")
	}
	
}

func timeout ( w http.ResponseWriter, r *http.Request){
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not timeout")
}

func main() {
    http.HandleFunc("/", helloGoPage)
	http.HandleFunc("/timeout", timeout)
    // http.ListenAndServe("", nil)

	// own server implementation
	server := http.Server{
		Addr: "",
		Handler: nil,
		ReadTimeout: 1000,
		WriteTimeout: 1000,
	}
	server.ListenAndServe()
}