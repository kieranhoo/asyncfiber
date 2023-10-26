package routes

import (
	"qrcheckin/api/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	r := a.Group("/api/v1/auth")
	r.Post("/signup", controller.SignUp)
	r.Post("/signin", controller.SignIn)
}
