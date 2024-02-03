package http_server

import (
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

	http.HandleFunc("/status", DiagnosticHandler)

	log.Println("Starting http_server on port " + port + "...")
	StartTime = time.Now()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}