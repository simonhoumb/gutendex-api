package http_server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var startTime time.Time
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
	startTime = time.Now()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func encodeJSON[T any](w http.ResponseWriter, v *T) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&v); err != nil {
		http.Error(w, "Error while encoding to json.", http.StatusInternalServerError)
		log.Println("Error while encoding to json: ", err.Error())
	}
}

func decodeJSON[T any](w http.ResponseWriter, res *http.Response, v *T) {
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&v); err != nil {
		http.Error(w, "Error while decoding json: ", http.StatusInternalServerError)
		log.Println("Error while decoding json: ", err.Error())
	}
}
