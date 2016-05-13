package authentication_service

import (
	"api-uptoo-jwt/api/parameters"
	"api-uptoo-jwt/core/authentication"
	"encoding/json"
	"net/http"
	"api-uptoo-jwt/services/mongo"
	"api-uptoo-jwt/models"
	"log"
)

func Login(r *http.Request) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	requestUser := new(models.ModelBase)

	requestUser = mongo_service.DecodebyModelNameToInterface("people",r)

	log.Print(requestUser)

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.PEOPLE)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func Init(r *http.Request) (int, []byte){
	authBackend := authentication.InitJWTAuthenticationBackend()

	requestUser := new(models.ModelBase)

	requestUser = mongo_service.DecodebyModelNameToInterface("people",r)

	if authBackend.Init(requestUser.PEOPLE) {
		token, err := authBackend.GenerateTokenGuest(requestUser.PEOPLE.EmailAddress)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		}else{
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}


func RefreshToken(r *http.Request) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	requestUser := new(models.ModelBase)

	requestUser = mongo_service.DecodebyModelNameToInterface("people",r)

	token, err := authBackend.GenerateToken(requestUser.PEOPLE)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

/*func VerifyActivation(r *http.Request) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	var Settings map[string]string
	User := new(models.PEOPLE)
	Settings["tokenActivation"] = requestToken.Token
	if authBackend.VerifyActivation(requestToken){
		result := mongo_service.GetItem("people",Settings)
		decode := json.NewDecoder(bytes.NewReader(result))
		decode.Decode(&User)
		token, err := authBackend.GenerateToken(User.EmailAddress)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		}else{
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}*/

func SignUp(r *http.Request) (int , *models.ModelBase){


	Objet := new(models.ModelBase)

	Objet = mongo_service.DecodebyModelNameToInterface("people",r)
	mongo_service.GetItem("people",Objet)



	return 200,nil
}

