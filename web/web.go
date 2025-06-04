package web

import (
	"encoding/json"
	"endpointMonitoringService/internal"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/cors"
)

type Server struct {
	EndpointController *internal.EndpointController
}

func NewServer(endpointController *internal.EndpointController) *Server {
	return &Server{EndpointController: endpointController}
}

func (s *Server) Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.rootHandler)
	mux.HandleFunc("/api/statuses", s.statusHandler)
	mux.HandleFunc("/api/targets", s.targetsHandler)
	mux.HandleFunc("/api/targets/create", s.createTargetHandler)

	// Add CORS middleware
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)

	http.ListenAndServe(":8000", handler)
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`Welcome! I am a backend for the endpoint monitoring service.`)
}

func (s *Server) createTargetHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)

	var payload internal.EndpointRecord
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println("Unmarshal error:", err)
		return
	}

	err := s.EndpointController.CreateEndpoints(payload)
	if err != nil {
		http.Error(w, "Failed to create target", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"message": "Endpoint created successfully"}`)
}

func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	statuses := s.EndpointController.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (s *Server) targetsHandler(w http.ResponseWriter, r *http.Request) {
	targets := s.EndpointController.GetEndpointCache()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}
