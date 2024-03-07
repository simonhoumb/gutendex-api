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
	var bookResults []gutendex.Book

	for _, code := range languageCodes {
		if isValidLanguageCode(code) {
			//refactor this into a function to get a result for one country (maybe move to gutendex.go)
			res := utils.GetResults(w, httpClient, GUTENDEXAPI_URL+"?languages="+code)
			var books gutendex.Books
			utils.DecodeJSON(w, res, &books)

			getAllPagesLoop(w, httpClient, &books, &bookResults)
			fmt.Printf("restult length: %d\n", len(bookResults))

			bookCountOutput = append(bookCountOutput, BookCount{
				Language: code,
				Books:    books.Count,
				Authors:  numberOfAuthors(w, httpClient, bookResults),
				Fraction: float64(books.Count) / float64(totalBookCount(w)),
			})
		}
	}
	if len(bookCountOutput) > 0 {
		utils.EncodeJSON(w, &bookCountOutput)
	} else {
		http.Error(w, "No valid parameters provided.", http.StatusBadRequest)
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

func totalBookCount(w http.ResponseWriter) int {
	res := utils.GetResults(w, httpClient, GUTENDEXAPI_URL)
	var books gutendex.Books
	utils.DecodeJSON(w, res, &books)

	return books.Count
}

func getAllPagesLoop(w http.ResponseWriter, client *http.Client, books *gutendex.Books, bookResults *[]gutendex.Book) {
	fmt.Println("entered GetAllPages....")
	*bookResults = append(*bookResults, *books.Results...)
	pageNum := 1
	for books.Next != nil {
		pageNum++
		fmt.Printf("Entered page: %d\n", pageNum)
		res := utils.GetResults(w, client, *books.Next)
		utils.DecodeJSON(w, res, &books)
		*bookResults = append(*bookResults, *books.Results...)
	}
}

func numberOfAuthors(w http.ResponseWriter, client *http.Client, bookResults []gutendex.Book) int {
	uniqueAuthors := make(map[string]bool)
	authorCount := 0

	for _, book := range bookResults {
		if book.Id == 28233 {
			fmt.Printf("found newton!! name:%s", book.Authors[0].Name)
		}
		for _, author := range book.Authors {
			authorCount++
			var birthYear string
			var deathYear string
			if author.BirthYear != 0 {
				birthYear = strconv.Itoa(author.BirthYear)
			}
			if author.DeathYear != 0 {
				deathYear = strconv.Itoa(author.DeathYear)
			}
			fmt.Printf("Name: %s, B.Y.: \"%s\", D.Y.: \"%s\"\n", author.Name, birthYear, deathYear)
			if _, found := uniqueAuthors[author.Name+birthYear+deathYear]; !found {

				uniqueAuthors[author.Name+birthYear+deathYear] = true
				fmt.Printf("found=%v\n", found)
				fmt.Println("added")
			} else {
				fmt.Printf("found=%v\n", found)
			}
		}
	}
	fmt.Printf("numOfResults: %d\n", len(bookResults))
	fmt.Printf("total authors found: %d\n", authorCount)
	return len(uniqueAuthors)
}
