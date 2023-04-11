package main

import (
	"errors"
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

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request){
	var x, y float32
	x = 100.0
	y = 0.0
	sum, err := divideValues(x, y)
	if err != nil {
		_, _ = fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}
	
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, sum))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}