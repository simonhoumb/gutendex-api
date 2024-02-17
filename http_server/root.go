package http_server

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Add("Content-Type", "text/html")

		output := fmt.Sprintf(`
Welcome to librarystats! <br>
<a href="%s">To Book Count</a> 
<input type="text" id="bookCountText" name="bookCountText"> <br>
<a href="%s">To Readership</a> 
<input type="text" id="readershipLanguage" name="readershipLanguage"> 
<input type="text" id="readershipLimit" name="readershipLimit"> <br>
<a href="%s">To Status</a>`,
			BOOKCOUNT_PATH, READERSHIP_PATH, STATUS_PATH)

		_, err := fmt.Fprintf(w, "%v", output)
		if err != nil {
			log.Println("Error when trying to output body: ", err.Error())
			http.Error(w, "Error with body.", http.StatusInternalServerError)
		}

	}
}
