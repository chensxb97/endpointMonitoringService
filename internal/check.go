package internal

import (
	"fmt"
	"log"
	"sync"
)

type EndpointStatus struct {
	Endpoint string          `json:"endpoint"`
	Labels   []EndpointLabel `json:"labels"`
	Status   string          `json:"status"`
}

var statusList = []EndpointStatus{}

var statusMutex sync.Mutex

func checkProbeStatus(endpoint string) (int, error) {
	// TO-DO: Replace with metric results from Prometheus Queries
	return 1, nil
}

func checkEndpoint(endpoint EndpointRecord, results chan<- EndpointStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	value, err := checkProbeStatus(endpoint.Endpoint)
	status := "down"
	if err == nil && value == 1 {
		status = "live"
	}

	results <- EndpointStatus{
		Endpoint: endpoint.Endpoint,
		Labels:   endpoint.Labels,
		Status:   status,
	}
	log.Printf("Checked %s: %s\n", endpoint, status)
}

func refreshStatuses() error {
	statusMutex.Lock()
	defer statusMutex.Unlock()

	var wg sync.WaitGroup
	results := make(chan EndpointStatus)
	endpointRecordsInMemory := fetchEndpointsMemory()

	for _, endpointRecord := range endpointRecordsInMemory {
		wg.Add(1)
		go func(endpointRecord EndpointRecord) {
			checkEndpoint(endpointRecord, results, &wg)
		}(endpointRecord)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	statusList = make([]EndpointStatus, len(results)) // reset statuses
	for status := range results {
		statusList = append(statusList, status)
	}

	fmt.Println("Inspected statuses: ", statusList)

	return nil
}

func fetchStatusList() []EndpointStatus {
	statusMutex.Lock()
	defer statusMutex.Unlock()
	return statusList
}
