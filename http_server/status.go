package http_server

import (
	"log"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 3 * time.Second}

type Status struct {
	GutendexAPI  int     `json:"gutendexapi"`
	LanguageAPI  int     `json:"languageapi"`
	CountriesAPI int     `json:"countriesapi"`
	Version      string  `json:"version"`
	Uptime       float64 `json:"uptime"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")

		statusOutput := Status{
			GutendexAPI:  getStatusCode(w, GUTENDEXAPI_URL),
			LanguageAPI:  getStatusCode(w, LANGUAGEAPI_URL),
			CountriesAPI: getStatusCode(w, COUNTRIESAPI_URL),
			Version:      VERSION,
			Uptime:       time.Since(StartTime).Seconds()}

		encodeJSON(w, statusOutput)

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
