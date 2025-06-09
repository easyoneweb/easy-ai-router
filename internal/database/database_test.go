package database

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Tests are written using deepseek-r1 model
const (
	testDBURI  = "mongodb://localhost:27017"
	testDBName = "testdb"
)

func setupTestDB(t *testing.T) {
	t.Helper()
	err := Connect(testDBURI, testDBName)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Clean up any existing data
	_, err = db.Logs.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func teardownTestDB(t *testing.T) {
	t.Helper()
	if db.DB != nil {
		err := db.DB.Drop(context.Background())
		if err != nil {
			t.Logf("Warning: failed to drop test database: %v", err)
		}
	}
}

func TestConnect(t *testing.T) {
	t.Run("successful connection", func(t *testing.T) {
		err := Connect(testDBURI, testDBName)
		if err != nil {
			t.Errorf("Connect() error = %v, want nil", err)
		}

		if db.DB == nil {
			t.Error("Connect() db.DB is nil, want *mongo.Database")
		}

		if db.Logs == nil {
			t.Error("Connect() db.Logs is nil, want *mongo.Collection")
		}
	})

	t.Run("invalid URI", func(t *testing.T) {
		err := Connect("invalid_uri", testDBName)
		if err == nil {
			t.Error("Connect() with invalid URI should return error, got nil")
		}
	})
}

func TestCreateLog(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	testCases := []struct {
		name    string
		log     Log
		wantErr bool
	}{
		{
			name: "valid log",
			log: Log{
				Text:     "test log",
				Provider: "test-provider",
				Type:     "info",
			},
			wantErr: false,
		},
		{
			name: "empty text",
			log: Log{
				Text:     "",
				Provider: "test-provider",
				Type:     "info",
			},
			wantErr: false, // Assuming empty text is allowed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateLog(tc.log)
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateLog() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				// Verify the log was actually created
				var result Log
				err := db.Logs.FindOne(context.Background(), bson.M{"text": tc.log.Text}).Decode(&result)
				if err != nil {
					t.Errorf("Failed to find created log: %v", err)
				}
				if result.Text != tc.log.Text {
					t.Errorf("Created log text = %v, want %v", result.Text, tc.log.Text)
				}
			}
		})
	}
}

func TestGetLogs(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	// Insert test data
	testLogs := []Log{
		{
			Text:      "log 1",
			Provider:  "provider-1",
			Type:      "info",
			CreatedAt: time.Now().Add(-time.Hour),
		},
		{
			Text:      "log 2",
			Provider:  "provider-2",
			Type:      "error",
			CreatedAt: time.Now().Add(-30 * time.Minute),
		},
		{
			Text:      "log 3",
			Provider:  "provider-1",
			Type:      "warning",
			CreatedAt: time.Now(),
		},
	}

	for _, log := range testLogs {
		_, err := db.Logs.InsertOne(context.Background(), log)
		if err != nil {
			t.Fatalf("Failed to insert test logs: %v", err)
		}
	}

	t.Run("get all logs", func(t *testing.T) {
		logs, err := GetLogs(10, 0)
		if err != nil {
			t.Errorf("GetLogs() error = %v, want nil", err)
		}

		if len(logs) != len(testLogs) {
			t.Errorf("GetLogs() returned %d logs, want %d", len(logs), len(testLogs))
		}
	})

	t.Run("get logs with limit", func(t *testing.T) {
		limit := int64(2)
		logs, err := GetLogs(limit, 0)
		if err != nil {
			t.Errorf("GetLogs() error = %v, want nil", err)
		}

		if len(logs) != int(limit) {
			t.Errorf("GetLogs() returned %d logs, want %d", len(logs), limit)
		}
	})

	t.Run("get logs with skip", func(t *testing.T) {
		skip := int64(1)
		logs, err := GetLogs(10, skip)
		if err != nil {
			t.Errorf("GetLogs() error = %v, want nil", err)
		}

		if len(logs) != len(testLogs)-int(skip) {
			t.Errorf("GetLogs() returned %d logs, want %d", len(logs), len(testLogs)-int(skip))
		}
	})
}

func TestGetLogsByProvider(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	// Insert test data
	now := time.Now()
	testLogs := []Log{
		{
			Text:      "log 1",
			Provider:  "provider-1",
			Type:      "info",
			CreatedAt: now.Add(-2 * time.Hour),
		},
		{
			Text:      "log 2",
			Provider:  "provider-2",
			Type:      "error",
			CreatedAt: now.Add(-time.Hour),
		},
		{
			Text:      "log 3",
			Provider:  "provider-1",
			Type:      "warning",
			CreatedAt: now,
		},
	}

	for _, log := range testLogs {
		_, err := db.Logs.InsertOne(context.Background(), log)
		if err != nil {
			t.Fatalf("Failed to insert test logs: %v", err)
		}
	}

	t.Run("get logs by provider", func(t *testing.T) {
		timeFilter := bson.M{
			"$gte": now.Add(-24 * time.Hour),
			"$lte": now,
		}

		logs, err := GetLogsByProvider("provider-1", timeFilter)
		if err != nil {
			t.Errorf("GetLogsByProvider() error = %v, want nil", err)
		}

		expectedCount := 2 // log 1 and log 3
		if len(logs) != expectedCount {
			t.Errorf("GetLogsByProvider() returned %d logs, want %d", len(logs), expectedCount)
		}

		for _, log := range logs {
			if log.Provider != "provider-1" {
				t.Errorf("GetLogsByProvider() returned log with provider %s, want provider-1", log.Provider)
			}
		}
	})

	t.Run("get logs by provider with time filter", func(t *testing.T) {
		timeFilter := bson.M{
			"$gte": now.Add(-90 * time.Minute),
			"$lte": now,
		}

		logs, err := GetLogsByProvider("provider-1", timeFilter)
		if err != nil {
			t.Errorf("GetLogsByProvider() error = %v, want nil", err)
		}

		expectedCount := 1 // only log 3 should match
		if len(logs) != expectedCount {
			t.Errorf("GetLogsByProvider() returned %d logs, want %d", len(logs), expectedCount)
		}
	})

	t.Run("no matching logs", func(t *testing.T) {
		timeFilter := bson.M{
			"$gte": now.Add(-24 * time.Hour),
			"$lte": now,
		}

		logs, err := GetLogsByProvider("nonexistent-provider", timeFilter)
		if err != nil {
			t.Errorf("GetLogsByProvider() error = %v, want nil", err)
		}

		if len(logs) != 0 {
			t.Errorf("GetLogsByProvider() returned %d logs, want 0", len(logs))
		}
	})
}
