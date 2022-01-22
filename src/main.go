package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
