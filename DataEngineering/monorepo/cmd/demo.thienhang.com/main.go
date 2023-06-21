package main

import (
	"net/http"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	demoController "api_thienhang_com/services/demo/controllers"
	log "github.com/sirupsen/logrus"
)

func init() {
	runtime.GOMAXPROCS(1)
	err := godotenv.Load("config.env")
	if err != nil {
		log.Error("Error loading .env file")
	}
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	m := mux.NewRouter()
	
	controller := demoController.New(m)
	controller.ServeHTTP()

	router := handlers.CORS(handlers.AllowedHeaders(
		[]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(m)
	log.Fatal(http.ListenAndServe(":9000", router))

}
