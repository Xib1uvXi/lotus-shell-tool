package cli

import "github.com/urfave/cli/v2"

var MinerCmd = &cli.Command{
	Name:  "miner",
	Usage: "lotus-miner",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "start lotus miner",
			Action: startMiner,
		},
		{
			Name:   "stop",
			Usage:  "stop lotus miner",
			Action: stopMiner,
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

func startMiner(c *cli.Context) error {
	exec := builder(c)
	return exec.StartMiner()
}

func stopMiner(c *cli.Context) error {
	exec := builder(c)
	return exec.StopMiner(c.Bool("force"))
}
