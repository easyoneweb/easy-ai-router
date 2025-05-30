package openrouter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ChatResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageWithImage struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string   `json:"type"`
	Text     string   `json:"text"`
	ImageUrl ImageUrl `json:"image_url"`
}

type ContentImage struct {
	Type     string   `json:"type"`
	ImageUrl ImageUrl `json:"image_url"`
}

type ImageUrl struct {
	Url string `json:"url"`
}

type Model struct {
	DeepSeekR1 string
	Gemma3     string
}

type PostBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type PostWithImageBody struct {
	Model    string             `json:"model"`
	Messages []MessageWithImage `json:"messages"`
}

var Models = Model{
	DeepSeekR1: "deepseek/deepseek-r1:free",
	Gemma3:     "google/gemma-3-12b-it:free",
}

func Chat(messages []Message, requestIdentity string) (ChatResponse, error) {
	config := getConfig()

	postBody, err := json.Marshal(PostBody{
		Model:    Models.DeepSeekR1,
		Messages: messages,
	})
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.MarshalJson)
	}

	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost, config.host+config.urls.apiV1.chatCompletion, requestBody)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.CreateRequest)
	}
	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.apiKey))

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.DoRequest)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.ReadBody)
	}

	chat := ChatResponse{}
	err = json.Unmarshal(body, &chat)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.UnmarshalJson)
	}

	err = CreateLimitLog("chat", requestIdentity)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.LimitLog)
	}

	return chat, nil
}

func ChatWithImage(messages []MessageWithImage, requestIdentity string) (ChatResponse, error) {
	config := getConfig()

	postBody, err := json.Marshal(PostWithImageBody{
		Model:    Models.Gemma3,
		Messages: messages,
	})
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.MarshalJson)
	}

	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost, config.host+config.urls.apiV1.chatCompletion, requestBody)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.CreateRequest)
	}
	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.apiKey))

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.DoRequest)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.ReadBody)
	}

	chat := ChatResponse{}
	err = json.Unmarshal(body, &chat)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.UnmarshalJson)
	}

	err = CreateLimitLog("chat_with_image", requestIdentity)
	if err != nil {
		return ChatResponse{}, errors.New(openrouterErrors.LimitLog)
	}

	return chat, nil
}
