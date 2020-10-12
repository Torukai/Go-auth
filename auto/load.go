package auto

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"torukai.net/auth/api/database"
)

var db *mongo.Client

// func GetDBCollection(db *mongo.Client) (*mongo.Collection, error) {

// 	// Check the connection
// 	// err := db.Ping(context.Background(), nil)
// 	// if err != nil {
// 	// 	fmt.Println("here 2")
// 	// 	return nil, err
// 	// }
// 	// collection := db.Database("myDB").Collection("users")
// 	// fmt.Println("here")
// 	// return collection, nil

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	db, err := database.Connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	collection := db.Database("myDB").Collection("users")
// 	return collection, nil
// }

func Load() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.Connect(ctx)
	//controllers.Init(db)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Disconnect(ctx)

	return db
}
