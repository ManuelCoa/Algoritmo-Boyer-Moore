package handlers

import (
	algorithms "boyer-moore-service/algoritmo"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	Text   string `json:"text"`
	Pattern string `json:"pattern"`
}

type SearchResponse struct {
	Indices []int `json:"indices"` // Índices donde se encuentra el patrón
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	indices := algorithms.BoyerMoore(req.Text, req.Pattern)
	response := SearchResponse{Indices: indices}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
