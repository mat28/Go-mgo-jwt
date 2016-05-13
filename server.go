package main

import (
	"api-uptoo-jwt/routers"
	"api-uptoo-jwt/settings"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE","PUT", "OPTIONS","PATCH"},
	})

	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	http.ListenAndServe(":5000",n)
}

