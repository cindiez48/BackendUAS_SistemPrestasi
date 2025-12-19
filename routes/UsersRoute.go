package routes

import (
	"backenduas_sistemprestasi/app/service"
	"backenduas_sistemprestasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App) {

	api := app.Group("/api/v1")

	users := api.Group("/users")
	users.Use(middleware.Protect())
	users.Use(middleware.HasPermission("user:manage"))
	users.Get("/", service.GetAllUsers)
	users.Post("/", service.CreateUser)
	users.Get("/:id", service.GetUserByID)
	users.Put("/:id", service.UpdateUser)
	users.Delete("/:id", service.DeleteUser)
	users.Put("/:id/role", service.AssignRole)

}