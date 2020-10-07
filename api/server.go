package api

import (
	"fmt"
	"log"
	"net/http"

	"torukai.net/auth/api/router"
	"torukai.net/auth/auto"
	"torukai.net/auth/config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
