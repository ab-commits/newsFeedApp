package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/service"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	articles, err := service.SearchArticles(r.Context(), MongoDB, query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
