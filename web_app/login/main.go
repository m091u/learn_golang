package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var userDb = map[string]string{
	"Wallace": "goodPass",
}

type Person struct {
	Name string
	Age int
}

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, Person{"Mira", 39})
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username:", username, password)

	if userDb[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login success")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Login failed")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "/function":
		fileName := "function.html"
		funcMap := map[string]interface{}{
			"upper": strings.ToUpper,		
		}
		t, err := template.New(fileName).Funcs(funcMap).ParseFiles(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = t.ExecuteTemplate(w, fileName, "Hello People!")
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Fprintf(w, "Hello Go")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on default port...")
	http.ListenAndServe("", nil)
}
