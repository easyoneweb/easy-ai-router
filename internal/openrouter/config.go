package openrouter

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type OpenrouterConfig struct {
	host   string
	apiKey string
	urls   Urls
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
}

var envVars = EnvVars{
	host:   "OPENROUTER_HOST",
	apiKey: "OPENROUTER_API_KEY",
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

	openrouterConfig := OpenrouterConfig{
		host:   hostString,
		apiKey: apiKeyString,
		urls:   urls,
	}

	return openrouterConfig
}
