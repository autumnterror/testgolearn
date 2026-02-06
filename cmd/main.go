package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// StatusResponse определяет структуру ответа для /status
type StatusResponse struct {
	Status string `json:"status"`
	System string `json:"system"`
}

// statusHandler обрабатывает запросы к /status
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := StatusResponse{Status: "OK", System: "Go Web App"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/status", statusHandler)
	log.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
