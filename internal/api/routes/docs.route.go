package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Swagger(a *fiber.App) {
	a.Get("/docs/*", swagger.HandlerDefault)
}
