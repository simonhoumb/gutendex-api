package http_server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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

		path := r.URL.Path
		parts := strings.Split(path, "/")

		var languageCode string

		if len(parts) >= 5 {
			languageCode = parts[4]
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			log.Println("Value of parameter 'limit' is not valid: ", err.Error())
		}

		if limit == 0 {
			limit = 1
		}

		if languageCode != "" && len(languageCode) == 2 {
			for i := 0; i < limit; i++ {
				readershipOutput := Readership{
					Country:    "Norway",
					IsoCode:    languageCode,
					Books:      211,
					Authors:    14,
					Readership: 5400000 / 211}

				encodeJSON(w, readershipOutput)
			}
		} else {
			http.Error(w, "No language code provided.", http.StatusBadRequest)
		}

	}
}
