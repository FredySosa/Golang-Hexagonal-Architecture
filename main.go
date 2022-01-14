package main

import (
	"log"

	"github.com/FredySosa/hexArch/internal/core/domain"

	"github.com/FredySosa/hexArch/internal/adapters"
	"github.com/FredySosa/hexArch/internal/core/services"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	countriesRepository := adapters.NewCountryRepository(map[string]domain.Country{
		"MX": {
			ID:   "MX",
			Name: "Mexico",
		},
		"USA": {
			ID:   "USA",
			Name: "United States",
		},
	})
	countriesService := services.NewCountryService(countriesRepository)
	eventService := services.NewEventService(countriesService)
	echoHandler := adapters.NewHTTPHandler(eventService)
	echoHandler.Pre(middleware.RemoveTrailingSlash())
	echoHandler.Use(middleware.Logger())

	log.Fatal(echoHandler.Start(":8080"))
}
