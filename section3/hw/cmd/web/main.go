package main

import (
	"fmt"
	"net/http"

	"github.com/simonwardle/goWeb/pkg/handlers"
)

const portNumber = ":8080"



// main is the main application starting point
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
