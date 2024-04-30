package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
    response, err := http.Get("http://localhost")

	if err != nil{
		fmt.Println("Server error", err)
	} else{
		data,_ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//URL key value form
	response, err = http.PostForm(
		"http://localhost/url", 
		url.Values{"name": {"Mira"}},
	)
	if err != nil{
		fmt.Println(err)
	} else {
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
	}

	//http://localhost/body
	//with body in json

	type Person struct {
		Name string
	}

	person := Person{Name: "John Doe"}
	personJson, _ := json.Marshal(person)


    client:= http.Client{}
	request, err := http.NewRequest(
		"GET",
		"http://localhost/body",
		bytes.NewBuffer(personJson),
	)
	
	if err != nil {
		fmt.Println("Error", err)
	}
		
	request.Header.Set("Content-Type", "application/json")
	response, err = client.Do(request)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	
}