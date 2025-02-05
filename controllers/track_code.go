package controllers

import (
	"fargo-api/database"
	"fargo-api/models"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"github.com/xuri/excelize/v2"
	"github.com/labstack/echo/v4"
)

const telegramBotToken = "7430261806:AAGKpzpzdSLg1wG6CBgYuJmH_eE4xD3HWpU"

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

	var client models.Client
	if err := database.DB.Where("unique_code = ?", track.ClientID).First(&client).Error; err == nil {
		telegramID := extractTelegramID(client.FullName)
		if telegramID != "" {
			sendTelegramNotification(telegramID, track.TrackCode, string(track.Status))
		}
	}

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

func sendTelegramNotification(telegramID, trackCode, status string) {
	message := fmt.Sprintf("📦 У вашего заказа с трек-кодом %s изменился статус на: %s.", trackCode, status)

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	data := url.Values{}
	data.Set("chat_id", telegramID)
	data.Set("text", message)

	_, err := http.PostForm(apiURL, data)
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения в Telegram:", err)
	}
}

func ExportTrackCodesToExcel(c echo.Context) error {
	var tracks []models.TrackCode
	result := database.DB.Find(&tracks)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch track codes"})
	}

	file := excelize.NewFile()
	sheetName := "Track Codes"
	file.SetSheetName("Sheet1", sheetName)

	headers := []string{"ID", "Client ID", "Track Code", "Status", "Created At"}
	for i, header := range headers {
		cell := string(rune('A' + i)) + "1"
		file.SetCellValue(sheetName, cell, header)
	}

	for i, track := range tracks {
		row := strconv.Itoa(i + 2) 
		file.SetCellValue(sheetName, "A"+row, track.ID)
		file.SetCellValue(sheetName, "B"+row, track.ClientID)
		file.SetCellValue(sheetName, "C"+row, track.TrackCode)
		file.SetCellValue(sheetName, "D"+row, string(track.Status)) 
		file.SetCellValue(sheetName, "E"+row, track.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=track_codes.xlsx")
	return file.Write(c.Response().Writer)
}

func extractTelegramID(fullName string) string {
	start := strings.Index(fullName, "(")
	end := strings.Index(fullName, ")")
	if start != -1 && end != -1 && start < end {
		return fullName[start+1 : end]
	}
	return ""
}
