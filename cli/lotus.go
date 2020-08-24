package cli

import "github.com/urfave/cli/v2"

var LotusCmd = &cli.Command{
	Name:  "lotus",
	Usage: "lotus node",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "start lotus node",
			Action: startLotus,
		},
		{
			Name:   "stop ",
			Usage:  "stop lotus node",
			Action: stopLotus,
		},
	},
}

func startLotus(c *cli.Context) error {
	exec := builder(c)
	return exec.StartLotus()
}

func stopLotus(c *cli.Context) error {
	panic("no support stop lotus, use lotus daemon stop")
}
