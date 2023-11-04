package server

import (
	"asyncfiber/internal/config"

	"github.com/urfave/cli/v2"
)

var Flag = []cli.Flag{
	&cli.StringFlag{
		Name:    "queue",
		Aliases: []string{"q"},
		Usage:   "worker queue name",
		Value:   "default_queue",
		Action: func(ctx *cli.Context, t string) error {
			print(ctx.Args().First())
			config.WorkerQueue = t
			return nil
		},
	},
}
