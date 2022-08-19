package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#754",
	}
	colors["white"] = "#fffff"

	// var colors map[string]string
	// colors := make(map[string]string)

	for key, value := range colors {
		fmt.Println(key, value)
	}

}
