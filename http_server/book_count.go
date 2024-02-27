package http_server

import (
	"assignment-1/api/gutendex"
	"assignment-1/utils"
	"fmt"
	"net/http"
	"strconv"
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
		bookCountGetRequest(w, r)
	default:
		http.Error(w, fmt.Sprintf("The method %s is not implemented. Currently only %s is supported.",
			r.Method, http.MethodGet), http.StatusMethodNotAllowed)
	}
}

func bookCountGetRequest(w http.ResponseWriter, r *http.Request) {
	defer httpClient.CloseIdleConnections()
	w.Header().Add("Content-Type", "application/json")

	languageParameter := r.URL.Query().Get("language")
	fmt.Println("parameter: \"" + languageParameter + "\"")

	var bookCountOutput []BookCount
	languageCodes := strings.Split(languageParameter, ",")

	for _, code := range languageCodes {

		//refactor this into a function to get a result for one country (maybe move to gutendex.go)
		res := utils.GetResults(w, httpClient, GUTENDEXAPI_URL+"?languages="+code)
		var bookResults []gutendex.Book
		var books gutendex.Books
		utils.DecodeJSON(w, res, &books)

		if books.Next != "" {
			getAllPages(w, httpClient, books, bookResults)
		}

		bookCountOutput = append(bookCountOutput, BookCount{
			Language: code,
			Books:    books.Count,
			Authors:  numberOfAuthors(w, httpClient, books),
			Fraction: float64(books.Count) / float64(totalBookCount(w)),
		})
	}
	utils.EncodeJSON(w, &bookCountOutput)
}

func isValidLanguageCode(code string) bool {
	for _, c := range code {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return len(code) == 2
}

func totalBookCount(w http.ResponseWriter) int {
	res := utils.GetResults(w, httpClient, GUTENDEXAPI_URL)
	var books gutendex.Books
	utils.DecodeJSON(w, res, &books)

	return books.Count
}

func getAllPages(w http.ResponseWriter, client *http.Client, books gutendex.Books) {

}

func numberOfAuthors(w http.ResponseWriter, client *http.Client, books gutendex.Books) int {
	uniqueAuthors := make(map[string]bool)
	var authorList []string

	/*
		for _, book := range books.Results {
			for _, author := range book.Authors {
				if _, found := uniqueAuthors[author.Name]; !found && author.Name != "" {
					uniqueAuthors[author.Name] = true
					authorList = append(authorList, author.Name)
				}
			}
		}
	*/

	hasNext := true
	pageNum := 1
	for hasNext {
		for _, book := range books.Results {
			for _, author := range book.Authors {
				if _, found := uniqueAuthors[author.Name]; !found && author.Name != "" {
					uniqueAuthors[author.Name] = true
					authorList = append(authorList, author.Name)
				}
			}
		}
		fmt.Println("page: " + strconv.Itoa(pageNum))
		fmt.Println("next: " + books.Next)
		if books.Next != "" {
			res := utils.GetResults(w, client, books.Next)
			utils.DecodeJSON(w, res, &books)
			pageNum++

		} else {
			hasNext = false
		}
	}
	return len(uniqueAuthors)
}
