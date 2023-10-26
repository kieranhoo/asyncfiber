package app

import (
	"github.com/urfave/cli/v2"
)

var Command = []*cli.Command{
	//{
	//	Name:    "worker",
	//	Aliases: []string{"w"},
	//	Usage:   "run worker handle tasks in queue",
	//	Action: func(c *cli.Context) error {
	//		print(c.Args().First())
	//		consume := fmt.Sprintf("comsume#%s", uuid.New().String())
	//		return WorkerLaunch(config.WorkerQueue, consume, 10)
	//	},
	//},
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(c *cli.Context) error {
			return AsyncWorker(10)
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(c *cli.Context) error {
			print(c.Args().First())
			Server()
			return nil
		},
	},
}
