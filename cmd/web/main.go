package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjayashri/go-practice/pkg/config"
	"github.com/mjayashri/go-practice/pkg/handlers"
	"github.com/mjayashri/go-practice/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port number %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
