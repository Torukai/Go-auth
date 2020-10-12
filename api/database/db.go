package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"torukai.net/auth/config"
)

func Connect(ctx context.Context) (*mongo.Client, error) {

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.DBURL,
	))
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Disconnect(ctx)
	return db, nil
}

func GetDBCollection() (*mongo.Collection, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.DBURL,
	))
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	//defer db.Disconnect(ctx)
	collection := db.Database("myDB").Collection("users")
	return collection, nil
}

// func GetDBCollection (){

// 	return db, nil
// }
