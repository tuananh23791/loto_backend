package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx context.Context

func Connect() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.HOST + ":" + config.PORT + ""))
	// if err != nil {

	// }
	// err = client.Connect(ctx)
	// db = client.Database("travel")
	if db != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// opts := options.ClientOptions{}
	// opts.SetDirect(true)
	// opts.ApplyURI("mongodb+srv://tony:Tony!123@cluster0-3tfep.gcp.mongodb.net/test?retryWrites=true&w=majority")

	// client, err := mongo.Connect(context.TODO(), &opts)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://tony:Tony!123@cluster0-3tfep.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	db = client.Database("travel")

	// collection := client.Database("travel").Collection("city")
	// ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	// res, err := collection.InsertOne(ctx, bson.M{"name": "HCM"})
	// // fmt.Println("ID:::::" + id)
	// if err != nil {
	// 	fmt.Println("error" + err.Error())
	// }
	// fmt.Println("Inserted a single document: ", res.InsertedID)
}
