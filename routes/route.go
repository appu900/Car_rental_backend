package routes

import (
	"github.com/appu900/carrental/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1/")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Backend is uppp",
		})
	})

	api.Post("/user/register", handler.RegisterUser)
	api.Post("/user/auth", handler.LoginUser)
}
