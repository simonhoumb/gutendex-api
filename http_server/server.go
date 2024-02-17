package http_server

import (
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
