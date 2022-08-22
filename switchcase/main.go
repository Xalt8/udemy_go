package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1 //Generate a random number between 1 and 6 (inclusive)
	fmt.Println("The dice value is ", diceNumber)

	switch diceNumber {
	case 1:
		println("Dice number is 1")
	case 2:
		println("Dice number is 2")
	case 3:
		println("Dice number is 3")
		fallthrough
	case 4:
		println("Dice number is 4")
		fallthrough
	case 5:
		println("Dice number is 5")
	case 6:
		println("Dice number is 6")
	default:
		fmt.Println("Default case to catch outlier")
	}
}
