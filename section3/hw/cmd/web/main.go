package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/simonwardle/goWeb/pkg/config"
	"github.com/simonwardle/goWeb/pkg/handlers"
	"github.com/simonwardle/goWeb/pkg/render"
)

const portNumber = ":8080"



// main is the main application starting point
func main() {
	var app config.AppConfig
	var err error

	//populate app config template cache via render package
	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error: cannot load template cache!")
	}

	//pass app config populated with template cache to render.
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
