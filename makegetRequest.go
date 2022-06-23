package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET Request")
	PerfomGetRequest()
}

func PerfomGetRequest() {
	const url = "http://localhost:8000/get"

	response, erro := http.Get(url)

	if erro != nil {
		panic(erro)
	}
	//fmt.Println(response)

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("content length :", response.ContentLength)

	//content, errors := ioutil.ReadAll(response.Body)
	//fmt.Println(content) //this line will give byte format we need to convert that using line 36

	//if errors != nil {
	//	panic(errors)
	//}

	//fmt.Println(string(content))

	//new way to translate this string
	//using strings package

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is :", byteCount) //here only will get the byte count

	//we need to convert to actual string
	//responseString is holding the actual data but need to convert

	fmt.Println(responseString.String())
}
