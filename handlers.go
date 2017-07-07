//TODO: Check if answer is correct/incorrect done
//TODO: Reprompt if user is taking too long to answer done
//TODO: Handle multiple answers in sessions
//TODO: RepeatIntent
//TODO: HelpIntent - instructions
//TODO: SLOTS struct improvement
//TODO: Handle when AnswerRiddle is initiated without AskRiddle
package main

import (
	"log"
	"net/http"

	"github.com/carlqt/alexariddles/alexaskill"
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

	if alexaReq.AppID() != "amzn1.ask.skill.3aebac54-38a0-4dd3-9f17-4942972e4136" {
		w.WriteHeader(404)
		log.Println("Invalid application id")
		return
	}

	switch alexaReq.Type() {
	case "IntentRequest":
		intentRequestResponse(alexaReq).Respond(w, 200)
	case "LaunchRequest":
		alexaskill.NewAlexaText("Questor has been launched. Let the games begin").SimpleCard("Questor", "Questor has been launched. Let the games begin").Respond(w, 200)
	default:
		alexaskill.NewAlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true).Respond(w, 200)
	}
}

func intentRequestResponse(alexaReq *alexaskill.AlexaRequest) *alexaskill.AlexaResponse {
	alexaResp := alexaskill.NewAlexaResponse("1.0")

	switch alexaReq.IntentName() {
	case "AMAZON.CancelIntent":
		alexaResp.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true)
	case "AMAZON.StopIntent":
		alexaResp.AlexaText("Questor stopped").SimpleCard("Questor", "stop").EndSession(true)
	case "AMAZON.HelpIntent":
		alexaResp.AlexaText("Answer the riddles with the phrase, The answer is my answer").
			SimpleCard("Questor", "Answer the riddles with the phrase, the answer is <my answer>").
			EndSession(true)
	case "AskRiddle":
		answer, riddle := riddles.Ask()
		alexaResp.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer).
			EndSession(false)

	case "RepeatRiddle":
		answer := alexaReq.GetSessionAttr("answer")
		riddle := riddles.GetRiddle(answer)

		if len(answer) == 0 {
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		}

		alexaResp.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer).
			EndSession(false)

	case "AnswerRiddle":
		sessionAnswer := alexaReq.GetSessionAttr("answer")
		userAnswer := alexaReq.GetUserAnswer()

		if sessionAnswer == userAnswer {
			alexaResp.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).EndSession(true)
		}

		if len(sessionAnswer) == 0 {
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		}

		if len(userAnswer) == 0 {
			alexaResp.AlexaText("Sorry, it is not the answer. Try again").
				SessionAttr("answer", sessionAnswer).
				RepromptText("Time is up. The answer is, "+sessionAnswer).
				SimpleCard("Riddle me this", "Sorry, it is not the answer. Try again").
				EndSession(false)

		}

		alexaResp.AlexaText("Sorry, "+userAnswer+" is not the answer. Try again").
			SessionAttr("answer", sessionAnswer).
			SimpleCard("Riddle me this", "Sorry, "+userAnswer+" is not the answer. Try again").
			RepromptText("Time is up. The answer is, " + sessionAnswer).
			EndSession(false)
	default:
		alexaResp.AlexaText("I do not know how to answer").SimpleCard("Riddle me this", "I do not know how to answer").EndSession(true)
	}

	return alexaResp
}

func repeatRiddleHandler() {

}
