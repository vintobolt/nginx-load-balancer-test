package main

import (
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.StatusText(200)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Server", "1")
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
