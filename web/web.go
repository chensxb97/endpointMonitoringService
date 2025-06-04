package web

import (
	"encoding/json"
	"endpointMonitoringService/internal"
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
	mux.HandleFunc("/statuses", s.statusHandler)
	mux.HandleFunc("/targets/create", s.createTargetHandler)
	mux.HandleFunc("/targets", s.targetsHandler)

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
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var payload internal.EndpointRecord
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = s.EndpointController.CreateEndpoints(payload)
	if err != nil {
		http.Error(w, "Failed to create target", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "true",
		"message": "Target created successfully",
	})
}

func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	statuses := s.EndpointController.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (s *Server) targetsHandler(w http.ResponseWriter, r *http.Request) {
	endpointCache := s.EndpointController.GetEndpointCache()
	if len(endpointCache) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]interface{}{})
		return
	}

	// Converting the list of records to the appropriate HTTP SD payload format
	var targets []map[string]interface{}
	for _, endpoint := range endpointCache {
		target := map[string]interface{}{
			"targets": []string{endpoint.Endpoint},
			"labels":  endpoint.Labels,
		}
		targets = append(targets, target)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}
