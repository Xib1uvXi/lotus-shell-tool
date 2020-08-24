package main

import (
	toolCli "github.com/Xib1uvXi/lotus-shell-tool/cli"
	"github.com/urfave/cli/v2"
	"os"
)

//go:generate go-bindata -pkg bind_data -o bind-data/bind_data.go scripts/

func main() {
	local := []*cli.Command{
		toolCli.LotusCmd,
		toolCli.MinerCmd,
	}

	app := &cli.App{
		EnableBashCompletion: true,
		Commands:             local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "conf",
				Value:   "",
				Usage:   "run config",
				Aliases: []string{"c"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
