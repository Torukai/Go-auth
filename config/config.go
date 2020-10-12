package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT     = 0
	DBDRIVER = ""
	DBURL    = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		//log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.w6unm.mongodb.net/%s?retryWrites=true&w=majority", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

}
