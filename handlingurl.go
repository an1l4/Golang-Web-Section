package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjhjjjnm"

func main() {
	fmt.Println(myurl)

	//parsing url

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) //Port() is a method
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type is %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams)

	//iterating
	for _, value := range qparams {
		fmt.Println("param is :", value)
	}

	//constructing url
	partsOfUrl := &url.URL{ //import thing is always passing the refernce not the copy (&)
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=jhfdej",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
