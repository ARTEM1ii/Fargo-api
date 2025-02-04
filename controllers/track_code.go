package controllers

import (
	"fargo-api/database"
	"fargo-api/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func CreateTrackcode(c echo.Context) error {
	track := new(models.TrackCode)

	if err := c.Bind(track); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var client models.Client
	if err := database.DB.Where("unique_code = ?", track.ClientID).First(&client).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Client not found"})
	}

	database.DB.Create(track)
	return c.JSON(http.StatusCreated, track)
}

func GetTrackCodes(c echo.Context) error {

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1 
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 20
	}

	offset := (page - 1) * limit

	var tracks []models.TrackCode
	result := database.DB.Limit(limit).Offset(offset).Find(&tracks)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch track codes"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"page":  page,
		"limit": limit,
		"data":  tracks,
	})
}

func UpdateTrackCodeStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var track models.TrackCode
	if err := database.DB.First(&track, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Track code not found"})
	}

	updatedData := new(models.TrackCode)
	if err := c.Bind(updatedData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	track.Status = updatedData.Status
	database.DB.Save(&track)
	
	return c.JSON(http.StatusOK, track)
}

func DeleteTrackCode(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var track models.TrackCode
	if err := database.DB.First(&track, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Track code not found"})
	}

	database.DB.Delete(&track)
	return c.JSON(http.StatusOK, map[string]string{"message": "Track code deleted successfully"})
}