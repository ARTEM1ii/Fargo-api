package middleware

import (
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token missing"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("admin_id", claims["admin_id"])
		}

		return next(c)
	}
}