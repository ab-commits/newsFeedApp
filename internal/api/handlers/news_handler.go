package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/llm"
	"newsApp/internal/service"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDB *mongo.Database

func SearchNewsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, _ := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)

	llmResult, err := llm.AnalyzeQuery(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	articles, err := service.GetRelevantArticles(r.Context(), MongoDB, llmResult, lat, lng)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
