package main

import (
	"fmt"

	"securenotesapp.com/internal/service"
)

func main() {
	fmt.Println("Hello there")
	newService := service.MakeService()
	newService.CreateServer()
	// newService.HealthCheck()
}
