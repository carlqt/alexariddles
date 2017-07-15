//TODO: Check if answer is correct/incorrect done
//TODO: Reprompt if user is taking too long to answer done
//TODO: Handle multiple answers in sessions
//TODO: RepeatIntent
//TODO: HelpIntent - instructions
//TODO: SLOTS struct improvement
//TODO: Handle when AnswerRiddle is initiated without AskRiddle
//TODO: Study alexa request and refactor session getters
//TODO: add unit test to alexaskill package
//TODO: Add tests when we can make requests again using the amazon
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/carlqt/alexariddles/alexaskill"
	"github.com/carlqt/alexariddles/helpers/httpdebug"
	"github.com/carlqt/alexariddles/riddles"
)

func ApiHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func RiddleHandler(w http.ResponseWriter, r *http.Request) {
	myAppID := "amzn1.ask.skill.3aebac54-38a0-4dd3-9f17-4942972e4136"
	// myAppID := "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"

	alexaReq, err := alexaskill.AlexaNewRequest(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
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
		alexaskill.NewAlexaText("Riddley has been launched. Let the games begin").SimpleCard("Riddley", "Riddley has been launched. Let the games begin").Respond(w, 200)
	default:
		alexaskill.NewAlexaText("There was something wrong").SimpleCard("Riddley", "something occured").EndSession(true).Respond(w, 200)
	}
}

func intentRequestResponse(alexaReq *alexaskill.AlexaRequest) *alexaskill.AlexaResponse {
	alexaResp := alexaskill.NewAlexaResponse("1.0")

	switch alexaReq.IntentName() {
	case "AMAZON.CancelIntent":
		alexaResp.AlexaText("Riddley cancelled").SimpleCard("Riddley", "cancel").EndSession(true)
	case "AMAZON.StopIntent":
		alexaResp.AlexaText("Riddley stopped").SimpleCard("Riddley", "stop").EndSession(true)
	case "AMAZON.HelpIntent":
		text := `Riddley will give you random riddles for you to answer. To start, use the phrase, Alexa ask riddley for riddles. To Answer his riddles, start your phrase with, Alexa, tell riddley the answer is your answer. Riddley will wait a couple of seconds, then if you're not able to answer his riddle, riddley will give you the answer. You can give up by saying the phrase, Alexa tell riddley I give up. You can exit by saying close, goodbye or cancel. You can repeat this by simpley saying Help. Is there anything else?`

		alexaResp.AlexaText(text).
			SimpleCard("Riddley", text)
	case "AskRiddle":
		answer, riddle := riddles.Ask()
		alexaResp.AlexaText(riddle).
			SimpleCard("Riddle me this", riddle).
			SessionAttr("answer", answer).
			RepromptText("Time is up. The answer is, " + answer + ". Would you like another riddle?").
			EndSession(false)

	case "RepeatRiddle":
		answer := strings.ToLower(alexaReq.GetSessionAttr("answer"))
		riddle := riddles.GetRiddle(answer)

		if len(answer) == 0 {
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		} else {
			alexaResp.AlexaText(riddle).
				SimpleCard("Riddle me this", riddle).
				SessionAttr("answer", answer).
				RepromptText("Time is up. The answer is, " + answer + ". Would you like another riddle?").
				EndSession(false)

		}
	case "DontKnow":
		sessionAnswer := strings.ToLower(alexaReq.GetSessionAttr("answer"))

		if len(sessionAnswer) == 0 {
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddle me this", "No riddles have been given yet").
				EndSession(true)
		} else {
			alexaResp.AlexaText("The answer is "+sessionAnswer).
				SimpleCard("Riddley", "Then answer is "+sessionAnswer).
				EndSession(true)
		}

	case "AnswerRiddle":
		sessionAnswer := strings.ToLower(alexaReq.GetSessionAttr("answer"))
		userAnswer := strings.ToLower(alexaReq.Request.Intent.Slots.Value("RiddleAnswer"))

		switch {
		case len(sessionAnswer) == 0:
			alexaResp.AlexaText("No riddles have been given yet").
				SimpleCard("Riddley", "No riddles have been given yet").
				EndSession(true)
		case len(userAnswer) == 0:
			alexaResp.AlexaText("Sorry, it is not the answer. Try again").
				SessionAttr("answer", sessionAnswer).
				RepromptText("Time is up. The answer is, "+sessionAnswer+". Would you like another riddle?").
				SimpleCard("Riddle me this", "Sorry, it is not the answer. Try again").
				EndSession(false)
		case sessionAnswer == userAnswer:
			alexaResp.AlexaText("You are correct. The answer is "+sessionAnswer).SimpleCard("Riddle me this", "You are correct. Then answer is "+sessionAnswer).EndSession(true)
		default:
			alexaResp.AlexaText("Sorry, "+userAnswer+" is not the answer. Try again").
				SessionAttr("answer", sessionAnswer).
				SimpleCard("Riddley", "Sorry, "+userAnswer+" is not the answer. Try again").
				RepromptText("Time is up. The answer is, " + sessionAnswer + ". Would you like another riddle?").
				EndSession(false)
		}
	default:
		alexaResp.AlexaText("I do not know how to answer").SimpleCard("Riddle me this", "I do not know how to answer").EndSession(true)
	}

	return alexaResp
}

func logRequest(r *http.Request) {
	requestCopy, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(requestCopy))
	httpdebug.PrettyJson(bytes.NewReader(requestCopy))
}
