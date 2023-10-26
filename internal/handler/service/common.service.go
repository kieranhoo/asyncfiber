package service

import (
	"asyncfiber/internal/config"
	"asyncfiber/internal/handler/tasks"
	"asyncfiber/internal/mailers"
	"asyncfiber/internal/worker"
)

func Ping() {
	// mailers.SendHTML("iduchungho@gmail.com")
	// log.Println("PONG")
}

func HealthCheck() error {
	return worker.Exec(config.CriticalQueue, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	))
}

func Email() (string, error) {
	err := mailers.SendHTML("iduchungho@gmail.com")
	return "", err
}
