package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/status/{code:[0-9]+}", StatusEchoHandler)

	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	body := "<html><head><title>Crashr</title></head><body> <h3> Click a link to return that status code: </h3>"

	for _, statusCode := range []int{100, 101, 200, 201, 202, 203, 204, 205, 206, 300, 301, 302, 303, 304, 305, 307, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 500, 501, 502, 503, 504, 505} {
		body += fmt.Sprintf(`<br><a href="/status/%d">/status/%d</a>`, statusCode, statusCode)
	}

	body += "</body></html>"
	fmt.Fprintln(w, body)
}

func StatusEchoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	statusCodeStr := vars["code"]
	statusCode, _ := strconv.Atoi(statusCodeStr)

	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf("Returning status code: %d", statusCode)))
}
