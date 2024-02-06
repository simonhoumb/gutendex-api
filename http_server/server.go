package http_server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var StartTime time.Time

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Defaulting to 8080")
		port = "8080"
	}

	http.HandleFunc(STATUS_PATH, StatusHandler)
	http.HandleFunc(READERSHIP_PATH, ReadershipHandler)

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
