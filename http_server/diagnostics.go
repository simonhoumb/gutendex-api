package http_server

import (
	"encoding/json"
	"net/http"
	"time"
)

func DiagnosticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")

		statusOutput := Status{
			Gutendexapi:  http.StatusText(http.StatusBadGateway),
			Languageapi:  http.StatusText(http.StatusBadGateway),
			Countriesapi: http.StatusText(http.StatusBadGateway),
			Version:      "v1.0.0",
			Uptime:       time.Since(StartTime).Seconds()}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(statusOutput)

		if err != nil {
			http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		}
	}
}
