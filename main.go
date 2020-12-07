package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// Call homePage when we go to the / path of our app
	http.HandleFunc("/", homePage)
	// Listen on port 10000
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
