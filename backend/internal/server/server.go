package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sapienfrom2000s/trident/backend/internal/core/models"
	"github.com/sapienfrom2000s/trident/backend/internal/webhook/github"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func RootHandler(h http.ResponseWriter, r *http.Request) {
	welcomeString := "Trident is ready!! 🔱"
	h.Write([]byte(welcomeString))
}

func HeartBeatHandler(h http.ResponseWriter, r *http.Request) {
}

func (s Server) GithubWebhookHandler(h http.ResponseWriter, r *http.Request) {
	handler := github.Handler{
		DB: s.db,
	}
	handler.WebhookHandler(h, r)
}

func Main() {
	fmt.Println("Connecting with Sqlite DB...")
	db, err := gorm.Open(sqlite.Open("trident.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected with Sqlite DB!")
	fmt.Println("Running Migrations...")
	db.AutoMigrate(&models.Event{}, &models.HeartBeat{}, &models.Job{})
	fmt.Println("Ran Migrations!")

	server := Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", RootHandler)
	mux.HandleFunc("POST /event/github", server.GithubWebhookHandler)

	fmt.Println("Starting Server...")
	http.ListenAndServe(":8080", mux)
}
