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

func RiddleHandler(w http.ResponseWriter, r *http.Request) {
	//myAppID := "amzn1.ask.skill.3aebac54-38a0-4dd3-9f17-4942972e4136"
	myAppID := "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"

	alexaReq, err := alexaskill.AlexaNewRequest(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if alexaReq.AppID() != myAppID {
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
		text := `Questor will give you random riddles for you to answer. To start, use the phrase, tell me a questor riddle or i am ready for a riddle. Answer his riddles starting with the phrase, The answer is your answer. Questor will wait a couple of seconds, then if you are not able to answer his riddle, questor will give you the answer. You can give up by saying the phrase, I don't know or I give up. You can exit by saying close, goodbye or cancel. What can I help you with?`

		alexaResp.AlexaText(text).
			SimpleCard("Questor", text)
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
	case "DontKnow":
		sessionAnswer := alexaReq.GetSessionAttr("answer")

		if len(sessionAnswer) == 0 {
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		}

		alexaResp.AlexaText("The answer is "+sessionAnswer).
			SimpleCard("Riddle me this", "Then answer is "+sessionAnswer).
			EndSession(true)

	case "AnswerRiddle":
		sessionAnswer := alexaReq.GetSessionAttr("answer")
		userAnswer := alexaReq.Request.Intent.Slots.Value("RiddleAnswer")

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

func HeartBeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
