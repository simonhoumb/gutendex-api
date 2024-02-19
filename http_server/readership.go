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
	urlParts := strings.Split(path, "/")

	var languageParameter string

	if len(urlParts) >= 5 {
		languageParameter = urlParts[4]
	}
	if validReadershipParameter(languageParameter) {

		res := utils.GetResults(w, httpClient, LANGUAGEAPI_URL+languageParameter)
		var countries []language2countries.Country
		utils.DecodeJSON(w, res, &countries)

		limit := r.URL.Query().Get("limit")
		countriesToShow := determineCountriesToShow(w, limit, countries)

		var readershipOutput []Readership
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
		utils.EncodeJSON(w, &readershipOutput)
	} else {
		http.Error(w, "No valid language code provided.", http.StatusBadRequest)
		log.Println("No valid language code provided")
	}
}

func determineCountriesToShow(w http.ResponseWriter, limitParam string, countries []language2countries.Country) int {
	limit, err := parseLimit(limitParam, countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0 // Or handle the error differently if needed
	}

	return int(math.Min(float64(len(countries)), float64(limit)))
}

func parseLimit(limitParam string, listOfCountries []language2countries.Country) (int, error) {
	if limitParam == "" {
		return len(listOfCountries), nil
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, fmt.Errorf("error parsing limit parameter: %s", limitParam)
	}

	return limit, err
}

func validReadershipParameter(parameter string) bool {
	return parameter != "" && len(parameter) == 2
}
