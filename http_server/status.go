package http_server

import (
	"log"
	"net/http"
	"time"
)

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
			GutendexAPI:  getStatusCode(GUTENDEXAPI_URL),
			LanguageAPI:  getStatusCode(LANGUAGEAPI_URL),
			CountriesAPI: getStatusCode(COUNTRIESAPI_URL),
			Version:      VERSION,
			Uptime:       time.Since(StartTime).Seconds()}

		encodeJSON(w, statusOutput)

	}
}

func getStatusCode(url string) int {
	res, err := httpClient.Get(url)
	if err != nil {
		log.Println("Error when requesting remote endpoint: ", err.Error())
		return http.StatusServiceUnavailable
	}

	return res.StatusCode
}
