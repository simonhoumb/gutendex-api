package http_server

import (
	"assignment-1/api/gutendex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"
)

type BookCount struct {
	Language string  `json:"language"`
	Books    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}

func BookCountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, fmt.Sprintf("The method %s is not implemented. Currently only %s is supported.",
			r.Method, http.MethodGet), http.StatusMethodNotAllowed)
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	defer httpClient.CloseIdleConnections()
	w.Header().Add("Content-Type", "application/json")

	languageParameter := r.URL.Query().Get("language")
	fmt.Println("parameter: \"" + languageParameter + "\"")

	var bookCountOutput []BookCount
	languageCodes := strings.Split(languageParameter, ",")

	for _, code := range languageCodes {
		res, err := httpClient.Get(GUTENDEXAPI_URL + "?languages=" + code)
		if err != nil {
			log.Println("Error while getting response:", err.Error())
		}

		decoder := json.NewDecoder(res.Body)
		var books gutendex.Books
		if err := decoder.Decode(&books); err != nil {
			log.Println("Error while decoding from json: ", err.Error())
		}

		bookCountOutput = append(bookCountOutput, BookCount{
			Language: code,
			Books:    books.Count,
			Authors:  numberOfAuthors(books),
			Fraction: float64(books.Count) / float64(totalBookCount(w, r)),
		})
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&bookCountOutput); err != nil {
		log.Println("Error while encoding to json: ", err.Error())
	}
}

func getResults(query string) (*http.Response, error) {
	fmt.Println("Url used: " + GUTENDEXAPI_URL + query)
	result, err := httpClient.Get(GUTENDEXAPI_URL + query)
	return result, err
}

func isValidLanguageCode(code string) bool {
	for _, c := range code {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return len(code) == 2
}

func invalidLanguageCodeError(w http.ResponseWriter) {
	http.Error(w, "No valid language code provided.", http.StatusBadRequest)
}

func totalBookCount(w http.ResponseWriter, r *http.Request) int {
	res, err := httpClient.Get(GUTENDEXAPI_URL)
	if err != nil {
		log.Println("Error while getting response:", err.Error())
	}

	decoder := json.NewDecoder(res.Body)
	var books gutendex.Books
	if err := decoder.Decode(&books); err != nil {
		log.Fatal(err)
	}
	return books.Count
}

func numberOfAuthors(books gutendex.Books) int {
	uniqueAuthors := make(map[string]bool)
	var authorList []string
	for _, book := range books.Results {
		for _, author := range book.Authors {
			if _, found := uniqueAuthors[author.Name]; !found && author.Name != "" {
				uniqueAuthors[author.Name] = true
				authorList = append(authorList, author.Name)
			}
		}
	}
	return len(uniqueAuthors)
}
