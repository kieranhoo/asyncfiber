// Copyright 2023 Duc Hung Ho. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.
package main

import (
	"asyncfiber/cmd/cli/app"
	_ "asyncfiber/docs"
	"asyncfiber/internal/api"
	"asyncfiber/internal/api/middleware"
	"asyncfiber/internal/api/routes"
	"asyncfiber/internal/config"
	"asyncfiber/pkg/sentry"
	"asyncfiber/pkg/x/mailers"
	"asyncfiber/pkg/x/worker"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

// @title Student Checkin System
// @version 1.0.0
// @description This is a documentation for the Student Checkin System API
// @host localhost:8000
// @BasePath
func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	coreAPI := api.New().Shutdown(sigChan)
	coreAPI.BackgroundTask(app.JobLaunch)
	coreAPI.Middleware(
		middleware.FiberMiddleware,
		middleware.SentryMiddleware,
	)
	coreAPI.Route(
		routes.Gateway,
		routes.NotFoundRoute,
	)
	coreAPI.Run()
}

// func main() {
// 	utils.LogINFO()
// }
