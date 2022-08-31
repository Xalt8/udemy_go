package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Xalt8/udemy_go/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
