package service

import (
	"asyncfiber/internal/handler/tasks"
	"asyncfiber/internal/mailers"
	"asyncfiber/internal/worker"
)

func Ping() {
	// mailers.SendHTML("iduchungho@gmail.com")
	// log.Println("PONG")
}

func HealthCheck() error {
	return worker.Exec(tasks.WorkerHealthCheck, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	))
}

func Email() (string, error) {
	err := mailers.SendHTML("iduchungho@gmail.com")
	return "", err
}
