package mongo_service


import (
	"api-uptoo-jwt/core/mongo"
	"encoding/json"
	"log"
	"net/http"
	"api-uptoo-jwt/models"
	"reflect"
	"fmt"
	"strings"
)



func CreateItem(modelName string, r *http.Request) {
	session := mongo.Connection()
	Objet := DecodebyModelNameToInterface(modelName, r)

	err := mongo.Collection(session,modelName).Insert(Objet)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetItems(modelName string, Settings map[string]string,model *models.ModelBase) interface {}{
	session := mongo.Connection()
	Objet:= mongo.ChargeStructByModelName(modelName)

	TypeObjet := reflect.TypeOf(&Objet)
	Field, _ := TypeObjet.FieldByName(modelName)
	ValueField := reflect.ValueOf(&Field).Elem()
	ObjetFinal := ValueField.Field(2).Interface()
	bson := mongo.ObjectToBson(modelName,model)


	err := mongo.Collection(session,modelName).Find(bson).All(&ObjetFinal)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ObjetFinal
}

func GetItem(modelName string, model *models.ModelBase) interface{} {


	session := mongo.Connection()

	//Objet := mongo.ChargeStructByModelName(modelName)
	Collection := mongo.ChargeCollectionByModelName(modelName)

	fmt.Println(strings.ToUpper(modelName))
	fmt.Print(models.ModelBase{})


	mStruct, _ := reflect.TypeOf(models.ModelBase{}).FieldByName(strings.ToUpper(modelName))
	mInstance := reflect.New(mStruct.Type.Elem())

	bson := mongo.ObjectToBson(modelName,model)
	err := mongo.Collection(session,Collection).Find(bson).One(mInstance)
	if err != nil {
		log.Fatal(err.Error())
	}
	return mInstance
}

func GetItemById(modelName string, model *models.ModelBase) interface {}{
	session := mongo.Connection()
	Objet := mongo.ChargeStructByModelName(modelName)

	TypeObjet := reflect.TypeOf(&Objet)
	Field, _ := TypeObjet.FieldByName(modelName)
	ValueField := reflect.ValueOf(&Field).Elem()
	ObjetFinal := ValueField.Field(2).Interface()


	bson := mongo.ObjectToBson(modelName+"ById",model)
	err := mongo.Collection(session,modelName).Find(bson).One(&ObjetFinal)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ObjetFinal
}

func UpdateItem(modelName string, model *models.ModelBase) []byte{
	session := mongo.Connection()
	bson := mongo.ObjectToBson(modelName+"ById",model)
	err := mongo.Collection(session,modelName).UpdateId(bson,model)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}



func DecodebyModelNameToInterface(modelName string, r *http.Request) *models.ModelBase{
	requestDecode:= mongo.ChargeStructByModelName(modelName)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestDecode)

	return requestDecode
}