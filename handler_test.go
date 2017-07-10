// test different requests/responses from alexa
// test the intents
// test slots
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/carlqt/alexariddles/alexaskill"
	"github.com/carlqt/alexariddles/riddles"
	"github.com/stretchr/testify/assert"
)

func TestRiddleHandler(t *testing.T) {
	assert := assert.New(t)

	request := strings.NewReader(`{
  "session": {
    "sessionId": "SessionId.7098bcf8-9994-4bbf-8ae7-b41d85723a7d",
    "application": {
      "applicationId": "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"
    },
    "attributes": {},
    "user": {
      "userId": "amzn1.ask.account.AH2Y4V7T4UXBMQEQKKIV7WMMZKVOBMGGGKWGVES3KAHDNPUOD6BMR3WD3ZL2RFT6VD47DNMGGKLG4XKAYYLBHEQ2TAKJ5PA525SKS3GZOMJC7PQZHIYLMCDHCMTDOV6AKPLZWPAQN6HJ5VZ4RGRIUBB7FK7TRO72T6BTIFJH3NXJM2P6JBKWFKJ5SKQI4LDLJYMXX2T6BZSRCDA"
    },
    "new": true
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.1c73c114-22e7-4cf8-b8b1-720bc2123122",
    "locale": "en-US",
    "timestamp": "2017-07-08T14:41:46Z",
    "intent": {
      "name": "AskRiddle",
      "slots": {}
    }
  },
  "version": "1.0"
}`)

	req := httptest.NewRequest("POST", "/", request)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RiddleHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(http.StatusOK, rr.Code, "Status code should be 200")
}

func TestRepeatRiddleIntentNoSession(t *testing.T) {
	assert := assert.New(t)

	request := strings.NewReader(`{
  "session": {
    "sessionId": "SessionId.7098bcf8-9994-4bbf-8ae7-b41d85723a7d",
    "application": {
      "applicationId": "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"
    },
    "attributes": {},
    "user": {
      "userId": "amzn1.ask.account.AH2Y4V7T4UXBMQEQKKIV7WMMZKVOBMGGGKWGVES3KAHDNPUOD6BMR3WD3ZL2RFT6VD47DNMGGKLG4XKAYYLBHEQ2TAKJ5PA525SKS3GZOMJC7PQZHIYLMCDHCMTDOV6AKPLZWPAQN6HJ5VZ4RGRIUBB7FK7TRO72T6BTIFJH3NXJM2P6JBKWFKJ5SKQI4LDLJYMXX2T6BZSRCDA"
    },
    "new": true
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.1c73c114-22e7-4cf8-b8b1-720bc2123122",
    "locale": "en-US",
    "timestamp": "2017-07-08T14:41:46Z",
    "intent": {
      "name": "RepeatRiddle",
      "slots": {}
    }
  },
  "version": "1.0"
}`)

	req := httptest.NewRequest("POST", "/", request)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RiddleHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	alexaResp := alexaskill.AlexaResponse{}
	json.Unmarshal(rr.Body.Bytes(), &alexaResp)
	assert.Equal("No riddles have been given yet", alexaResp.Response.OutputSpeech.Text)
}
func TestRepeatRiddleWithSession(t *testing.T) {
	assert := assert.New(t)
	request := strings.NewReader(`{
  "session": {
    "sessionId": "SessionId.7098bcf8-9994-4bbf-8ae7-b41d85723a7d",
    "application": {
      "applicationId": "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"
    },
    "attributes": {
			"answer": "man"
		},
    "user": {
      "userId": "amzn1.ask.account.AH2Y4V7T4UXBMQEQKKIV7WMMZKVOBMGGGKWGVES3KAHDNPUOD6BMR3WD3ZL2RFT6VD47DNMGGKLG4XKAYYLBHEQ2TAKJ5PA525SKS3GZOMJC7PQZHIYLMCDHCMTDOV6AKPLZWPAQN6HJ5VZ4RGRIUBB7FK7TRO72T6BTIFJH3NXJM2P6JBKWFKJ5SKQI4LDLJYMXX2T6BZSRCDA"
    },
    "new": true
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.1c73c114-22e7-4cf8-b8b1-720bc2123122",
    "locale": "en-US",
    "timestamp": "2017-07-08T14:41:46Z",
    "intent": {
      "name": "RepeatRiddle",
      "slots": {}
    }
  },
  "version": "1.0"
}`)

	req := httptest.NewRequest("POST", "/", request)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RiddleHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	alexaResp := alexaskill.AlexaResponse{}
	json.Unmarshal(rr.Body.Bytes(), &alexaResp)
	expect := riddles.GetRiddle("man")

	assert.Equal(expect, alexaResp.Response.OutputSpeech.Text)
}

