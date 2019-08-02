package commands

import (
	"github.com/nknorg/NKNDataPump/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"strconv"
)

func SetAction(app *cli.App) {
	app.Action = func(c *cli.Context) error {
		port := c.Int(cli_FLAG_PORT)
		cfgFile := c.String(cli_FLAG_CONFIG)

		//need at least one of these two parameters: [port | config]
		if 0 == port && "" == cfgFile {
			cli.ShowCommandHelp(c, "")
			return nil
		}

		//port first
		if 0 != port {
			config.PumpConfig.SetNodePort(strconv.Itoa(port))
		}

		//config file
		if "" != cfgFile {
			gwErr := config.NewConfig(cfgFile)

			if nil != gwErr {
				log.Warn(gwErr.FmtOutput())
				return gwErr
			}
		}

		return nil
	}
}
