package mocks

import (
	"testing"

	"github.com/easyoneweb/easy-ai-router/internal/database"
)

const (
	testDBURI  = "mongodb://localhost:27017"
	testDBName = "testdb"
)

func SetupTestDB(t *testing.T) {
	t.Helper()
	err := database.Connect(testDBURI, testDBName)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	err = database.DeleteLogsByProdvider("openrouter")
	if err != nil {
		t.Fatalf("Failed to clean logs by provider: %v", err)
	}
}

func ClearTestDBLogs(t *testing.T) {
	t.Helper()
	err := database.DeleteLogsByProdvider("openrouter")
	if err != nil {
		t.Fatalf("Failed to clean logs by provider: %v", err)
	}
}
