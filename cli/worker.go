package cli

import "github.com/urfave/cli/v2"

var WorkerCmd = &cli.Command{
	Name:  "worker",
	Usage: "lotus-worker",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "start lotus worker",
			Action: startWorker,
		},
		{
			Name:   "stop",
			Usage:  "stop lotus worker",
			Action: stopWorker,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "force",
					Value: false,
					Usage: "force stop lotus miner, will use kill -9",
				},
			},
		},
	},
}

func startWorker(c *cli.Context) error {
	exec := builder(c)
	return exec.StartWorker()
}

func stopWorker(c *cli.Context) error {
	exec := builder(c)
	return exec.Stop(c.Bool("force"))
}
