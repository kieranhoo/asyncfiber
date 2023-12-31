package tasks

import (
	"asyncfiber/pkg/x/worker"
)

const (
	DefaultQueue      string = "default_queue"
	WorkerHealthCheck string = "Worker.HealthCheck"
	WorkerSaveUser    string = "Worker.SaveUser"
)

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck: HandleHealthCheck,
		WorkerSaveUser:    HandleSaveUser,
	}
}
