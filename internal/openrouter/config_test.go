package openrouter

import "testing"

const (
	host   = "http://localhost:3333"
	apiKey = "test-api-key"
	limit  = 50
)

func setupTestOpenrouterConfig(t *testing.T) {
	t.Helper()
	newConfig := OpenrouterConfig{
		Host:   host,
		ApiKey: apiKey,
		Limit:  limit,
	}
	SetConfig(newConfig)
}

func TestSetConfig(t *testing.T) {
	setupTestOpenrouterConfig(t)

	t.Run("Set config", func(t *testing.T) {
		if config.Host != host {
			t.Errorf("Failed to assert Host, expected: %v, got: %v", host, config.Host)
		}
		if config.ApiKey != apiKey {
			t.Errorf("Failed to assert ApiKey, expected: %v, got: %v", apiKey, config.ApiKey)
		}
		if config.Limit != limit {
			t.Errorf("Failed to assert Limit, expected: %v, got: %v", limit, config.Limit)
		}
	})
}
