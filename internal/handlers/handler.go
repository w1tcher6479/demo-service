package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/w1tcher6479/demo-service/internal/models"
)

func GetOrder(cacheData map[string]models.Order) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		result, ok := cacheData[id]
		if ok {
			jsonData, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write(jsonData)
		} else {
			http.Error(w, "data not found", http.StatusNotFound)
		}
	}
}
