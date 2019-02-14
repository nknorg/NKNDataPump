package commands

import "github.com/urfave/cli"

type SetPortFlag struct{}

func (*SetPortFlag) NewFlag() cli.Flag {
	return &(cli.IntFlag{
		Name:   cli_FLAG_PORT,
		Usage:  "set destination NKN blockchain node port",
		Hidden: false,
	})
}
