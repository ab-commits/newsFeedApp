package main

import (
	"log"
	"net/http"
	"newsApp/internal/api/handlers"
	"newsApp/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to Mongo
	mongoDB, err := db.ConnectMongo("mongodb://localhost:27017", "news_db")
	if err != nil {
		log.Fatal(err)
	}

	// Set global MongoDB reference for handlers
	handlers.MongoDB = mongoDB

	// Create router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/api/news/category", handlers.CategoryHandler).Methods("GET")
	r.HandleFunc("/api/news/source", handlers.SourceHandler).Methods("GET")
	r.HandleFunc("/api/news/search", handlers.SearchHandler).Methods("GET")
	r.HandleFunc("/api/news/score", handlers.ScoreHandler).Methods("GET")
	r.HandleFunc("/api/news/nearby", handlers.NearbyHandler).Methods("GET")

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
