package controller

import (
	"qrcheckin/internal/handler/service"
	"qrcheckin/pkg/sentry"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck.
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/healthcheck [GET]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

// HealthCheck.
// @Description health check worker.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/workercheck [GET]
func HealthCheckWorker(c *fiber.Ctx) error {
	go service.HealthCheck()
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

// Email.
// @Description test email.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/email [GET]
func Email(c *fiber.Ctx) error {
	_, err := service.Email()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

// Sentry.
// @Description test sentry.io.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/sentry [GET]
func Sentry(c *fiber.Ctx) error {
	panic("y tho")
}

// Capture.
// @Description test sentry.io capture.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/capture [GET]
func Capture(c *fiber.Ctx) error {
	sentry.CaptureMessage(c, "User provided unwanted query string, but we recovered just fine")
	return c.SendStatus(fiber.StatusOK)
}
