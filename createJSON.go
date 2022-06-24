package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias
	Price    int
	Platform string   `json:"website"`        //alias
	Password string   `json:"-"`              //this "-" dash inside double quotes will not reflect whoever is consuming my API
	Tags     []string `json:"tags,omitempty"` //data type is slice of string
}

func main() {
	fmt.Println("Creating JSON data")
	EncodingJson()
}

//encoding Json- converting different type of data to json format
func EncodingJson() {
	lcoCourses := []course{ //slice of type course (struct)
		{"ReacrJS", 299, "learncodeonline.in", "abc123", []string{"web-dev", "JS"}},
		{"NodeJS", 399, "Learncodeonline.in", "xyz123", []string{"full-stack", "JS"}},
		{"Golang", 599, "learncodeonline.in", "efg123", nil},
	}

	//package this data as JSON data

	//finalJSONData, err := json.Marshal(lcoCourses) //if we use Marshal we will get data but difficult to read
	//using MarshalIndent to get more readable data
	finalJSONData, err := json.MarshalIndent(lcoCourses, "", "\t") //indent based on new tab \t
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJSONData)
}
