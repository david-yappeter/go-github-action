package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var hitted = 0

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", handleDefaultRoute).Methods(http.MethodGet)
	router.HandleFunc("/hit", handleHit).Methods(http.MethodGet)

	log.Printf("Listen and Serve at http://localhost:%s\n", port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func handleDefaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleHit(w http.ResponseWriter, r *http.Request) {
	hitted++
	fmt.Fprintf(w, "I have been hitted %d times", hitted)
}
