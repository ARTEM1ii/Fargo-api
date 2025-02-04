package controllers

import (
	"fargo-api/database"
	"fargo-api/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func CreateCompanyContact(c echo.Context) error {
	contact := new(models.CompanyContact)

	if err := c.Bind(contact); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	database.DB.Create(contact)
	return c.JSON(http.StatusAccepted, contact)
}

func GetAllCompanyContacts(c echo.Context) error {
	var contacts []models.CompanyContact
	database.DB.Find(&contacts)
	return(c.JSON(http.StatusOK, contacts))
}

func DeleteCompanyContact(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	
	var contact models.CompanyContact
	if err := database.DB.First(&contact, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Company contact not found"})
	}

	database.DB.Delete(&contact)
	return c.JSON(http.StatusOK, map[string]string{"message": "Company contact deleted successfully"})
}

func UpdateCompanyContact(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var contact models.CompanyContact
	if err := database.DB.First(&contact, id).Error; err != nil {
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Company contact not found"})
	}

	updatedData := new(models.CompanyContact)
	if err := c.Bind(updatedData); err != nil {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	
	database.DB.Model(&contact).Updates(updatedData)
	database.DB.First(&contact, id)
	
	return c.JSON(http.StatusOK, contact)
}