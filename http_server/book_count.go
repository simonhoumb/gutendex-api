package http_server

import (
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
	if r.Method == http.MethodGet {
		var bookCountOutput []BookCount
		w.Header().Add("Content-Type", "application/json")

		languageParameter := r.URL.Query().Get("language")
		if languageParameter != "" {

			languageCodes := strings.Split(languageParameter, ",")
			for _, code := range languageCodes {
				//add function to find data corresponding with code
				if isValidLanguageCode(code) {
					bookCountOutput = append(bookCountOutput, BookCount{
						Language: code,
						Books:    321,
						Authors:  14,
						Fraction: 0.0000023})
				}
			}
			if bookCountOutput != nil {
				encodeJSON(w, bookCountOutput)
			} else {
				invalidLanguageCodeError(w)
			}

		} else {
			invalidLanguageCodeError(w)
		}

	}
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
