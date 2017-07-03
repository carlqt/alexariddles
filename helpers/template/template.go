package template

import (
	"encoding/json"
	"log"
	"net/http"
)

type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

func JSON(w http.ResponseWriter, h interface{}, status int) {
	resp, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(resp)
}

func Alexa(w http.ResponseWriter, speech string, status int) {
	outputSpeech := OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	alexa := AlexaResponse{
		Version:  "1.0",
		Response: Response{OutputSpeech: outputSpeech, ShouldEndSession: true},
	}

	resp, err := json.Marshal(alexa)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(status)
	w.Write(resp)
}
