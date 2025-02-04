package routes

import (
	"fargo-api/controllers"
	"fargo-api/middleware"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Echo) {
	e.POST("/login", controllers.AdminLogin)
	e.GET("/clients", controllers.GetClients, middleware.JWTMiddleware)
	e.DELETE("/clients/:id", controllers.DeleteClient, middleware.JWTMiddleware)
	e.POST("/register", controllers.RegisterAdmin)	
}

func CompanyContactRoutes(e *echo.Echo) {
	e.POST("/company_contacts", controllers.CreateCompanyContact, middleware.JWTMiddleware)
	e.GET("/company_contacts", controllers.GetAllCompanyContacts, middleware.JWTMiddleware)
	e.DELETE("/company_contacts/:id", controllers.DeleteCompanyContact, middleware.JWTMiddleware)
	e.PATCH("/company_contacts/:id", controllers.UpdateCompanyContact, middleware.JWTMiddleware)
}

func TrackCodeRoutes(e *echo.Echo) {
	e.POST("/track_codes", controllers.CreateTrackcode, middleware.JWTMiddleware)
	e.GET("/track_codes", controllers.GetTrackCodes, middleware.JWTMiddleware)
	e.PATCH("/track_codes/:id", controllers.UpdateTrackCodeStatus, middleware.JWTMiddleware)
	e.DELETE("/track_codes/:id", controllers.DeleteTrackCode, middleware.JWTMiddleware)
}