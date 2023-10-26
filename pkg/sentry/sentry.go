package sentry

import (
	"asyncfiber/internal/config"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Init() {
	if config.StageStatus == "prod" && config.SentryDsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: config.SentryDsn,
			BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
				if hint.Context != nil {
					if c, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
						// You have access to the original Context if it panicked
						fmt.Println(utils.CopyString(c.Hostname()))
					}
				}
				fmt.Println(event)
				return event
			},
			Debug:            true,
			AttachStacktrace: true,
		}); err != nil {
			panic(err.Error())
		}

	}
}

func CaptureMessage(c *fiber.Ctx, message string) {
	if config.StageStatus == "prod" && config.SentryDsn != "" {
		if hub := fibersentry.GetHubFromContext(c); hub != nil {
			hub.CaptureMessage(message)
		}
	}
}

func CaptureMessageWithScope(c *fiber.Ctx, key, value, message string) {
	if config.StageStatus == "prod" && config.SentryDsn != "" {
		if hub := fibersentry.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetExtra(key, value)
				hub.CaptureMessage(message)
			})
		}
	}
}
