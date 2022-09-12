package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/show", showProduct)
	mux.HandleFunc("/show/create", createProduct)

	log.Println("zapusk serverom")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
