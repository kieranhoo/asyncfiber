package service

import (
	"asyncfiber/internal/config"
	"asyncfiber/internal/mod/tasks"
	"asyncfiber/pkg/x/mailers"
	"asyncfiber/pkg/x/worker"
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
	err := mailers.SendHTML("pkg/template/email_test.html", "iduchungho@gmail.com")
	return "", err
}
