package openrouter

import (
	"time"

	"github.com/easyoneweb/easy-ai-router/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Gets total limit for openrouter provider from env variable, used limits for today
// based on db logs.
// Return how many requests left, total limit, error if error occured.
// Return of -1 for how many requests left means infinite.
func GetTodayLimits() (int, int, error) {
	config := getConfig()
	if config.limit == -1 {
		return -1, config.limit, nil
	}

	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 1, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())

	timeFilter := bson.M{"createdAt": bson.M{"$gte": startTime, "$lt": endTime}}
	logs, err := database.GetLogsByProvider("openrouter", timeFilter)
	if err != nil {
		return 0, config.limit, err
	}

	return config.limit - len(logs), config.limit, nil
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
