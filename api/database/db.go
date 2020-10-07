package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"torukai.net/auth/config"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	//db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	//var token = os.Getenv("TOKEN")
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(
		config.DBURL,
	))
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Disconnect(ctx)
	return db, nil
}
