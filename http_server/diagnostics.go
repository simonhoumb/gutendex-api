package http_server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 3 * time.Second}

func DiagnosticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")

		statusOutput := Status{
			GutendexAPI:  getStatusCode(w, GUTENDEXAPI_URL),
			LanguageAPI:  getStatusCode(w, LANGUAGEAPI_URL),
			CountriesAPI: getStatusCode(w, COUNTRIESAPI_URL),
			Version:      VERSION,
			Uptime:       time.Since(StartTime).Seconds()}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(statusOutput)
		if err != nil {
			log.Println("Error during JSON encoding: ", err.Error())
			http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
		}

	}
}

func getStatusCode(w http.ResponseWriter, url string) int {
	defer client.CloseIdleConnections()

	res, err := client.Get(url)
	if err != nil {
		log.Println("Error when requesting remote endpoint: ", err.Error())
		return http.StatusServiceUnavailable
	}

	return res.StatusCode
}
