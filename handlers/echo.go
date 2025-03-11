package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Time    string `json:"time,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetEcho(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Time:    time.Now().String(),
		Message: "Hello World",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Fatal("[handlers:json-encoding] GET /echo could not encode response")
	}
}
