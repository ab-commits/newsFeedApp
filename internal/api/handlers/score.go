package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/service"
)

func ScoreHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := service.GetByScore(r.Context(), MongoDB, 0.7)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
