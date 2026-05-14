package api

import (
	"encoding/json"
	"net/http"

	"synthview/internal/metrics"
)

func StartServer(m metrics.Metrics) {

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {

		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		json.NewEncoder(w).Encode(m)
	})

	http.ListenAndServe(":8080", nil)
}

