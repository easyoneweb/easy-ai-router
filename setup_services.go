package main

import (
	"log"
	"os"
	"strconv"

	"github.com/easyoneweb/easy-ai-router/internal/database"
	"github.com/easyoneweb/easy-ai-router/internal/openrouter"
	"github.com/joho/godotenv"
)

func setupDB() {
	godotenv.Load()

	dbUri := os.Getenv(envVars.DBURI)
	if dbUri == "" {
		log.Fatal("DB_URI env variable not provided")
	}

	dbName := os.Getenv(envVars.DBName)
	if dbName == "" {
		log.Fatal("DB_NAME env variable not provided")
	}

	err := database.Connect(dbUri, dbName)
	if err != nil {
		log.Fatal("could not connect to database")
	}
}

func setupOpenrouter() {
	openrouterHost := os.Getenv(envVars.OpenrouterHost)
	if openrouterHost == "" {
		log.Fatal("OPENROUTER_HOST env variable not provided")
	}

	openrouterApiKey := os.Getenv(envVars.OpenrouterApiKey)
	if openrouterApiKey == "" {
		log.Fatal("OPENROUTER_API_KEY env variable not provided")
	}

	openrouterLimit := os.Getenv(envVars.OpenrouterLimit)
	if openrouterLimit == "" {
		log.Fatal("OPENROUTER_LIMIT env var not provided")
	}

	openrouterLimitInt, err := strconv.Atoi(openrouterLimit)
	if err != nil {
		log.Fatal("OPENROUTER_LIMIT env variable could not be converted to int")
	}

	openrouter.SetConfig(openrouter.OpenrouterConfig{
		Host:   openrouterHost,
		ApiKey: openrouterApiKey,
		Limit:  openrouterLimitInt,
	})
}
