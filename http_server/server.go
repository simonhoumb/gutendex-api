package http_server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var StartTime time.Time
var httpClient = &http.Client{Timeout: 3 * time.Second}

func Start() {
	defer httpClient.CloseIdleConnections()
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Defaulting to 8080")
		port = DEFAULT_PORT
	}

	http.HandleFunc(ROOT_PATH, RootHandler)
	http.HandleFunc(STATUS_PATH, StatusHandler)
	http.HandleFunc(READERSHIP_PATH, ReadershipHandler)
	http.HandleFunc(BOOKCOUNT_PATH, BookCountHandler)

	log.Println("Starting HTTP Server on port " + port + "...")
	StartTime = time.Now()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func encodeJSON(w http.ResponseWriter, v any) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	if err != nil {
		log.Println("Error during JSON encoding: ", err.Error())
		http.Error(w, "Error during JSON encoding.", http.StatusInternalServerError)
	}
}

func decodeJSON(w http.ResponseWriter, res *http.Response, v any) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&v); err == io.EOF {
		return
	} else if err != nil {
		log.Println("Error during JSON decoding: ", err.Error())
		http.Error(w, "Error during JSON decoding.", http.StatusInternalServerError)
	}
}
