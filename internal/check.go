package internal

import (
	"log"
	"sync"
)

func checkProbeStatus(endpoint string) (int, error) {
	// TO-DO: Replace with Prometheus Query calls
	return 1, nil
}

func checkEndpoint(endpoint string, results chan<- EndpointStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	value, err := checkProbeStatus(endpoint)
	status := "down"
	if err == nil && value == 1 {
		status = "live"
	}

	results <- EndpointStatus{
		Endpoint: endpoint,
		Status:   status,
	}
	log.Printf("Checked %s: %s\n", endpoint, status)
}
