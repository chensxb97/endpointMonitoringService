package internal

import (
	"fmt"
	"sync"
)

type EndpointLabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EndpointRecord struct {
	Endpoint string          `json:"endpoint"`
	Module   string          `json:"module"`
	Labels   []EndpointLabel `json:"labels"`
}

var endpointsInMemory = []EndpointRecord{}

var inMemoryMutex sync.Mutex

func saveEndpointsToMemory(payload EndpointRecord) error {
	inMemoryMutex.Lock()
	defer inMemoryMutex.Unlock()

	// push endpint to in-memory storage
	newEndpoint := EndpointRecord{
		Endpoint: payload.Endpoint,
		Module:   payload.Module,
		Labels:   payload.Labels,
	}
	endpointsInMemory = append(endpointsInMemory, newEndpoint)
	fmt.Printf("Saved endpoint to memory: %v", newEndpoint.Endpoint)
	return nil
}

func fetchEndpointsMemory() []EndpointRecord {
	inMemoryMutex.Lock()
	defer inMemoryMutex.Unlock()
	return endpointsInMemory
}
