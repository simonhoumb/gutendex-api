package http_server

import (
	"assignment-1/api/language2countries"
	"assignment-1/utils"
	"log"
	"math"
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

		var languageParameter string
		var countriesToShow int

		if len(parts) >= 5 {
			languageParameter = parts[4]
		}

		res := utils.GetResults(w, httpClient, LANGUAGEAPI_URL+languageParameter)
		var countries []language2countries.Country
		decodeJSON(w, res, &countries)

		limit := r.URL.Query().Get("limit")
		if limit != "" {
			limit, err := strconv.Atoi(limit)
			if err != nil {
				http.Error(w, "Error when converting parameter.", http.StatusInternalServerError)
				log.Println("Error when converting parameter: ", err.Error())
			}
			countriesToShow = int(math.Min(float64(len(countries)), float64(limit)))
		} else {
			countriesToShow = len(countries)
		}

		var readershipOutput []Readership
		if languageParameter != "" && len(languageParameter) == 2 {
			for i := 0; i < countriesToShow; i++ {
				readershipOutput = append(readershipOutput, Readership{
					Country:    countries[i].OfficialName,
					IsoCode:    countries[i].Iso31661Alpha2,
					Books:      211,
					Authors:    14,
					Readership: 5400000 / 211})

			}
		} else {
			http.Error(w, "No language code provided.", http.StatusBadRequest)
		}

		encodeJSON(w, &readershipOutput)

	}
}