func TestAnswerRiddleWrongAnswer(t *testing.T) {
	assert := assert.New(t)
	request := strings.NewReader(`{
  "session": {
    "sessionId": "SessionId.8a422ce6-3243-46ee-afa9-6c8f33dabe7c",
    "application": {
      "applicationId": "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"
    },
    "attributes": {
			"answer": "man"
		},
    "user": {
      "userId": "amzn1.ask.account.AH2Y4V7T4UXBMQEQKKIV7WMMZKVOBMGGGKWGVES3KAHDNPUOD6BMR3WD3ZL2RFT6VD47DNMGGKLG4XKAYYLBHEQ2TAKJ5PA525SKS3GZOMJC7PQZHIYLMCDHCMTDOV6AKPLZWPAQN6HJ5VZ4RGRIUBB7FK7TRO72T6BTIFJH3NXJM2P6JBKWFKJ5SKQI4LDLJYMXX2T6BZSRCDA"
    },
    "new": false
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.4c232378-091f-4f4a-85d8-879c74d8b083",
    "locale": "en-US",
    "timestamp": "2017-07-09T07:36:49Z",
    "intent": {
      "name": "AnswerRiddle",
      "slots": {
        "RiddleAnswer": {
          "name": "ANSWER_LIST",
          "value": "manner"
        }
      }
    }
  },
  "version": "1.0"
}`)

	req := httptest.NewRequest("POST", "/", request)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RiddleHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	alexaResp := alexaskill.AlexaResponse{}
	json.Unmarshal(rr.Body.Bytes(), &alexaResp)
	expect := "Sorry, manner is not the answer. Try again"

	assert.Equal(expect, alexaResp.Response.OutputSpeech.Text)
}

func TestAnswerRiddleNoSession(t *testing.T) {
	assert := assert.New(t)
	request := strings.NewReader(`{
  "session": {
    "sessionId": "SessionId.8a422ce6-3243-46ee-afa9-6c8f33dabe7c",
    "application": {
      "applicationId": "amzn1.ask.skill.61e24a88-0159-4f67-983f-d974aa6b8d64"
    },
    "attributes": {},
    "user": {
      "userId": "amzn1.ask.account.AH2Y4V7T4UXBMQEQKKIV7WMMZKVOBMGGGKWGVES3KAHDNPUOD6BMR3WD3ZL2RFT6VD47DNMGGKLG4XKAYYLBHEQ2TAKJ5PA525SKS3GZOMJC7PQZHIYLMCDHCMTDOV6AKPLZWPAQN6HJ5VZ4RGRIUBB7FK7TRO72T6BTIFJH3NXJM2P6JBKWFKJ5SKQI4LDLJYMXX2T6BZSRCDA"
    },
    "new": false
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.4c232378-091f-4f4a-85d8-879c74d8b083",
    "locale": "en-US",
    "timestamp": "2017-07-09T07:36:49Z",
    "intent": {
      "name": "AnswerRiddle",
      "slots": {
        "RiddleAnswer": {
          "name": "ANSWER_LIST",
          "value": "manner"
        }
      }
    }
  },
  "version": "1.0"
}`)

	req := httptest.NewRequest("POST", "/", request)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RiddleHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	alexaResp := alexaskill.AlexaResponse{}
	json.Unmarshal(rr.Body.Bytes(), &alexaResp)
	expect := "No riddles have been given yet"

	assert.Equal(expect, alexaResp.Response.OutputSpeech.Text)
	assert.True(alexaResp.Response.ShouldEndSession)
}
