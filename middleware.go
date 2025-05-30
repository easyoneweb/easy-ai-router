package main

import (
	"net/http"

	"github.com/easyoneweb/easy-ai-router/internal/openrouter"
)

func (apiCfg *ApiConfig) checkOpenrouterAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("access-openrouter")
		if apiCfg.AccessOpenrouterApiKey != apiKey {
			w.WriteHeader(401)
			w.Write([]byte("not allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkOpenrouterLimits(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usedLimit, limit, err := openrouter.GetTodayLimits()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("error"))
			return
		}

		if usedLimit != -1 && usedLimit >= limit {
			w.WriteHeader(429)
			w.Write([]byte("now allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
