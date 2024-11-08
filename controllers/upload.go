// controllers/upload.go
package controllers

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"recruitment-system/models"
	"recruitment-system/repositories"

	"github.com/labstack/echo/v4"
)

func UploadResume(c echo.Context) error {
	userID := c.Get("userID").(string)

	// Get the file from the request
	file, err := c.FormFile("resume")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "No file is received")
	}

	// Save the file locally
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to open the file")
	}
	defer src.Close()

	// Create a path to save the file
	dst, err := os.Create(filepath.Join("uploads", file.Filename))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to create the file")
	}
	defer dst.Close()

	// Copy the file to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to save the file")
	}

	// Save the resume file path in the database
	profile := models.Profile{
		Applicant:         models.User{ID: userID},
		ResumeFileAddress: filepath.Join("uploads", file.Filename),
	}

	err = repositories.SaveProfile(context.TODO(), profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not save profile")
	}

	return c.JSON(http.StatusOK, "Resume uploaded successfully")
}
