package cli

import (
	"github.com/Xib1uvXi/lotus-shell-tool/env"
	"github.com/Xib1uvXi/lotus-shell-tool/exec"
	"github.com/urfave/cli/v2"
)

func builder(c *cli.Context) *exec.Executor {
	confPath := c.String("conf")

	if confPath == "" {
		panic("need set toml config path")
	}

	conf := env.NewDecoder().Decode(confPath)

	return exec.NewExecutor(conf)
}
