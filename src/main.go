package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aigaki/products-api/src/dto"
	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func AllProductsHandler(w http.ResponseWriter, r *http.Request) {
	dummyProducts := [3]dto.ProductDto{
		{
			Name:   "Selfie Project",
			Price:  1.99,
			Id:     345594,
			ImgUrl: "https://ros.net.pl/GalleryImages/product_photos/360_350/253711_336096_360_350_83107.png"},
		{
			Name:   "korektor matujący jasny",
			Price:  8.99,
			Id:     336096,
			ImgUrl: "https://ros.net.pl/GalleryImages/product_photos/360_350/34791_45248_360_350_100793.png"},
		{
			Name:   "lubrykant",
			Price:  9.79,
			Id:     167244,
			ImgUrl: "https://ros.net.pl/GalleryImages/product_photos/360_350/81270_251571_360_350.png"},
	}

	jsonResponse, err := json.Marshal(dummyProducts)
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
	r.HandleFunc("/products", AllProductsHandler)

	fmt.Printf("Listening on port %d\n", portNumber)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portNumber), r))
}
