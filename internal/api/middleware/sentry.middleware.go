package middleware

import (
	"asyncfiber/internal/config"

	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
)

func SentryMiddleware(a *fiber.App) {
	if config.StageStatus == "prod" && config.SentryDsn != "" {
		a.Use(fibersentry.New(fibersentry.Config{
			Repanic:         true,
			WaitForDelivery: true,
		}))
	}
}
