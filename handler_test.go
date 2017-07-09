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
)

func TestHeartBeat(t *testing.T) {
	req, err := http.NewRequest("GET", "/heartbeat", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HeartBeat)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `OK`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestRiddleHandler(t *testing.T) {
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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	alexaResp := alexaskill.AlexaResponse{}
	json.Unmarshal(rr.Body.Bytes(), &alexaResp)
}

func TestRepeatRiddleIntentNoSession(t *testing.T) {
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
	if alexaResp.Response.OutputSpeech.Text != "No riddles have been given yet" {
		t.Errorf("Wrong respond message: expected %v, actual %v", "No riddles have been given yet", alexaResp.Response.OutputSpeech.Text)
	}
}
func TestRepeatRiddleWithSession(t *testing.T) {
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

	if alexaResp.Response.OutputSpeech.Text != expect {
		t.Errorf("Wrong respond message: expected %v, actual %v", "No riddles have been given yet", alexaResp.Response.OutputSpeech.Text)
	}
}

func TestAnswerRiddleWrongAnswer(t *testing.T) {
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

	if alexaResp.Response.OutputSpeech.Text != expect {
		t.Errorf("Wrong respond message: expected: %v; actual: %v", expect, alexaResp.Response.OutputSpeech.Text)
	}
}

func TestAnswerRiddleNoSession(t *testing.T) {
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

	if alexaResp.Response.OutputSpeech.Text != expect {
		t.Errorf("Wrong respond message: expected: %v; actual: %v", expect, alexaResp.Response.OutputSpeech.Text)
	}

	if alexaResp.Response.ShouldEndSession != true {
		t.Errorf("Session should be open")
	}
}
