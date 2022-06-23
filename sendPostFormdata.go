package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("POST Form")
	PerfomPostFormRequest()
}

func PerfomPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data

	data := url.Values{} //form data coming through url
	data.Add("First Name", "Anila")
	data.Add("Last Name", "Soman")
	data.Add("Email", "example@yahoo.com")

	//going to make request

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
