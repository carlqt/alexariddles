package main

import (
	"log"
	"net/http"
	"os"

	alexaMiddleware "github.com/carlqt/alexariddles/alexaskill/middleware"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func init() {
	if !exists("log") {
		os.MkdirAll("log", os.ModePerm)
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(ApiHandler)
	r.Use(alexaMiddleware.AlexaValidation)
	r.Use(logRequest)

	r.Post("/", RiddleHandler)

	log.Println("listening to port ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}
