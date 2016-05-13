package controllers

import (
	"net/http"
	"api-uptoo-jwt/services/mongo"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	mongo_service.CreateItem("people",r)
}
