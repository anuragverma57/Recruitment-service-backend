package controllers

import (
	"context"
	"net/http"
	"recruitment-system/config"
	"recruitment-system/models"
	"recruitment-system/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	UserType string `json:"user_type"`
	UserID   string `json:"userId"`
	jwt.StandardClaims
}

func Signup(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)

	err := repositories.CreateUser(context.TODO(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Could not create user",
			"success": false,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User  created successfully",
		"success": true,
	})
}

func Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	storedUser, err := repositories.GetUserByEmail(context.TODO(), user.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(user.PasswordHash)) != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := CustomClaims{
		UserType: storedUser.UserType,
		UserID:   storedUser.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   storedUser.ID,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}
