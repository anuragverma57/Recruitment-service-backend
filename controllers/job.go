package controllers

import (
	"context"
	"fmt"
	"net/http"
	"recruitment-system/models"
	"recruitment-system/repositories"

	"github.com/labstack/echo/v4"
)

func CreateJob(c echo.Context) error {
	job := models.Job{}
	if err := c.Bind(&job); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userType := c.Get("userType").(string)

	if userType == "Applicant" {
		return c.JSON(http.StatusInternalServerError, "Applicant can not create job")
	}

	userID := c.Get("userID").(string)
	fmt.Print(userID)
	job.PostedBy = repositories.GetUserByID(context.TODO(), userID)

	err := repositories.CreateJob(context.TODO(), job)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not create job")
	}

	return c.JSON(http.StatusCreated, "Job created successfully")
}

func GetJobs(c echo.Context) error {
	jobs, err := repositories.GetAllJobs(context.TODO())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not fetch jobs")
	}

	return c.JSON(http.StatusOK, jobs)
}

func GetJob(c echo.Context) error {
	jobID := c.Param("job_id")
	fmt.Println("jobID", jobID)
	job, err := repositories.GetJobByID(context.TODO(), jobID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Job not found")
	}

	return c.JSON(http.StatusOK, job)
}

func ApplyJob(c echo.Context) error {
	userID := c.Get("userID").(string)
	jobID := c.QueryParam("job_id")

	application := models.Application{
		UserID: userID,
		JobID:  jobID,
	}

	err := repositories.ApplyForJob(context.TODO(), application)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not apply for the job")
	}

	return c.JSON(http.StatusOK, "Application submitted successfully")
}
