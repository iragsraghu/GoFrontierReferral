package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Running...")
	})
	router.HandleFunc("/generate", GenerateLink).Methods("POST")
	router.HandleFunc("/listReferrals", ListReferrals).Methods("GET")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
