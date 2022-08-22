package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("Error: You need to specify a file")
	// }
	// // data, err := ioutil.ReadFile(os.Args[1])
	// data, err := os.Open(os.Args[1])

	// if err != nil {
	// 	fmt.Println("Error: Something went wrong with reading a file", err)
	// 	os.Exit(1)
	// }
	// io.Copy(os.Stdout, data)
	// fmt.Println(string(data))

	// Try writing to a file
	content := "This needs to go into a file."
	file, err := os.Create("./createfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./createfile.txt")
}

func readFile(filename string) {
	// When you read a file its in bytes format
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("File contents: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
