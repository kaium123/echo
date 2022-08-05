package utility

import (
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func CreateMessage(msg string) *Message {
	return &Message{
		Message: msg,
		Status:  http.StatusOK,
	}
}

func UpdateMessage(msg string) *Message {
	return &Message{
		Message: msg,
		Status:  http.StatusOK,
	}
}

func DeleteMessage(msg string) *Message {
	return &Message{
		Message: msg,
		Status:  http.StatusOK,
	}
}
