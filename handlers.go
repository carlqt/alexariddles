//TODO: Check if answer is correct/incorrect done
//TODO: Reprompt if user is taking too long to answer done
//TODO: Handle multiple answers in sessions
//TODO: RepeatIntent
//TODO: HelpIntent - instructions
//TODO: SLOTS struct improvement
//TODO: Handle when AnswerRiddle is initiated without AskRiddle
package main

import (
	"net/http"

	"github.com/carlqt/alexariddles/alexaskill"
	"github.com/carlqt/alexariddles/alexaskill/response"
	"github.com/carlqt/alexariddles/riddles"
)

func ApiHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func riddleHandler(w http.ResponseWriter, r *http.Request) {

	alexaReq, err := alexaskill.AlexaNewRequest(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	answer, riddle := riddles.Ask()
	if alexaReq.Type() == "IntentRequest" {
		switch alexaReq.IntentName() {
		case "AMAZON.CancelIntent":
			response.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").Respond(w, 200, true)
		case "AMAZON.StopIntent":
			response.AlexaText("Questor stopped").SimpleCard("Questor", "stop").Respond(w, 200, true)
		case "AskRiddle":
			response.AlexaText(riddle).
				SimpleCard("Riddle me this", riddle).
				SessionAttr("answer", answer).
				RepromptText("Time is up. The answer is, "+answer).
				Respond(w, 200, false)

		case "RepeatRiddle":
			response.AlexaText(riddle).
				SimpleCard("Riddle me this", riddle).
				SessionAttr("answer", answer).
				RepromptText("Time is up. The answer is, "+answer).
				Respond(w, 200, false)

		case "AnswerRiddle":
			sessionAnswer := alexaReq.GetSessionAttr("answer")
			userAnswer := alexaReq.GetUserAnswer()

			if sessionAnswer == userAnswer {
				response.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).Respond(w, 200, true)
			} else {
				if len(userAnswer) == 0 {
					response.AlexaText("Sorry, it is not the answer. Try again").
						SessionAttr("answer", sessionAnswer).
						RepromptText("Time is up. The answer is, "+answer).
						SimpleCard("Riddle me this", "Sorry, it is not the answer. Try again").
						Respond(w, 200, false)

				} else {
					response.AlexaText("Sorry, "+userAnswer+" is not the answer. Try again").
						SessionAttr("answer", sessionAnswer).
						SimpleCard("Riddle me this", "Sorry, "+userAnswer+" is not the answer. Try again").
						RepromptText("Time is up. The answer is, "+answer).
						Respond(w, 200, false)
				}
			}
		default:
			response.AlexaText("I do not know how to answer").SimpleCard("Riddle me this", "I do not know how to answer").Respond(w, 200, true)
		}
	}
}
