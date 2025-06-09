package openrouter

import (
	"time"

	"github.com/easyoneweb/easy-ai-router/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Gets limits for openrouter provider.
// Total limit is defined in env variable OPENROUTER_LIMIT.
// Returns how many requests were made today and total limit, error if error occured.
// Used limit and total limit -1 means infinite.
func GetTodayLimits() (int, int, error) {
	if config.Limit == -1 {
		return -1, config.Limit, nil
	}

	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 1, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

	timeFilter := bson.M{"$gte": startTime, "$lt": endTime}
	logs, err := database.GetLogsByProvider("openrouter", timeFilter)
	if err != nil {
		return 0, config.Limit, err
	}

	return len(logs), config.Limit, nil
}

// Creates log in db.
// Should be used for every openrouter request since logs are used to track the limit
// of request to openrouter api.
func CreateLimitLog(text string, typeOfLog string) error {
	log := database.Log{
		Text:     text,
		Provider: "openrouter",
		Type:     typeOfLog,
	}

	err := database.CreateLog(log)
	if err != nil {
		return err
	}

	return nil
}
