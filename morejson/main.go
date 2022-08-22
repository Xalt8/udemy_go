package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("=== JSON ====")
	// encodeJSON()
	decodeJson()

}

func encodeJSON() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncoding.com", "abc123", []string{"web", "js"}},
		{"Mern Bootcamp", 199, "Learncoding.com", "bcd123", []string{"full stack", "js"}},
		{"Angular Bootcamp", 299, "Learncoding.com", "hit123", nil},
	}

	// package this as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "Learncoding.com",   
		"Password": "abc123",
		"Tags": ["web","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON invalid")
	}

	// Some cases you want data in key,value pairs

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v, value: %v, type: %T\n", k, v, v)
	}
}
