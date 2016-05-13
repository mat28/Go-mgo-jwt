package controllers

import (
	"api-uptoo-jwt/services/auth"
	"api-uptoo-jwt/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	responseStatus, token := authentication_service.Login(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(authentication_service.RefreshToken(r))
}


func Signup(w http.ResponseWriter, r *http.Request){

	responseStatusCreate,_ := authentication_service.SignUp(r)
	responseStatusCreate = 200
	if responseStatusCreate == 200 {
		responseStatus, token := authentication_service.Login(r)
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(responseStatus)
		w.Write(token)
	}else{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(responseStatusCreate)
	}

}
/*
func Activation(w http.ResponseWriter, r *http.Request){

	responseStatus, token := authentication_service.VerifyActivation(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)


}
*/
func Init(w http.ResponseWriter, r *http.Request){

	responseStatus, token := authentication_service.Init(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	requestUser := models.People{}
	decoder :=json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	//API Mandrill
}
