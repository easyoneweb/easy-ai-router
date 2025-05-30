package database

import "time"

type Log struct {
	CreatedAt time.Time `json:"created_at" bson:"createdAt"`
	Text      string    `json:"text" bson:"text"`
	Provider  string    `json:"provider" bson:"provider"`
	Type      string    `json:"type" bson:"type"`
}
