package internal

import (
	"fmt"
	"log"
)

type EndpointController struct {
}

func NewEndpointController() *EndpointController {
	return &EndpointController{}
}

var inspectionCount = 0
var refreshCount = 0

func (endpointController *EndpointController) RefreshStatuses() {
	err := refreshStatuses()
	if err != nil {
		log.Printf("Error refreshing statuses: %v", err)
		return
	}
	inspectionCount += 1
	log.Printf("Inspection Count: %d", inspectionCount)
}

func (endpointController *EndpointController) GetStatuses() []EndpointStatus {
	statuses := fetchStatusList()
	log.Println("Returned endpoint statuses.")
	return statuses
}

func (endpointController *EndpointController) RefreshEndpoints() {
	err := refreshCache()
	if err != nil {
		log.Printf("Error refreshing endpoints: %v", err)
		return
	}
	refreshCount += 1
	log.Printf("Refresh Count: %d", refreshCount)
}

func (endpointController *EndpointController) GetEndpointCache() []EndpointCacheRecord {
	endpointCache := fetchCache()
	if len(endpointCache) != 0 {
		log.Println("Returned endpoint cache.")
		return endpointCache
	}
	log.Println("Empty endpoint cache returned.")
	return []EndpointCacheRecord{}
}

func (endpointController *EndpointController) CreateEndpoints(payload EndpointRecord) error {
	fmt.Println("Creating Endpoint with payload: ", payload)
	if err := saveEndpointsToMemory(payload); err != nil {
		log.Printf("Error creating endpoint: %v", err)
	}
	return nil
}
