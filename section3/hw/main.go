package main

import (
	"fmt"
	"net/http"
)
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page!")
}

// Abbout is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(3, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 3 + 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}