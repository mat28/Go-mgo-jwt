package routers

import (
	"api-uptoo-jwt/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("POST")

	return router
}
