package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/service"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("name")

	articles, err := service.GetByCategory(r.Context(), MongoDB, category)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
