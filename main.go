package main

import (
	"log"
	"fargo-api/database"
	"fargo-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"go mod tidy
)

func main() {
	database.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Можно указать конкретные домены
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	routes.AdminRoutes(e)
	routes.CompanyContactRoutes(e)
	routes.TrackCodeRoutes(e)

	log.Println("🚀 Server running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}