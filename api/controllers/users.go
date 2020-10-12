package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"torukai.net/auth/api/database"
	"torukai.net/auth/api/models"
	"torukai.net/auth/helpers"
)

var db *mongo.Client

func Init(_client *mongo.Client) {
	db = _client
}

func DefaultHandler(response http.ResponseWriter, request *http.Request) {

	var body, _ = helpers.LoadFile("templates/index.html")
	fmt.Fprintf(response, body)
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user models.User
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &user)
	var res models.ResponseResult

	fmt.Printf("1" + user.Username)
	fmt.Printf("2" + user.FirstName)
	fmt.Printf("3" + user.LastName)

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(response).Encode(res)
		return

	}

	collection, err := database.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		fmt.Println("here")
		json.NewEncoder(response).Encode(res)
		return
	}

	var result models.User
	err = collection.FindOne(context.Background(), bson.D{
		{"username", user.Username},
	}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(response).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)

			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(response).Encode(res)
				return
			}

			res.Result = "Registration Successful"
			json.NewEncoder(response).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(response).Encode(res)
		return
	}
	res.Result = "Username already Exists!"
	json.NewEncoder(response).Encode(res)
	return
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	var user models.User
	body, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := database.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}

	var result models.User
	var res models.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{
		{"username", user.Username},
	}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(response).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"firstname": result.FirstName,
		"lastname":  result.LastName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(response).Encode(res)
		return
	}

	result.Token = tokenString
	result.Password = ""

	json.NewEncoder(response).Encode(result)

}

func ProfileHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	tokenString := request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result models.User
	var res models.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		json.NewEncoder(response).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(response).Encode(res)
		return
	}

}

func GetUsers(response http.ResponseWriter, request *http.Request) {
	//response.Write([]byte("List users"))
	response.Header().Add("content-type", "application-json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := db.Database("myDB").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)

}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("A user"))

}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Create a user"))

}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Update user"))

}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Delete user"))

}
