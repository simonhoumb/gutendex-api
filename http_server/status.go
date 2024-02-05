package http_server

type Status struct {
	Gutendexapi  string  `json:"gutendexapi"`
	Languageapi  string  `json:"languageapi"`
	Countriesapi string  `json:"countriesapi"`
	Version      string  `json:"version"`
	Uptime       float64 `json:"uptime"`
}
