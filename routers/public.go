package routers


import (
	"api-uptoo-jwt/controllers/auth"
	"github.com/gorilla/mux"
)


func SetPublicRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/public/init", controllers.Init).Methods("POST")

	router.HandleFunc("/v1/public/login", controllers.Login).Methods("POST")
	router.HandleFunc("/v1/public/signin", controllers.Login).Methods("POST")
	router.HandleFunc("/v1/public/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/v1/public/resetPassword",controllers.ResetPassword).Methods("POST")
	//router.HandleFunc("/v1/public/activation", controllers.Activation).Methods("POST")

	return router
}

