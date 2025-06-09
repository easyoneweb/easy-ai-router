package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Port                   string
	AccessOpenrouterApiKey string
}

type EnvVars struct {
	Port                   string
	DBURI                  string
	DBName                 string
	AccessOpenrouterApiKey string
	OpenrouterHost         string
	OpenrouterApiKey       string
	OpenrouterLimit        string
}

var envVars = EnvVars{
	Port:                   "PORT",
	DBURI:                  "DB_URI",
	DBName:                 "DB_NAME",
	AccessOpenrouterApiKey: "ACCESS_OPENROUTER_API_KEY",
	OpenrouterHost:         "OPENROUTER_HOST",
	OpenrouterApiKey:       "OPENROUTER_API_KEY",
	OpenrouterLimit:        "OPENROUTER_LIMIT",
}

func main() {
	godotenv.Load()

	portString := os.Getenv(envVars.Port)
	if portString == "" {
		log.Fatal("PORT env variable not provided")
	}

	accessOpenrouter := os.Getenv(envVars.AccessOpenrouterApiKey)
	if accessOpenrouter == "" {
		log.Fatal("API_KEY env variable not provided")
	}

	apiCfg := ApiConfig{
		Port:                   portString,
		AccessOpenrouterApiKey: accessOpenrouter,
	}

	setupDB()
	setupOpenrouter()

	router := chi.NewRouter()

	router.Route("/openrouter", func(r chi.Router) {
		r.Use(apiCfg.checkOpenrouterAccess)

		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/ping", openrouterPing)
			r.Get("/key", openrouterKey)
			r.Get("/limits", openrouterLimits)

			r.Route("/chat", func(r chi.Router) {
				r.Use(checkOpenrouterLimits)

				r.Post("/", openrouterChat)
				r.Post("/image", openrouterChatWithImage)
			})
		})
	})

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + apiCfg.Port,
	}

	log.Printf("[server]: starting on port %v", apiCfg.Port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
