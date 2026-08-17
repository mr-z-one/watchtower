package db

import (
	"context"
	"time"
	custom_color "watchtower/custom_Color"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var current_db *mongo.Database

var DBNAME string = "WatchTower"

func Mongo_connect() *mongo.Database {
	if current_db != nil {
		return current_db
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		custom_color.Error()("\n%v\n", err)
		panic(err)
	}

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		custom_color.Error()("\n%v\n", err)
		panic(err)
	}

	// Save and return
	current_db = client.Database(DBNAME)
	custom_color.Succeed()("[+] Connect to mongo is Successful\n")
	return current_db
}
