package controllers

import (
	"fargo-api/database"
	"fargo-api/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAdmin(c echo.Context) error {
	admin := new(models.Admin)

	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)

	admin.Password = string(hashedPassword)

	database.DB.Create(admin)
	return c.JSON(http.StatusCreated, admin)
}

func AdminLogin(c echo.Context) error {

	input := new(models.Admin)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var admin models.Admin
	if err := database.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token, err := generateAdminJWT(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func generateAdminJWT(admin models.Admin) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": admin.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}