package main

import (
	"log"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	// Logging for request handling
	log.Printf("Handling request: %s %s", r.Method, r.URL.Path)

	_ = app.writeJSON(w, http.StatusOK, payload)
}
