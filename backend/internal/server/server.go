package server

import (
	"net/http"

	"github.com/sapienfrom2000s/trident/backend/internal/webhook/github"
)

func RootHandler(h http.ResponseWriter, r *http.Request) {
	welcomeString := "Trident is ready!! 🔱"
	h.Write([]byte(welcomeString))
}

func HeartBeatHandler(h http.ResponseWriter, r *http.Request) {
}

func GithubWebhookHandler(h http.ResponseWriter, r *http.Request) {
	handler := github.Handler{}
	handler.WebhookHandler(h, r)
}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", RootHandler)
	mux.HandleFunc("POST /event/github", GithubWebhookHandler)
	http.ListenAndServe(":8080", mux)
}
