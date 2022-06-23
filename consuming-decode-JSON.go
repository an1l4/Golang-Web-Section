package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Decode Json Data")
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "NodeJS",
		"Price": 399,
		"website": "Learncodeonline.in",
		"tags": ["full-stack","JS"]
}
		
	`)
	//checking or varifiying our given json is valid format or not

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) //it will give u bool value

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse) //interface
	} else {
		fmt.Println("json is not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{} //key is string we dont know getting values type.it could be numbers,could be string so we mention value type as interface
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v and Type is %T \n", key, value, value)
	}
}
