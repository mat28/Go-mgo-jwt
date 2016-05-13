package routers


import (
	"api-uptoo-jwt/controllers/ws"
	"github.com/gorilla/mux"
)


func SetwebSocketRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/ws", ws.Echo)
	return router
}
