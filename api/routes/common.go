package routes

import (
	"asyncfiber/api/controller"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(a *fiber.App) {
	r := a.Group("/common")
	r.Get("/healthcheck", controller.HealthCheck)
	r.Get("/workercheck", controller.HealthCheckWorker)
	r.Get("/email", controller.Email)
	r.Get("/sentry", controller.Sentry)
	r.Get("/capture", controller.Capture)
}
