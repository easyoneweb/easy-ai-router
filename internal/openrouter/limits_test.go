package openrouter

import (
	"testing"

	"github.com/easyoneweb/easy-ai-router/internal/mocks"
)

func TestGetTodayLimits(t *testing.T) {
	setupTestOpenrouterConfig(t)
	mocks.SetupTestDB(t)
	mocks.ClearTestDBLogs(t)

	t.Run("Get today's limits", func(t *testing.T) {
		v1, v2, err := GetTodayLimits()
		if err != nil {
			t.Errorf("Failed to get today's limits: %v", err)
		}
		if v1 != 0 {
			t.Errorf("Failed to assert used limit, expected: %v, got: %v", 0, v1)
		}
		if v2 != config.Limit {
			t.Errorf("Failed to assert total limit, expected: %v, got: %v", config.Limit, v2)
		}
	})
}

func TestCreateLimitLog(t *testing.T) {
	mocks.SetupTestDB(t)
	defer mocks.ClearTestDBLogs(t)

	testLogText := "test-log"
	testLogType := "test-identity"
	t.Run("Create limit log", func(t *testing.T) {
		err := CreateLimitLog(testLogText, testLogType)
		if err != nil {
			t.Errorf("Failed to create limit log: %v", err)
		}
	})
}
