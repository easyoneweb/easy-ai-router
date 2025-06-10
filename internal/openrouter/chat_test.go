package openrouter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/easyoneweb/easy-ai-router/internal/mocks"
)

func getTestChatResponseData(t *testing.T) ChatResponse {
	t.Helper()
	message := Message{
		Role:    "assistant",
		Content: "chat-completion-text",
	}
	choices := make([]Choice, 0, 1)
	choices = append(choices, Choice{Message: message})

	return ChatResponse{
		ID:      "test-id",
		Choices: choices,
	}
}

func getTestChatPostBody(t *testing.T) PostBody {
	t.Helper()

	message := Message{
		Role:    "user",
		Content: "chat-text",
	}
	messages := make([]Message, 0, 1)
	messages = append(messages, message)

	return PostBody{
		Model:    "test-model",
		Messages: messages,
	}
}

func getTestChatPostWithImageBody(t *testing.T) PostWithImageBody {
	t.Helper()
	contentText := Content{
		Type: "text",
		Text: "test-text",
	}
	contentImage := Content{
		Type:     "image_url",
		ImageUrl: ImageUrl{Url: "text-image"},
	}
	contents := make([]Content, 0, 2)
	contents = append(contents, contentText, contentImage)
	message := MessageWithImage{
		Role:    "user",
		Content: contents,
	}
	messages := make([]MessageWithImage, 0, 1)
	messages = append(messages, message)

	return PostWithImageBody{
		Model:    "test-model",
		Messages: messages,
	}
}

func handlerChat(t *testing.T) http.HandlerFunc {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		reqBody := PostBody{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reqBody)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		err = validateChatRequestBody(t, reqBody)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		chatResponse := getTestChatResponseData(t)
		json, err := json.Marshal(chatResponse)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	}
}

func validateChatRequestBody(t *testing.T, body PostBody) error {
	t.Helper()

	testPostBody := getTestChatPostBody(t)
	if body.Model != testPostBody.Model {
		return errors.New("Model")
	}
	if len(body.Messages) != 1 {
		return errors.New("length of Messages")
	}
	if body.Messages[0].Role != testPostBody.Messages[0].Role {
		return errors.New("Messages.Role")
	}
	if body.Messages[0].Content != testPostBody.Messages[0].Content {
		return errors.New("Messages.Content")
	}

	return nil
}

func handlerChatWithImage(t *testing.T) http.HandlerFunc {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		reqBody := PostWithImageBody{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reqBody)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		err = validateChatWithImageRequestBody(t, reqBody)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		chatResponse := getTestChatResponseData(t)
		json, err := json.Marshal(chatResponse)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	}
}

func validateChatWithImageRequestBody(t *testing.T, body PostWithImageBody) error {
	t.Helper()

	testPostBody := getTestChatPostWithImageBody(t)
	if body.Model != testPostBody.Model {
		return errors.New("Model")
	}
	if len(body.Messages) != 1 {
		return errors.New("length of Messages")
	}
	if body.Messages[0].Role != testPostBody.Messages[0].Role {
		return errors.New("Messages.Role")
	}
	if len(body.Messages[0].Content) != 2 {
		return errors.New("length of Content")
	}
	if body.Messages[0].Content[0].Type != testPostBody.Messages[0].Content[0].Type {
		return errors.New("Messages.Content[0].Type")
	}
	if body.Messages[0].Content[0].Text != testPostBody.Messages[0].Content[0].Text {
		return errors.New("Messages.Content[0].Text")
	}
	if body.Messages[0].Content[1].Type != testPostBody.Messages[0].Content[1].Type {
		return errors.New("Messages.Content[1].Type")
	}
	if body.Messages[0].Content[1].Text != testPostBody.Messages[0].Content[1].Text {
		return errors.New("Messages.Content[1].Text")
	}

	return nil
}

func TestChat(t *testing.T) {
	setupTestOpenrouterConfig(t)
	// Setup TEST DB because Chat() creates log in DB
	mocks.SetupTestDB(t)
	defer mocks.ClearTestDBLogs(t)

	server := httptest.NewServer(http.HandlerFunc(handlerChat(t)))
	defer server.Close()

	// Set Openrouter config host to mock server's URL
	config.Host = server.URL
	// Set Chat completion url to root of mock server
	config.Urls.apiV1.chatCompletion = ""

	t.Run("Chat", func(t *testing.T) {
		testPostData := getTestChatPostBody(t)
		testChatResponse := getTestChatResponseData(t)

		resp, err := Chat(testPostData.Messages, testPostData.Model, "test-identity")
		if err != nil {
			t.Errorf("Failed to post chat: %v", err)
		}
		if resp.ID != testChatResponse.ID {
			t.Errorf("Failed to assert ID, expected: %v, got: %v", testChatResponse.ID, resp.ID)
		}
		if len(resp.Choices) != 1 {
			t.Errorf("Failed to assert number of Choices, expected: %v, got: %v", len(testChatResponse.Choices), len(resp.Choices))
		}
		if resp.Choices[0].Message.Role != testChatResponse.Choices[0].Message.Role {
			t.Errorf("Failed to assert Message.Role, expected: %v, got: %v", testChatResponse.Choices[0].Message.Role, resp.Choices[0].Message.Role)
		}
		if resp.Choices[0].Message.Content != testChatResponse.Choices[0].Message.Content {
			t.Errorf("Failed to assert Message.Content, expected: %v, got: %v", testChatResponse.Choices[0].Message.Content, resp.Choices[0].Message.Content)
		}
	})
}

func TestChatWithImage(t *testing.T) {
	setupTestOpenrouterConfig(t)
	// Setup TEST DB because ChatWithImage() creates log in DB
	mocks.SetupTestDB(t)
	defer mocks.ClearTestDBLogs(t)

	server := httptest.NewServer(http.HandlerFunc(handlerChatWithImage(t)))
	defer server.Close()

	// Set Openrouter config host to mock server's URL
	config.Host = server.URL
	// Set Chat completion url to root of mock server
	config.Urls.apiV1.chatCompletion = ""

	t.Run("Chat with image", func(t *testing.T) {
		testPostData := getTestChatPostWithImageBody(t)
		testChatResponse := getTestChatResponseData(t)

		resp, err := ChatWithImage(testPostData.Messages, testPostData.Model, "test-identity")
		if err != nil {
			t.Errorf("Failed to post chat with image: %v", err)
		}
		if resp.ID != testChatResponse.ID {
			t.Errorf("Failed to assert ID, expected: %v, got: %v", testChatResponse.ID, resp.ID)
		}
		if len(resp.Choices) != 1 {
			t.Errorf("Failed to assert number of Choices, expected: %v, got: %v", len(testChatResponse.Choices), len(resp.Choices))
		}
		if resp.Choices[0].Message.Role != testChatResponse.Choices[0].Message.Role {
			t.Errorf("Failed to assert Message.Role, expected: %v, got: %v", testChatResponse.Choices[0].Message.Role, resp.Choices[0].Message.Role)
		}
		if resp.Choices[0].Message.Content != testChatResponse.Choices[0].Message.Content {
			t.Errorf("Failed to assert Message.Content, expected: %v, got: %v", testChatResponse.Choices[0].Message.Content, resp.Choices[0].Message.Content)
		}
	})
}
