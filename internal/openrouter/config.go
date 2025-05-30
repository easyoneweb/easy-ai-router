package openrouter

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type OpenrouterConfig struct {
	host   string
	apiKey string
	urls   Urls
	limit  int
}

type Urls struct {
	apiV1 Endpoints
}

type Endpoints struct {
	key            string
	chatCompletion string
}

type EnvVars struct {
	host   string
	apiKey string
	limit  string
}

var envVars = EnvVars{
	host:   "OPENROUTER_HOST",
	apiKey: "OPENROUTER_API_KEY",
	limit:  "OPENROUTER_LIMIT",
}

var urls = Urls{
	apiV1: Endpoints{
		key:            "/api/v1/key",
		chatCompletion: "/api/v1/chat/completions",
	},
}

func getConfig() OpenrouterConfig {
	godotenv.Load()

	hostString := os.Getenv(envVars.host)
	if hostString == "" {
		log.Fatal("OPENROUTER_HOST env variable not provided")
	}

	apiKeyString := os.Getenv(envVars.apiKey)
	if apiKeyString == "" {
		log.Fatal("OPENROUTER_API_KEY env variable not provided")
	}

	limitString := os.Getenv(envVars.limit)
	if limitString == "" {
		log.Fatal("OPENROUTER_LIMIT env variable not provided")
	}
	if limitString == "-1" {
		log.Println("WARNING: OPENROUTER_LIMIT is set to -1, which means the limit is infinite")
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Fatal("OPENROUTER_LIMIT env variable could not be converted to int")
	}

	openrouterConfig := OpenrouterConfig{
		host:   hostString,
		apiKey: apiKeyString,
		urls:   urls,
		limit:  limit,
	}

	return openrouterConfig
}
