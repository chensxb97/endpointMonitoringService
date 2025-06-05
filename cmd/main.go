package main

import (
	"endpointMonitoringService/internal"
	"endpointMonitoringService/web"
	"time"
)

func main() {
	endpointController := internal.NewEndpointController()

	// Additionally, we can define a DB client here.
	// For this project, I'm keeping it simple and using an in-memory cache to simulate the database.

	go func() {
		ticker := time.NewTicker(5 * time.Minute) // Refresh every 5 minutes
		defer ticker.Stop()
		for range ticker.C {
			go endpointController.RefreshEndpoints() // Refresh endpoint cache
			go endpointController.RefreshStatuses()  // Refresh endpoint probe statuses
		}
	}()

	// Start web server
	server := web.NewServer(endpointController) // Setup server
	server.Start()
}
