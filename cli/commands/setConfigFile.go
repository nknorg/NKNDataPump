package commands

import "github.com/urfave/cli"

type SetConfigFlag struct{}

func (*SetConfigFlag) NewFlag() cli.Flag {
	return &(cli.StringFlag{
		Name:   cli_FLAG_CONFIG,
		Usage:  "set config file name",
		Hidden: false,
	})
}
