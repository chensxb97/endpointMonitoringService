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
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			go endpointController.RefreshEndpoints() // Refresh endpoint cache every 5 seconds
			go endpointController.RefreshStatuses()  // Refresh statuses every 5 seconds
		}
	}()

	// Start web server
	server := web.NewServer(endpointController) // Setup server with endpoint controller
	server.Start()
}
