package routes

import (
	"recruitment-system/controllers"
	"recruitment-system/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)
	e.POST("/uploadResume", controllers.UploadResume, middlewares.AuthMiddleware)
	e.POST("/admin/job", controllers.CreateJob, middlewares.AuthMiddleware)
	e.GET("/admin/job/:job_id", controllers.GetJob, middlewares.AuthMiddleware)
	e.GET("/admin/applicants", controllers.GetApplicants, middlewares.AuthMiddleware)
	e.GET("/admin/applicant/:applicant_id", controllers.GetApplicant, middlewares.AuthMiddleware)
	e.GET("/jobs", controllers.GetJobs)
	e.GET("/jobs/apply", controllers.ApplyJob, middlewares.AuthMiddleware)
}
