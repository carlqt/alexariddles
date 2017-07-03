package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/carlqt/alexariddles/riddles"
	"github.com/carlqt/alexaskill"
	"github.com/carlqt/alexaskill/response"
)

func ApiHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// logFile(r)
		next.ServeHTTP(w, r)
	})
}

func riddleHandler(w http.ResponseWriter, r *http.Request) {
	logFile(r)

	alexaReq, err := alexaskill.AlexaNewRequest(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	riddle, answer := riddles.Ask()
	if alexaReq.Type() == "IntentRequest" {
		switch alexaReq.IntentName() {
		case "AMAZON.CancelIntent":
			response.AlexaText("Cancelled").SimpleCard("Cancel", "cancel").Respond(w, 200, true)
		case "AskRiddle":
			response.AlexaText(riddle).SimpleCard("Riddle me this", riddle).SessionAttr("answer", answer).Respond(w, 200, false)
		case "AnswerRiddle":
			response.AlexaText("perhaps").SimpleCard("Riddle me this", "perhaps").Respond(w, 200, true)
		default:
			response.AlexaText("I do not know how to answer").SimpleCard("Riddle me this", "I do not know how to answer").Respond(w, 200, true)
		}
	}
}

func logFile(r *http.Request) {
	requestCopy, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(requestCopy))

	logrus.WithFields(logrus.Fields{
		"request": string(requestCopy),
	}).Info("request info")

}
