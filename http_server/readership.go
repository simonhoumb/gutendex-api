package http_server

import (
	"net/http"
)

type Readership struct {
	Country    string `json:"country"`
	IsoCode    string `json:"isocode"`
	Books      int    `json:"books"`
	Authors    int    `json:"authors"`
	Readership int    `json:"readership"`
}

func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "application/json")

		readershipOutput := Readership{
			Country:    "Norway",
			IsoCode:    "NO",
			Books:      211,
			Authors:    14,
			Readership: 5400000 / 211}

		encodeJSON(w, readershipOutput)

	}
}
