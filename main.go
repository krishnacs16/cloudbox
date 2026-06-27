package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	addr string
}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type TestResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := TestResponse{
		Message: "test endpoint working",
		Data:    map[string]string{"environment": "testing"},
	}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) Start() error {
	http.HandleFunc("/health", s.healthHandler)
	http.HandleFunc("/test", s.testHandler)

	log.Printf("Starting test server on %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

func main() {
	server := NewServer(":8080")
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}