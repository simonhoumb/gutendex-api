package utils

import (
	"log"
	"net/http"
)

func GetResults(w http.ResponseWriter, client *http.Client, url string) *http.Response {
	result, err := client.Get(url)
	if err != nil {
		http.Error(w, "Error while trying to reach endpoint.", http.StatusInternalServerError)
		log.Println("Error while trying to reach endpoint: ", err.Error())
	}
	return result
}
