package mongo


import (
	"log"
	"gopkg.in/mgo.v2/bson"
	"api-uptoo-jwt/models"
	"strings"
	"gopkg.in/mgo.v2"
)

type Config struct {
	ConnectString string
	Database string
}


func Init() (Config){

	config := Config {
		ConnectString: "localhost",
		Database: "uptoo",
	}
	return config
}

func Collection(session *mgo.Session,modelName string) (*mgo.Collection){
	config := Init()
	c := session.DB(config.Database).C(modelName)
	return c
}

func Connection() (*mgo.Session){
	config := Init()
	session, err := mgo.Dial(config.ConnectString)
	if err != nil {
		log.Fatal(err.Error())
	}
return session
}


func QueryBuilder(settings map[string]string) []byte{
	ReturnValue := ""
	for key,value := range settings {
		switch key {
			case key :
				if ReturnValue != "" {
					ReturnValue = ReturnValue +","+key+":"+value
				} else {
					ReturnValue = key+":"+value
				}
		}
	}

	bson,err  := bson.Marshal(strings.NewReader(ReturnValue))
	if err != nil {
		log.Fatal(err.Error())
	}
	return bson
}


func ChargeStructByModelName(modelName string)  (*models.ModelBase){
	Default := new(models.ModelBase)

	switch modelName {
		case "people" :  var people = new(models.ModelBase)
			return people
	}
	return Default
}

func ObjectToBson(settings string, model *models.ModelBase) bson.M {

	switch settings {
		case "people" :           return bson.M{"emailaddress":model.PEOPLE.EmailAddress}
		case "peopleById" :       return bson.M{"_id":model.PEOPLE.Id}
		case "peopleByLinkedin" : return bson.M{"linkedinid" : model.PEOPLE.Positions.Values[0].Company.LinkedinId}
	}
	return bson.M{}
}


func ChargePositionByModelName(modelName string) int{
	var position int

	switch modelName {
	case "people" : position = 0
	}

	return position
}

func ChargeCollectionByModelName(modelName string) string {
	switch modelName{
	case "people" : collection:=models.CollectionPeople
			return collection
	}
	return ""
}


func Json2Bson(i interface{}) (m *bson.M, err error){
	b, err := bson.Marshal(i)
	if err != nil {
		return
	}
	err = bson.Unmarshal(b, &m)
	return
}