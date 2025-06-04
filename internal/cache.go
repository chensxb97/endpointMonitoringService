package internal

import (
	"fmt"
	"sync"
)

type EndpointCacheRecord struct {
	Endpoint string            `json:"endpoint"`
	Labels   map[string]string `json:"labels"`
}

var endpointCache = []EndpointCacheRecord{}
var cacheMutex sync.Mutex

func refreshCache() error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	recordsInMemory := fetchEndpointsMemory()

	endpointCache = make([]EndpointCacheRecord, len(recordsInMemory)) // reset cache
	for i, record := range recordsInMemory {
		endpointCache[i] = EndpointCacheRecord{
			Endpoint: record.Endpoint,
			Labels:   make(map[string]string),
		}

		for _, label := range record.Labels {
			endpointCache[i].Labels["__meta_"+label.Key] = label.Value
		}
	}

	fmt.Println("Refreshed cache: ", endpointCache)

	return nil
}

func fetchCache() []EndpointCacheRecord {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	return endpointCache
}
