package internal

import (
	"log"
	"sync"
)

type EndpointStatus struct {
	Endpoint string `json:"endpoint"`
	Status   string `json:"status"`
}

type EndpointLabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EndpointRecord struct {
	Endpoint string                   `json:"endpoint"`
	Labels   map[string]EndpointLabel `json:"labels"`
}

type EndpointController struct {
	statuses []EndpointStatus
	mu       sync.Mutex
}

func NewEndpointController() *EndpointController {
	return &EndpointController{
		statuses: []EndpointStatus{},
	}
}

var inspectionCount = 0

var endpointRecordCache = []EndpointRecord{} // This is your cache of endpoint records

func (endpointController *EndpointController) CheckHealth(endpoints []string) {
	var wg sync.WaitGroup
	endpointController.statuses = []EndpointStatus{}
	results := make(chan EndpointStatus)
	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string) {
			checkEndpoint(endpoint, results, &wg)
		}(endpoint)
	}

	// Make wg.Wait a goroutine to prevent unbuffered channel deadlock
	go func() {
		wg.Wait()      // wait for all go routines to finish by wg.Done()
		close(results) // close channel
	}()

	for status := range results {
		endpointController.statuses = append(endpointController.statuses, status)
	}

	inspectionCount += 1
	log.Printf("Inspection Count: %d", inspectionCount)
}

func (endpointController *EndpointController) GetStatuses() []EndpointStatus {
	endpointController.mu.Lock()
	defer endpointController.mu.Unlock()
	return endpointController.statuses
}

func (endpointController *EndpointController) RefreshEndpoints() {
	log.Println("TO-DO: Implement Cache Refresh Logic")
}

func (endpointController *EndpointController) GetEndpointCache() []EndpointRecord {
	log.Println("TO-DO: Implement Cache Fetching Logic")
	return endpointRecordCache
}

func (endpointController *EndpointController) CreateEndpoints() {
	log.Println("TO-DO:Implement Endpoint Creation Logic ")
}
