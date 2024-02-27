package utils

import (
	"encoding/json"
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

func EncodeJSON[T any](w http.ResponseWriter, v *T) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(v); err != nil {
		http.Error(w, "Error while encoding to json.", http.StatusInternalServerError)
		log.Println("Error while encoding to json: ", err.Error())
	}
}

func DecodeJSON[T any](w http.ResponseWriter, res *http.Response, v *T) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(v); err != nil {
		http.Error(w, "Error while decoding json: ", http.StatusInternalServerError)
		log.Println("Error while decoding json: ", err.Error())
	}
}
