package openrouter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type KeyResponse struct {
	Data KeyData `json:"data"`
}

type KeyData struct {
	Label             string    `json:"label"`
	Usage             float32   `json:"usage"`
	IsFreeTier        bool      `json:"is_free_tier"`
	IsProvisioningKey bool      `json:"is_provisioning_key"`
	RateLimit         RateLimit `json:"rate_limit"`
	Limit             int       `json:"limit"`
	LimitRemaining    float32   `json:"limit_remaining"`
}

type RateLimit struct {
	Requests int    `json:"requests"`
	Interval string `json:"interval"`
}

func GetKeyInfo() (KeyResponse, error) {
	config := getConfig()

	resp, err := http.NewRequest(http.MethodGet, config.host+config.urls.apiV1.key, nil)
	if err != nil {
		return KeyResponse{}, errors.New(openrouterErrors.CreateRequest)
	}
	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.apiKey))

	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return KeyResponse{}, errors.New(openrouterErrors.DoRequest)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return KeyResponse{}, errors.New(openrouterErrors.ReadBody)
	}

	key := KeyResponse{}
	err = json.Unmarshal(body, &key)
	if err != nil {
		return KeyResponse{}, errors.New(openrouterErrors.UnmarshalJson)
	}

	return key, nil
}
