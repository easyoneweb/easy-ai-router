package main

import (
	"encoding/json"
	"net/http"

	"github.com/easyoneweb/easy-ai-router/internal/openrouter"
)

type OpenrouterChatRequest struct {
	Messages          []openrouter.Message          `json:"messages"`
	MessagesWithImage []openrouter.MessageWithImage `json:"messages_with_image"`
	RequestIdentity   string                        `json:"request_identity"`
}

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

func openrouterLimits(w http.ResponseWriter, r *http.Request) {
	type LimitsResponse struct {
		UsedLimit int `json:"used_limit"`
		Limit     int `json:"limit"`
	}
	usedLimit, limit, err := openrouter.GetTodayLimits()
	if err != nil {
		jsonErrorResponse(w, 400, handlerErrors.OpenrouterErrors.Limits)
		return
	}

	jsonResponse(w, 200, LimitsResponse{UsedLimit: usedLimit, Limit: limit})
}

func openrouterChat(w http.ResponseWriter, r *http.Request) {
	reqData := OpenrouterChatRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		jsonErrorResponse(w, 422, handlerErrors.OpenrouterErrors.ChatBody)
		return
	}

	chat, err := openrouter.Chat(reqData.Messages, reqData.RequestIdentity)
	if err != nil {
		jsonErrorResponse(w, 400, handlerErrors.OpenrouterErrors.Chat)
	}

	jsonResponse(w, 200, chat)
}

func openrouterChatWithImage(w http.ResponseWriter, r *http.Request) {
	reqData := OpenrouterChatRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		jsonErrorResponse(w, 422, handlerErrors.OpenrouterErrors.ChatBody)
		return
	}

	chat, err := openrouter.ChatWithImage(reqData.MessagesWithImage, reqData.RequestIdentity)
	if err != nil {
		jsonErrorResponse(w, 400, handlerErrors.OpenrouterErrors.Chat)
	}

	jsonResponse(w, 200, chat)
}
