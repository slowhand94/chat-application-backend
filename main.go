package main

import (
	"slowhand94/chat-application-backend/router"
)

/*
 * entry point for backend service
 */
func main() {
	// create router and setup routes
	router.CreateRouter()
	router.SetupRoutes()

	// start server on port :8080
	router.StartRouter()
}
