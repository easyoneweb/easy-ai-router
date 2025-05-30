package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBConfig struct {
	DB   *mongo.Database
	Logs *mongo.Collection
}

var db DBConfig

// Connects to local MongoDB instance.
// Local package variable db receives pointers to db instance, collection logs.
// Returns an error if connection is not possible or returned an error.
func Connect(dbURI string, dbName string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(dbURI))
	if err != nil {
		return err
	}

	db.DB = client.Database(dbName)
	db.Logs = db.DB.Collection("logs")
	createLogIndexes()
	return nil
}

// Creates/checks indexes on Log documents, fields: createdAt, provider.
func createLogIndexes() {
	indexModels := []mongo.IndexModel{
		{Keys: bson.D{{Key: "createdAt", Value: -1}}},
		{Keys: bson.D{{Key: "provider", Value: 1}}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := db.Logs.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Fatal(err)
	}
}

// Creates Log document in database.
func CreateLog(l Log) error {
	l.CreatedAt = time.Now()

	_, err := db.Logs.InsertOne(context.Background(), l)
	if err != nil {
		return errors.New("could not create new log")
	}

	return nil
}

// Gets all Log documents from database with provided limit and skip.
// Returns empty Log slice and an error if error occured.
func GetLogs(limit int64, skip int64) ([]Log, error) {
	result := []Log{}
	opts := options.Find().SetLimit(limit).SetSkip(skip)
	cursor, err := db.Logs.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return result, errors.New("could not execute collection logs find")
	}

	if err = cursor.All(context.Background(), &result); err != nil {
		return result, errors.New("could not execute collection logs cursor all")
	}

	return result, nil
}

// Gets all Log documents from database by provider string and in time range by time filter.
// Returns empty Log slice and an error if error occured.
func GetLogsByProvider(provider string, timeFilter bson.M) ([]Log, error) {
	result := []Log{}
	filter := bson.M{"createdAt": timeFilter, "provider": provider}
	cursor, err := db.Logs.Find(context.Background(), filter)
	if err != nil {
		return result, errors.New("could not execute collection logs find by provider")
	}

	if err = cursor.All(context.Background(), &result); err != nil {
		return result, errors.New("could not execute collection logs cursor all by provider")
	}

	return result, nil
}