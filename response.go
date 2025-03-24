package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonErrorResponse(w http.ResponseWriter, code int, message string) {
	js, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

func jsonResponse(w http.ResponseWriter, code int, data any) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Printf("handlers_openrouter: %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}
