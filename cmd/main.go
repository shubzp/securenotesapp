package main

import (
	"securenotesapp.com/internal/service"
)

func main() {
	newService := service.MakeService()
	newService.CreateServer()
	// newService.HealthCheck()
}
