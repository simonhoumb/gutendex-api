package http_server

import (
	"assignment-1/api/gutendex"
	"assignment-1/api/language2countries"
	"assignment-1/utils"
	"fmt"
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
	switch r.Method {
	case http.MethodGet:
		readershipGetRequest(w, r)
	default:
		http.Error(w, fmt.Sprintf("The method %s is not implemented. Currently only %s is supported.",
			r.Method, http.MethodGet), http.StatusMethodNotAllowed)
	}
}

func readershipGetRequest(w http.ResponseWriter, r *http.Request) {
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
	utils.DecodeJSON(w, res, &countries)

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

			//maybe refactor this
			bookRes := utils.GetResults(w, httpClient,
				GUTENDEXAPI_URL+"?languages="+countries[i].Iso31661Alpha2)

			var books gutendex.Books
			utils.DecodeJSON(w, bookRes, &books)

			readershipOutput = append(readershipOutput, Readership{
				Country:    countries[i].OfficialName,
				IsoCode:    countries[i].Iso31661Alpha2,
				Books:      books.Count,            //use same as inside for loop in book_count.go
				Authors:    numberOfAuthors(books), //use same as inside for loop in book_count.go
				Readership: 5400000})               //use rest countries api with isocode

		}
	} else {
		http.Error(w, "No language code provided.", http.StatusBadRequest)
	}

	utils.EncodeJSON(w, &readershipOutput)
}
