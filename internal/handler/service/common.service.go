package service

import (
	"qrcheckin/internal/handler/tasks"
	"qrcheckin/internal/mailers"
	"qrcheckin/internal/worker"
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
