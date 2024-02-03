package http_server

import (
	"fmt"
	"net/http"
	"time"
)

func DiagnosticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	output := "{\n"
	output += fmt.Sprintf("\"gutendexapi\": \"%s\",\n", http.StatusText(http.StatusBadGateway))
	output += fmt.Sprintf("\"languageapi\": \"%s\",\n", http.StatusText(http.StatusBadGateway))
	output += fmt.Sprintf("\"countriesapi\": \"%s\",\n", http.StatusText(http.StatusBadGateway))
	output += fmt.Sprintf("\"version\": \"%s\",\n", "v1.0.0")
	output += fmt.Sprintf("\"uptime\": \"%f\"\n", time.Since(StartTime).Seconds())
	output += "}"

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		http.Error(w, "Error when returning output.", http.StatusInternalServerError)
	}
}
