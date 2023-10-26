package app

import (
	"asyncfiber/api"
	"asyncfiber/api/middleware"
	"asyncfiber/api/routes"
	"asyncfiber/internal/app/service"
	"asyncfiber/internal/app/tasks"
	"asyncfiber/internal/config"
	"asyncfiber/internal/job"
	"asyncfiber/internal/worker"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server
// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Server() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	coreAPI := api.New().Shutdown(sigChan)

	coreAPI.BackgroundTask(JobLaunch)
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

func AsyncWorker(concurrency int) error {
	w := worker.NewServer(concurrency, worker.Queue{
		config.CriticalQueue: 6, // processed 60% of the time
		config.DefaultQueue:  3, // processed 30% of the time
		config.LowQueue:      1, // processed 10% of the time
	})
	w.HandleFunctions(tasks.Path())
	return w.Run()
}

func JobLaunch() {
	j := job.New()
	if err := j.Scheduler(service.Ping, 5*time.Second); err != nil {
		panic(err)
	}
	if err := j.Launch(); err != nil {
		panic(err)
	}
}
