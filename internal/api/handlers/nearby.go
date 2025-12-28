package handlers

import (
	"encoding/json"
	"net/http"
	"newsApp/internal/service"
	"strconv"
)

func NearbyHandler(w http.ResponseWriter, r *http.Request) {
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, _ := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)

	articles, err := service.GetNearby(r.Context(), MongoDB, lat, lng)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
