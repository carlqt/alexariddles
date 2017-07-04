//TODO: Check if answer is correct/incorrect
//TODO: Reprompt if user is taking too long to answer
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
			sessionAnswer := alexaReq.GetSessionAttr("answer")
			userAnswer := alexaReq.GetUserAnswer()

			if sessionAnswer == userAnswer {
				response.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).Respond(w, 200, true)
			} else {
				if userAnswer == "" {
					response.AlexaText("Sorry, it is not the answer. Try again").SessionAttr("answer", sessionAnswer).SimpleCard("Riddle me this", "Sorry, it is not the answer. Try again").Respond(w, 200, false)
				} else {
					response.AlexaText("Sorry, "+userAnswer+" is not the answer. Try again").SessionAttr("answer", sessionAnswer).SimpleCard("Riddle me this", "Sorry, "+userAnswer+" is not the answer. Try again").Respond(w, 200, false)
				}
			}
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
