package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("POST Request")
	PerformPostJSONRequest()
}

func PerformPostJSONRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	//creating some data in json format this data needs to go to server

	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go with Go",
		"price":10,
		"Platform":"w3-schools"
	}
	`)
	//Post request
	response, err := http.Post(myurl, "application/json", requestBody)
	//handling error
	if err != nil {
		panic(err)
	}
	//closing the connection after the request
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
