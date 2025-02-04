package main

import (
	"log"
	"fargo-api/database"
	"fargo-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.AdminRoutes(e)
	routes.CompanyContactRoutes(e)
	routes.TrackCodeRoutes(e)

	log.Println("Server is running on http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
	
}