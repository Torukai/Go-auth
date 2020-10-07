package main

import (
	api "torukai.net/auth/api"
)

func main() {

	api.Run()
	// var token = os.Getenv("TOKEN")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	// 	token,
	// ))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

}
