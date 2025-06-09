package openrouter

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestKeyData(t *testing.T) KeyData {
	t.Helper()
	return KeyData{
		Label:             "test-label",
		Usage:             0,
		IsFreeTier:        true,
		IsProvisioningKey: true,
		RateLimit: RateLimit{
			Requests: 10,
			Interval: "10s",
		},
		Limit:          50,
		LimitRemaining: 50,
	}
}

func handlerGetKeyInfo(t *testing.T) http.HandlerFunc {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		testKeyData := getTestKeyData(t)
		resp := KeyResponse{
			Data: testKeyData,
		}

		json, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte{})
			return
		}

		w.Header().Add("Contect-Type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	}
}

func TestGetKeyInfo(t *testing.T) {
	setupTestOpenrouterConfig(t)

	server := httptest.NewServer(http.HandlerFunc(handlerGetKeyInfo(t)))
	defer server.Close()

	// Set Openrouter config host to mock server's URL
	config.Host = server.URL
	// Set Key data url to root of mock server
	config.Urls.apiV1.key = ""

	t.Run("Get key info", func(t *testing.T) {
		testKeyData := getTestKeyData(t)
		resp, err := GetKeyInfo()
		if err != nil {
			t.Errorf("Failed to get key info: %v", err)
		}
		if resp.Data.Label != testKeyData.Label {
			t.Errorf("Failed to assert Label, expected: %v, got: %v", testKeyData.Label, resp.Data.Label)
		}
		if resp.Data.Usage != testKeyData.Usage {
			t.Errorf("Failed to assert Usage, expected: %v, got: %v", testKeyData.Usage, resp.Data.Usage)
		}
		if resp.Data.IsFreeTier != testKeyData.IsFreeTier {
			t.Errorf("Failed to assert IsFreeTier, expected: %v, got: %v", testKeyData.IsFreeTier, resp.Data.IsFreeTier)
		}
		if resp.Data.IsProvisioningKey != testKeyData.IsProvisioningKey {
			t.Errorf("Failed to assert IsProvisioningKey, expected: %v, got: %v", testKeyData.IsProvisioningKey, resp.Data.IsProvisioningKey)
		}
		if resp.Data.RateLimit.Interval != testKeyData.RateLimit.Interval {
			t.Errorf("Failed to assert RateLimit.Interval, expected: %v, got: %v", testKeyData.RateLimit.Interval, resp.Data.RateLimit.Interval)
		}
		if resp.Data.RateLimit.Requests != testKeyData.RateLimit.Requests {
			t.Errorf("Failed to assert RateLimit.Requests, expected: %v, got: %v", testKeyData.RateLimit.Requests, resp.Data.RateLimit.Requests)
		}
		if resp.Data.Limit != testKeyData.Limit {
			t.Errorf("Failed to assert Limit, expected: %v, got: %v", testKeyData.Limit, resp.Data.Limit)
		}
		if resp.Data.LimitRemaining != testKeyData.LimitRemaining {
			t.Errorf("Failed to assert LimitRemaining, expected: %v, got: %v", testKeyData.LimitRemaining, resp.Data.LimitRemaining)
		}
	})
}
