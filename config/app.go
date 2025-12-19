package config

import (
	"backenduas_sistemprestasi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "backenduas_sistemprestasi/docs"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New(LoggerConfig()))

	// swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// routes
	routes.UsersRoute(app)
	routes.AuthRoute(app)
	routes.AchievementRoutes(app)
	routes.LecturerRoute(app)
	routes.StudentRoutes(app)
	routes.Analytics(app)

	return app
}
