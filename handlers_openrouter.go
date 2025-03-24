package main

import (
	"encoding/json"
	"net/http"

	"github.com/ikirja/easy-ai-router/internal/openrouter"
)

func openrouterPing(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Message string `json:"message"`
	}{
		Message: "openrouter ping",
	}
	jsonResponse(w, 200, body)
}

func openrouterKey(w http.ResponseWriter, r *http.Request) {
	key, err := openrouter.GetKeyInfo()
	if err != nil {
		jsonErrorResponse(w, 500, handlerErrors.OpenrouterErrors.Key)
		return
	}

	jsonResponse(w, 200, key)
}

func openrouterChat(w http.ResponseWriter, r *http.Request) {
	messages := []openrouter.Message{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&messages)
	if err != nil {
		jsonErrorResponse(w, 422, handlerErrors.OpenrouterErrors.ChatBody)
		return
	}

	chat, err := openrouter.Chat(messages)
	if err != nil {
		jsonErrorResponse(w, 400, handlerErrors.OpenrouterErrors.Chat)
	}

	jsonResponse(w, 200, chat)
}
