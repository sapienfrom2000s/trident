package server

import (
	"net/http"
)

func RootHandler(h http.ResponseWriter, r *http.Request) {
	welcomeString := "Ready for Trident!! 🔱"
	h.Write([]byte(welcomeString))
}

func HeartBeatHandler(h http.ResponseWriter, r *http.Request) {
}

func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", RootHandler)
	mux.HandleFunc("POST /event/github", HeartBeatHandler)
	http.ListenAndServe(":8080", mux)
}
