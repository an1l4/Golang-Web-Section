package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response %T\n", response) //will get refernce (pointer) not a copy
	//fmt.Println(response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
