package middlewares

import (
	"net/http"
	"strings"

	"recruitment-system/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Missing or invalid token")
		}

		claims := &CustomClaims{}
		_, err := jwt.ParseWithClaims(strings.TrimPrefix(token, "Bearer "), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("userID", claims.Subject)
		c.Set("userType", claims.UserType) // Use claims.UserType instead of claims["user_type"]
		return next(c)
	}
}
