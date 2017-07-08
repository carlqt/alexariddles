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

	if alexaReq.AppID() != "amzn1.ask.skill.3aebac54-38a0-4dd3-9f17-4942972e4136" {
		w.WriteHeader(404)
		log.Println("Invalid application id")
		return
	}

	switch alexaReq.Type() {
	case "IntentRequest":
		alexaResp := intentRequestHandler(alexaReq)
		alexaResp.Respond(w, 200)
	case "LaunchRequest":
		response.AlexaText("Questor has been launched. Let the games begin").SimpleCard("Questor", "Questor has been launched. Let the games begin").Respond(w, 200)
	default:
		response.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true).Respond(w, 200)
	}
}

func intentRequestHandler(alexaReq *alexaskill.AlexaRequest) *response.AlexaResponse {
	switch alexaReq.IntentName() {
	case "AMAZON.CancelIntent":
		return response.AlexaText("Questor cancelled").SimpleCard("Questor", "cancel").EndSession(true)
	case "AMAZON.StopIntent":
		return response.AlexaText("Questor stopped").SimpleCard("Questor", "stop").EndSession(true)
	case "AMAZON.HelpIntent":
		text := `Questor will give you random riddles for you to answer. To start, use the phrase, tell me a questor riddle or i am ready for a riddle. Answer his riddles starting with the phrase, The answer is your answer. Questor will wait a couple of seconds, then if you are not able to answer his riddle, questor will give you the answer. You can give up by saying the phrase, I don't know or I give up. You can exit by saying close, goodbye or cancel. What can I help you with?`

		return response.AlexaText(text).
			SimpleCard("Questor", text)
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
				EndSession(true)
		}

		return response.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer).
			EndSession(false)
	case "DontKnow":
		sessionAnswer := alexaReq.GetSessionAttr("answer")

		if len(sessionAnswer) == 0 {
			return response.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		}

		return response.AlexaText("The answer is "+sessionAnswer).
			SimpleCard("Riddle me this", "Then answer is "+sessionAnswer).
			EndSession(true)

	case "AnswerRiddle":
		sessionAnswer := alexaReq.GetSessionAttr("answer")
		userAnswer := alexaReq.GetUserAnswer()

		if sessionAnswer == userAnswer {
			return response.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).EndSession(true)
		}

		if len(sessionAnswer) == 0 {
			return response.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
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
