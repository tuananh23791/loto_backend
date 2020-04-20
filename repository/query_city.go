package repository

import (
	"context"
	"fmt"
	"log"
	ErrorCode "travel/config"
	"travel/model"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionCity *mongo.Collection

func InsertCity(cityList []model.City, tokenString string) (int, string) {
	getCollectionCity()

	errorCode, errorMessage := validateToken(tokenString)

	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage
	}
	fmt.Println("cityList -- ", cityList)

	ui := []interface{}{}
	for _, v := range cityList {
		ui = append(ui, structs.Map(v))
	}
	fmt.Println("ui -- ", ui)
	_, err := collectionCity.InsertMany(context.TODO(), ui)

	if err != nil {
		fmt.Println("InsertMany error" + err.Error())
		return ErrorCode.SOME_THING_WENT_WRONG, "Insert City error - " + err.Error()
	}

	return ErrorCode.SUCCESS, ""
}

func GetListCity(tokenString string) (int, string, []model.City) {
	getCollectionCity()
	var results []model.City

	errorCode, errorMessage := validateToken(tokenString)

	if errorCode != ErrorCode.SUCCESS {
		return errorCode, errorMessage, results
	}

	cur, err := collectionCity.Find(context.TODO(), bson.D{})
	if err != nil {
		return ErrorCode.SOME_THING_WENT_WRONG, "GetListCity Find error - " + err.Error(), results
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var city model.City
		err := cur.Decode(&city)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, city)
	}

	return ErrorCode.SUCCESS, "", results
}

func getCollectionCity() {
	if collectionCity == nil {
		Connect()
		collectionCity = db.Collection("city")
	}
	// ctx, _ = context.WithTimeout(context.Background(), config.CONNECTION_TIME_OUT*time.Second)
}
