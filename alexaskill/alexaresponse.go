package alexaskill

import (
	"encoding/json"
	"net/http"
)

type AlexaResponse struct {
	Version           string      `json:"version"`
	SessionAttributes sessionAttr `json:"sessionAttributes"`
	Response          Response    `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
	Reprompt         struct {
		OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
	} `json:"reprompt,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Text    string `json:"text"`
}

// NewAlexaResponse is a constructor that sets the bare minimum for the AlexaResponse struct
func NewAlexaResponse(version string) *AlexaResponse {
	session := make(map[string]string)
	return &AlexaResponse{
		Version:           "1.0",
		SessionAttributes: session,
	}
}

// NewAlexaText is a constructor much like the NewAlexaResponse but defaults the outputSpeech.Type field to PlainText
func NewAlexaText(speech string) *AlexaResponse {
	outputSpeech := OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	resp := NewAlexaResponse("1.0")
	resp.Response.OutputSpeech = outputSpeech
	resp.Response.ShouldEndSession = false

	return resp
}

// AlexaText is a method much like the NewAlexaText constructor much like the NewAlexaResponse but defaults the outputSpeech.Type field to PlainText
func (a *AlexaResponse) AlexaText(speech string) *AlexaResponse {
	outputSpeech := OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	a.Response.OutputSpeech = outputSpeech
	a.Response.ShouldEndSession = false

	return a
}

func (a *AlexaResponse) SessionAttr(key, value string) *AlexaResponse {
	a.SessionAttributes[key] = value

	return a
}

func (a *AlexaResponse) SimpleCard(title, content string) *AlexaResponse {
	a.Response.Card = Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return a
}

func (a *AlexaResponse) Respond(w http.ResponseWriter, status int) {
	resp, _ := json.Marshal(a)

	w.WriteHeader(status)
	w.Write(resp)
}

func (a *AlexaResponse) EndSession(b bool) *AlexaResponse {
	a.Response.ShouldEndSession = b
	return a
}

func (a *AlexaResponse) RepromptText(speech string) *AlexaResponse {
	outputSpeech := &OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	a.Response.Reprompt.OutputSpeech = outputSpeech
	return a
}
