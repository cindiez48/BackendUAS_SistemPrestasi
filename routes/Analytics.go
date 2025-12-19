package routes

import (
	"backenduas_sistemprestasi/app/service"
	"backenduas_sistemprestasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func Analytics(app *fiber.App) {
	api := app.Group("/api/v1")
	reports := api.Group("/reports")

	reports.Use(middleware.Protect())
	reports.Use(middleware.HasPermission("analytics:read"))

	reports.Get("/statistics", service.GetStatisticsService)
	reports.Get("/student/:id", service.GetStudentReportService)

}
