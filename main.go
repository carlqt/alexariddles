package main

import (
	"log"
	"net/http"
	"os"

	alexaMiddleware "github.com/carlqt/alexariddles/alexaskill/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func init() {
	if !exists("log") {
		os.MkdirAll("log", os.ModePerm)
	}
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(ApiHandler)
	r.Use(alexaMiddleware.AlexaValidation)

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
