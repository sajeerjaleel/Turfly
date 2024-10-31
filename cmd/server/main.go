package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sajeerjaleel/Turfly/internal/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bookings", handlers.GetBookingsHandler).Methods("GET")

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
