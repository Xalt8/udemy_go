package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // should be a pointer
var mut sync.Mutex    // should be a pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://wikipedia.org",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) // Add 1 wait group
	}

	wg.Wait() // Used to wait for the added waitgroups to finish
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error with endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
