//TODO: Handle when user has incorrect answers
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	alexaMiddleware "github.com/carlqt/alexaskill/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var f *os.File

func init() {
	var err error
	f, err = os.OpenFile("./riddles.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(f)
}

func main() {
	port := os.Getenv("PORT")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(ApiHandler)
	r.Use(alexaMiddleware.AlexaValidation)

	r.Post("/", riddleHandler)

	log.Println("listening to port ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
