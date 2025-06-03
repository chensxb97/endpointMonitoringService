package main

import (
	"endpointMonitoringService/internal"
	"endpointMonitoringService/web"
	"time"
)

func main() {
	endpointController := internal.NewEndpointController()
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			go endpointController.RefreshEndpoints()
		}
	}()

	// Start web server
	server := web.NewServer(endpointController)
	server.Start()
}
