package controllers

import (
	"fargo-api/database"
	"fargo-api/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetClients(c echo.Context) error {

	var clients []models.Client
	result := database.DB.Find(&clients)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch clients"})
	}

	return c.JSON(http.StatusOK, clients)
}

func DeleteClient(c echo.Context) error {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid client ID"})
	}

	var client models.Client
	if err := database.DB.First(&client, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Client not found"})
	}

	database.DB.Delete(&client)

	return c.JSON(http.StatusOK, map[string]string{"message": "Client deleted successfully"})
}