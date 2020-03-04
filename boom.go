package boom

import (
	"encoding/json"
	"log"
)

type Message struct {
	StatusCode int         `json:"statusCode"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func marshalJson(code int, error string, text string, data ...interface{}) []byte {
	message := &Message{
		StatusCode: code,
		Error:      error,
		Message:    text,
	}

	if data != nil {
		message.Data = data
		if len(data) == 1 {
			message.Data = data[0]
		}
	}

	response, err := json.Marshal(message)
	if err != nil {
		log.Panicln(err)
	}
	return response
}
