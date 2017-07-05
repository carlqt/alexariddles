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

	if alexaReq.Type() == "IntentRequest" {
		alexaResp := intentRequestHandler(alexaReq)
		alexaResp.Respond(w, 200)
	} else {
		response.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true)
	}
}

func intentRequestHandler(alexaReq *alexaskill.AlexaRequest) *response.AlexaResponse {
	// if RepeatRiddle or AnswerRiddle is called without a session, should respond "No riddles yet"
	switch alexaReq.IntentName() {
	case "AMAZON.CancelIntent":
		return response.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true)
	case "AMAZON.StopIntent":
		return response.AlexaText("Questor stopped").SimpleCard("Questor", "stop").EndSession(true)
	case "AskRiddle":
		answer, riddle := riddles.Ask()
		return response.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer).
			EndSession(false)

	case "RepeatRiddle":
		answer := alexaReq.GetSessionAttr("answer")
		riddle := riddles.GetRiddle(answer)

		if len(answer) == 0 {
			return response.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				RepromptText("Time is up. The answer is, " + answer).
				EndSession(false)
		}

		return response.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer).
			EndSession(false)

	case "AnswerRiddle":
		sessionAnswer := alexaReq.GetSessionAttr("answer")
		userAnswer := alexaReq.GetUserAnswer()

		if sessionAnswer == userAnswer {
			return response.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).EndSession(true)
		}

		if len(userAnswer) == 0 {
			return response.AlexaText("Sorry, it is not the answer. Try again").
				SessionAttr("answer", sessionAnswer).
				RepromptText("Time is up. The answer is, "+sessionAnswer).
				SimpleCard("Riddle me this", "Sorry, it is not the answer. Try again").
				EndSession(false)

		}

		return response.AlexaText("Sorry, "+userAnswer+" is not the answer. Try again").
			SessionAttr("answer", sessionAnswer).
			SimpleCard("Riddle me this", "Sorry, "+userAnswer+" is not the answer. Try again").
			RepromptText("Time is up. The answer is, " + sessionAnswer).
			EndSession(false)
	default:
		return response.AlexaText("I do not know how to answer").SimpleCard("Riddle me this", "I do not know how to answer").EndSession(true)
	}
}

func repeatRiddleHandler() {

}
