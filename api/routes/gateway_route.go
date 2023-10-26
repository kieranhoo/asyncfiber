package routes

import "github.com/gofiber/fiber/v2"

func Gateway(a *fiber.App) {
	AuthRoutes(a)
	HealthCheck(a)
	Swagger(a)
}
