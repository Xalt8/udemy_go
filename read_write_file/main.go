package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: You need to specify a file")
	}

	// data, err := ioutil.ReadFile(os.Args[1])
	data, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: Something went wrong with reading a file", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, data)
	// fmt.Println(string(data))

}
