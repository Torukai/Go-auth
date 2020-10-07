package auto

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"torukai.net/auth/api/database"
)

func Load() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	databases, err := db.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	// defer db.Close()

	// err = db.Debug().DropTableIfExists(&models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().AutoMigrate(&models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, user := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&user).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	console.Pretty(user)
	// }
}
