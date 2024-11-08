package controllers

import (
	"context"
	"net/http"
	"recruitment-system/repositories"

	"github.com/labstack/echo/v4"
)

func GetApplicants(c echo.Context) error {
	applicants, err := repositories.GetAllApplicants(context.TODO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not fetch applicants")
	}

	return c.JSON(http.StatusOK, applicants)
}

func GetApplicant(c echo.Context) error {
	applicantID := c.Param("applicant_id")
	applicant, err := repositories.GetApplicantByID(context.TODO(), applicantID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Applicant not found")
	}

	return c.JSON(http.StatusOK, applicant)
}
