package main

import (
	"github.com/adijha/omicorn/models"
	"github.com/adijha/omicorn/routes"
)

// Entrypoint for app.
func main() {
	// Load the routes
	r := routes.SetupRouter()

	// Initialize database
	models.SetupDatabase()

	// Start the HTTP API
	r.Run()
}
