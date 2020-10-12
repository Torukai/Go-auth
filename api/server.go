package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"torukai.net/auth/api/controllers"
	"torukai.net/auth/api/router"
	"torukai.net/auth/auto"
	"torukai.net/auth/config"
)

var static string = "/static/"
var adminStatic string = "/admin/static/"

func Run() {
	config.Load()
	db := auto.Load()
	controllers.Init(db)
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)

	listen(config.PORT)
}

func listen(port int) {

	// http.Handle("/static/", //final url can be anything
	// 	http.StripPrefix("/static/",
	// 		http.FileServer(http.Dir("static"))))

	r := router.New()

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("templates/styles/"))))

	// serveStatic(r, static, "")
	// serveStatic(r, adminStatic, "/admin")

	// r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("/templates/styles"))))

	// _, err := os.Stat(filepath.Join(".", "templates/styles", "style.css"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func serveStatic(router *mux.Router, staticDir string, admin string) {
	staticPaths := map[string]string{
		"/css/":     staticDir + "/css/",
		"/test/":    staticDir + "/test/",
		"/images/":  staticDir + "/images/",
		"/scripts/": staticDir + "/scripts/",
		"/tinymce/": staticDir + "/scripts/tinymce/js/tinymce/",
		// If we use "/files/" as a prefix we get in conflict with the router which also use files.
		// Also it only works if the files folder is inside another folder also due to the conflict.
		"/uploads/": staticDir + "/uploads/",
	}
	for pathPrefix, pathValue := range staticPaths {
		router.PathPrefix(admin + pathPrefix).Handler(http.StripPrefix(admin+pathPrefix,
			http.FileServer(http.Dir("."+pathValue))))
	}
}
