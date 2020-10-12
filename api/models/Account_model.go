package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountModel struct {
	db *mongo.Client
}

// func (accountModel AccountModel) CheckUsernameAndPassword(username, password string) bool {
// 	var account User
// 	err := accountModel.db.Database("quickstart").Collection("users").Find(bson.M{
// 		"username": username,
// 	})
// }
