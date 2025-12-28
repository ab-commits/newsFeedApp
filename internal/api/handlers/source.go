package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/service"
)

func SourceHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("name")

	articles, err := service.GetBySource(r.Context(), MongoDB, source)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
