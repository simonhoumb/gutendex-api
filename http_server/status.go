package http_server

import "time"

type Status struct {
	Gutendexapi, Languageapi, Countriesapi, Version string
	Uptime                                          time.Duration
}
