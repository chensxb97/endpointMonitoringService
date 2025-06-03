package web

import (
	"encoding/json"
	"endpointMonitoringService/internal"
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
	mux.HandleFunc("/status", s.statusHandler)
	mux.HandleFunc("/targets", s.targetsHandler)
	mux.HandleFunc("/targets/create", s.createTargetHandler)

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`TO-DO: Implement target creation logic.`)
}

func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	statuses := s.EndpointController.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (s *Server) targetsHandler(w http.ResponseWriter, r *http.Request) {
	// TO-DO: Replace targets variable with actual targets from cache
	targets := []map[string]interface{}{
		{
			"targets": []string{"http://example.com"},
			"labels": map[string]string{
				"__meta_module":      "http_2xx",
				"__meta_application": "App 1",
				"__meta_team":        "Alpha",
				"__meta_environment": "prod",
				"__meta_datacenter":  "A",
			},
		},
		{
			"targets": []string{"http://localhost:8000"},
			"labels": map[string]string{
				"__meta_module":      "http_2xx",
				"__meta_application": "App 2",
				"__meta_team":        "Beta",
				"__meta_environment": "nonprod",
				"__meta_datacenter":  "B",
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}
