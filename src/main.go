package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	dto "github.com/kolaczyn/mann_backend/src/dto"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func ExampleProductHandler(w http.ResponseWriter, r *http.Request) {
	dummyProduct := dto.DummyProduct()

	jsonResponse, err := json.Marshal(dummyProduct)
	if err != nil {
		err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func main() {
	portNumber := 8000
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/12", ExampleProductHandler)

	fmt.Printf("Listening on port %d\n", portNumber)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portNumber), r))
}
